package middleware

import (
  "net/http"
)

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
		fmt.Println(err.Error())
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
