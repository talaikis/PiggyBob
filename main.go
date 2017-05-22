package main

import (
	//"./secure"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	//"./session"
	"fmt"
	//"github.com/gin-contrib/gzip"
	"./config"
	"./database"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/twitter"
	"github.com/nicksnyder/go-i18n/i18n"
	"./middleware"
)

var defaultLang = "en"
var store = sessions.NewCookieStore([]byte(os.Getenv("PIG_APP_KEY")), []byte(os.Getenv("PIG_PIG_ENCRYPT_KEY")))

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables.")
	}

	gothic.Store = store
}

func main() {
	goth.UseProviders(
		twitter.New(os.Getenv("PIG_TWITTER_KEY"), os.Getenv("PIG_TWITTER_SECRET"), os.Getenv("PIG_OAUTH_HOST")+"/auth/twitter/callback"),
		facebook.New(os.Getenv("PIG_FACEBOOK_KEY"), os.Getenv("PIG_FACEBOOK_SECRET"), os.Getenv("PIG_OAUTH_HOST")+"/auth/facebook/callback"),
		gplus.New(os.Getenv("PIG_GPLUS_KEY"), os.Getenv("PIG_GPLUS_SECRET"), os.Getenv("PIG_OAUTH_HOST")+"/auth/gplus/callback"),
		linkedin.New(os.Getenv("PIG_LINKEDIN_KEY"), os.Getenv("PIG_LINKEDIN_SECRET"), os.Getenv("PIG_OAUTH_HOST")+"/auth/linkedin/callback"),
	)

	// *****************************************************************************
	//	ROUTES
	// *****************************************************************************

	app := mux.NewRouter()
	app.PathPrefix("/static/assets/").Handler(http.FileServer(http.Dir(".")))

	// TODO load auth middleware
	// TODO load gzip middleware
	// TODO load recover, log, monitor midds

	app.HandleFunc("/", HomeHandler)
	app.HandleFunc("/auth/{provider}/callback", CallbackHandler)
	app.HandleFunc("/logout/{provider}", LogoutHandler)
	app.HandleFunc("/auth/{provider}", ProviderHandler)
	app.HandleFunc("/lang/{lang}/", HomeHandler)
	app.HandleFunc("/lang/{lang}/member_area/", MemberAreaHandler)
	app.HandleFunc("/lang/{lang}}/accounts/", AccountsHandler)
	app.HandleFunc("/lang/{lang}/journal/", JournalHandler)
	app.HandleFunc("/lang/{lang}/balance/", BalanceSheetHandler)
	app.HandleFunc("/lang/{lang}/timeline/", TimeStatsHandler)
	app.HandleFunc("/lang/{lang}/piechart/", PieHandler)
	app.HandleFunc("/lang/{lang}/general/", GeneralJournalHandler)
	app.HandleFunc("/lang/{lang}/privacy/", PrivacyHandler)
	app.HandleFunc("/lang/{lang}/about/", AboutHandler)
	app.NotFoundHandler = http.HandlerFunc(NotFound)

	// Wauthorized
	/*authorized := app.Group("/")
	authorized.Use(auth.RequireLogin())
	{
	authorized.GET("/", c.Login)*/
	if os.Getenv("RUNMODE") != "dev" {
		app.Host(os.Getenv("DOMAIN")).Schemes("https")
	}

	server := &http.Server{
		Handler:      app,
		Addr:         os.Getenv("PIG_HOST"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// *****************************************************************************
//	HANDLER GLOBS
// *****************************************************************************

type PageStruct struct {
	T           i18n.TranslateFunc
	PageTitle   string
	HeaderTitle string
	SiteTitle   string
	CurrentLang string
	L           *config.LangStruct
	P           *config.ProviderIndex
	Strings     map[string]string
}

// *****************************************************************************
//	MEMBER AREA
//*****************************************************************************

func MemberAreaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	lang, T, translations := middleware.Translate(w, r)

	session, _ := store.Get(r, "_sess_")
	email := session.Values["email"]
	if email == nil {
		http.Redirect(w, r, "/lang/" + lang + "/", 302)
	}

	db, _ := database.Connect().Acquire()
	defer database.Connect().Release(db)

	// FIXME parse rows
	query := fmt.Sprintf(`SELECT provider, name, email, avatar_url FROM users WHERE email='%s'`, email)
	rows, e := db.Exec(query)
	if e != nil {
		fmt.Fprintln(w, e)
		return
	}

	for rows.Next() {
		var provider string
		var name string
		var email string
		var avatar_url string

		err := rows.Scan(&provider, &name, &email, &avatar_url)
		if err != nil {
			return err
		}
	}

	translations["Provider"] = provider
	translations["Name"] = string(rows[1])
	translations["Email"] = string(rows[2])
	translations["AvatarUrl"] = string(rows[3])

	t, _ := template.ParseFiles("static/html/member_area.html", "static/html/footer.html", "static/html/header.html", "static/html/user.html")

	err = t.Execute(w, PageStruct{
		PageTitle:   T("members_area"),
		CurrentLang: lang,
		HeaderTitle: T("members_area"),
		SiteTitle:   " | " + os.Getenv("PIG_SITE_TITLE"),
		L:           config.Languages(),
		P:           config.Social(),
		Strings:     translations})
	if err != nil {
		fmt.Println(err.Error())
	}
}

// *****************************************************************************
//	CALLBACK
//*****************************************************************************

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	session, _ := store.Get(r, "_sess_")
	session.Values["email"] = user.Email
	session.Save(r, w)

	db, _ := database.Connect().Acquire()
	defer database.Connect().Release(db)

	cols := "provider, email, name, first_name, last_name, nickname, description, user_id, avatar_url, location, access_token, access_token_secret, refresh_token"
	query := fmt.Sprintf(`INSERT into users (%s) VALUES ('%s', '%s', '%s', '%s',
				'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') ON CONFLICT ON CONSTRAINT email_unique DO UPDATE
				SET access_token='%s', refresh_token='%s' WHERE users.email='%s';`, cols, user.Provider,
		user.Email, user.Name, user.FirstName, user.LastName, user.NickName,
		user.Description, user.UserID, user.AvatarURL, user.Location, user.AccessToken,
		user.AccessTokenSecret, user.RefreshToken, user.AccessToken, user.RefreshToken, user.Email)

	_, e := db.Exec(query)
	if e != nil {
		fmt.Fprintln(w, e)
		return
	}

	lang := session.Values["lang"]
	fmt.Println()
	fmt.Println(lang)
	fmt.Println(lang.(string))

	http.Redirect(w, r, "/lang/"+lang.(string)+"/member_area/", 302)
}

// *****************************************************************************
//	LOGOUT
// *****************************************************************************

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// *****************************************************************************
//	PROVIDER HANDLER
// *****************************************************************************

func ProviderHandler(w http.ResponseWriter, r *http.Request) {
	if user, err := gothic.CompleteUserAuth(w, r); err == nil {
		session, _ := store.Get(r, "_sess_")
		lang := session.Values["lang"]
		session.Values["email"] = user.Email
		if lang == nil {
			lang = defaultLang
		}
		http.Redirect(w, r, "/lang/" + lang.(string) + "/member_area/", 302)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

// *****************************************************************************
// 500
// *****************************************************************************

func ServerError(w http.ResponseWriter, r *http.Request, err ...string) {
	w.Header().Set("Content-Type", "text/html")
	_, T, translations := middleware.Translate(w, r)

	translations["Error"] = err

	t, _ := template.ParseFiles("static/html/error.html", "static/html/footer.html", "static/html/header.html", "static/html/user.html")
	e := t.Execute(w, PageStruct{
		PageTitle:   T("server_error"),
		CurrentLang: lang,
		HeaderTitle: T("server_error"),
		SiteTitle:   " | " + os.Getenv("PIG_SITE_TITLE"),
		L:           config.Languages(),
		P:           config.Social(),
		Strings:     translations})
	if e != nil {
		fmt.Println(err.Error())
	}
}

// *****************************************************************************
//	404
// *****************************************************************************

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// FIXME it not always can know the lang, get a session lang
	lang, T, translations := middleware.Translate(w, r)

	t, _ := template.ParseFiles("static/html/error.html", "static/html/footer.html", "static/html/header.html", "static/html/user.html")
	err = t.Execute(w, PageStruct{
		PageTitle:   T("not_found"),
		CurrentLang: lang,
		HeaderTitle: T("not_found"),
		SiteTitle:   " | " + os.Getenv("PIG_SITE_TITLE"),
		L:           config.Languages(),
		P:           config.Social(),
		Strings:     translations})
	if err != nil {
		fmt.Println(err.Error())
	}
}

//*****************************************************************************
//	HOME HANDLER
// *****************************************************************************

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	lang, T, translations := middleware.Translate(w, r)

	session, _ := store.Get(r, "_sess_")
	session.Values["lang"] = lang
	session.Save(r, w)

	t, _ := template.ParseFiles("static/html/index.html", "static/html/footer.html", "static/html/header.html", "static/html/user.html")

	err = t.Execute(w, PageStruct{
		PageTitle:   T("index_page_title"),
		CurrentLang: lang,
		HeaderTitle: T("index_page_title"),
		SiteTitle:   " | " + os.Getenv("PIG_SITE_TITLE"),
		L:           config.Languages(),
		P:           config.Social(),
		Strings:     translations})
	if err != nil {
		fmt.Println(err.Error())
	}
}

// *****************************************************************************
//	PRIVACY POLICY PAGE
//*****************************************************************************

func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	lang, T, translations := middleware.Translate(w, r)

	t, _ := template.ParseFiles("static/html/privacy.html", "static/html/footer.html", "static/html/header.html", "static/html/user.html")

	err = t.Execute(w, PageStruct{
		PageTitle:   T("privacy_policy"),
		CurrentLang: lang,
		HeaderTitle: T("privacy_policy"),
		SiteTitle:   " | " + os.Getenv("PIG_SITE_TITLE"),
		L:           config.Languages(),
		P:           config.Social(),
		Strings:     translations})
	if err != nil {
		fmt.Println(err.Error())
	}
}

// *****************************************************************************
//	ABOUT PAGE
//*****************************************************************************

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	lang, T, translations := middleware.Translate(w, r)

	t, _ := template.ParseFiles("static/html/about.html", "static/html/footer.html", "static/html/header.html", "static/html/user.html")
	err = t.Execute(w, PageStruct{
		PageTitle:   T("about_header"),
		CurrentLang: lang,
		HeaderTitle: T("about_header"),
		SiteTitle:   " | " + os.Getenv("PIG_SITE_TITLE"),
		L:           config.Languages(),
		P:           config.Social(),
		Strings:     translations})
	if err != nil {
		fmt.Println(err.Error())
	}
}

//*****************************************************************************
//	PIGGYBOB APP ACTUALLY
// *****************************************************************************

func AccountsHandler(w http.ResponseWriter, r *http.Request) {
}

func JournalHandler(w http.ResponseWriter, r *http.Request) {
}

func BalanceSheetHandler(w http.ResponseWriter, r *http.Request) {
}

func TimeStatsHandler(w http.ResponseWriter, r *http.Request) {
}

func PieHandler(w http.ResponseWriter, r *http.Request) {
}

func GeneralJournalHandler(w http.ResponseWriter, r *http.Request) {
}
