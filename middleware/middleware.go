package middleware

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

var defaultLang = "en"

func Translate(w http.ResponseWriter, r *http.Request) (string, i18n.TranslateFunc, map[string]string) {
	path := r.URL.Path
	lang := ""
	if strings.HasPrefix(path, "/lang") {
		lang = path[6:8]
	} else {
		lang = defaultLang
	}

	i18n.MustLoadTranslationFile("locale/" + lang + ".all.json")
	T, err := i18n.Tfunc(lang)
	if err != nil {
		log.Fatal("Error loading translation file ", err.Error())
	}

	translations := map[string]string{
		"Accounts":       T("accounts"),
		"AddIncome":      T("add_income"),
		"AddExpense":     T("add_expense"),
		"Journal":        T("journal"),
		"Reports":        T("reports"),
		"About":          T("about"),
		"BalanceSheet":   T("balance_sheet"),
		"Timeline":       T("timeline"),
		"ByCategory":     T("by_category"),
		"InOut":          T("in_out"),
		"GeneralJournal": T("general_journal"),
		"Copyright":      T("copyright"),
		"Privacy":        T("privacy"),
		"BaseUrl":        os.Getenv("PIG_BASE_URL")}
	return lang, T, translations
}

type PageStruct struct {
	T           i18n.TranslateFunc
	PageTitle   string
	HeaderTitle string
	SiteTitle   string
	CurrentLang string
	L           *LangStruct
	P           *ProviderIndex
	Strings     map[string]string
}

type LangStruct struct {
	Languages    []string
	LanguagesMap map[string]string
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func Languages() *LangStruct {
	langs := make(map[string]string)

	langs["en"] = "English"
	langs["de"] = "Deutch"
	langs["lt"] = "Lietuvių"

	var lang_keys []string
	for k := range langs {
		lang_keys = append(lang_keys, k)
	}
	sort.Strings(lang_keys)

	var langIndex = &LangStruct{Languages: lang_keys, LanguagesMap: langs}

	return langIndex
}

func Social() *ProviderIndex {
	providers := make(map[string]string)
	providers["facebook"] = "Facebook"
	providers["gplus"] = "Google"
	providers["linkedin"] = "Linkedin"
	providers["twitter"] = "Twitter"

	var provider_keys []string
	for k := range providers {
		provider_keys = append(provider_keys, k)
	}
	sort.Strings(provider_keys)

	var providerIndex = &ProviderIndex{Providers: provider_keys, ProvidersMap: providers}

	return providerIndex
}
