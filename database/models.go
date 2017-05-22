package database

import (
	"fmt"
	"io/ioutil"
)

func PopulateIncomeCategoryValues() {
	data := []map[string]string{
		{"Title": "Employment", "Description": "Employment income."},
		{"Title": "Business", "Description": "Business account."},
		{"Title": "Investment", "Description": "Investment account."},
		{"Title": "Savings", "Description": "Savings account."},
		{"Title": "Interest", "Description": "Interest income."},
		{"Title": "Inheritance", "Description": "Inheritance income."},
		{"Title": "Other", "Description": "Other income."},
	}

	db, _ := Connect().Acquire()
	defer Connect().Release(db)

	for _, val := range data {
		_, err := db.Exec(fmt.Sprintf("INSERT INTO income_category(title, description) VALUES('%s', '%s')", val["Title"], val["Description"]))
		if err != nil {
			fmt.Println("Failed with query. ", err)
		}
	}
}

func PopulateExpenseCategoryValues() {
	data := []map[string]string{
		{"Title": "Auto & Transport", "Description": "Auto, transport, liability insurance, tolls, fuel,..."},
		{"Title": "Banking services", "Description": "All banking related services."},
		{"Title": "Charity", "Description": "Charity contributions and gifts."},
		{"Title": "Entertainment", "Description": "Entretainment, travel, movies, TV, games, restaurants."},
		{"Title": "Housing", "Description": "Housing, retn, loans, maintenance, repair, utilities,..."},
		{"Title": "Health Care", "Description": "Health care, drugs,..."},
		{"Title": "Food and Groceries", "Description": "Everyday food, groceries, except restaurants,..."},
		{"Title": "Personal Care", "Description": "Personal care & self-development, sports,..."},
		{"Title": "Savings", "Description": "Contributions to savings account."},
		{"Title": "Business", "Description": "To be founded business, idea protyping and similar expenses."},
		{"Title": "Investments", "Description": "Investment account contributions."},
		{"Title": "Taxes", "Description": "Taxes."},
	}

	db, _ := Connect().Acquire()
	defer Connect().Release(db)

	for _, val := range data {
		_, err := db.Exec(fmt.Sprintf("INSERT INTO expense_category(title, description) VALUES('%s', '%s')", val["Title"], val["Description"]))
		if err != nil {
			fmt.Println("Failed with query. ", err)
		}
	}
}

func PopulateCurrencyValues() {
	data := []map[string]string{
		{"Title": "AUD", "Description": "Australian dollar"},
		{"Title": "CAD", "Description": "Canadian dollar"},
		{"Title": "CHF", "Description": "Swiss franc"},
		{"Title": "CNY", "Description": "Chinese yuan"},
		{"Title": "CZK", "Description": "Czech koruna"},
		{"Title": "DKK", "Description": "Danish krone"},
		{"Title": "EUR", "Description": "Euro"},
		{"Title": "HKD", "Description": "Hong Kong dollar"},
		{"Title": "HIF", "Description": "Hungarian forint"},
		{"Title": "ISK", "Description": "Icelandic króna"},
		{"Title": "JPY", "Description": "Japanese yen"},
		{"Title": "MXN", "Description": "Mexican peso"},
		{"Title": "NZD", "Description": "New Zealand dollar"},
		{"Title": "PLN", "Description": "Polish złoty"},
		{"Title": "RUB", "Description": "Russian ruble"},
		{"Title": "SEK", "Description": "Swedish krona/kronor"},
		{"Title": "SGD", "Description": "Singapore dollar"},
		{"Title": "TRY", "Description": "Turkish lira"},
		{"Title": "USD", "Description": "United States dollar"},
	}

	db, _ := Connect().Acquire()
	defer Connect().Release(db)

	for _, val := range data {
		_, err := db.Exec(fmt.Sprintf("INSERT INTO currency(title, description) VALUES('%s', '%s')", val["Title"], val["Description"]))
		if err != nil {
			fmt.Println("Failed with query. ", err)
		}
	}
}

func CreateTables() {
	db, _ := Connect().Acquire()
	defer Connect().Release(db)

	contents, err := ioutil.ReadFile("database/schema.sql")
	if err != nil {
		fmt.Print(err)
	}

	query := string(contents)

	_, dbErr := db.Exec(query)
	if dbErr != nil {
		fmt.Println("Failed with query. ", dbErr)
	}

}
