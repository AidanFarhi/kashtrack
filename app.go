package main

import (
	"database/sql"
	"html/template"
	"kashtrack/handler"
	"kashtrack/logger"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"
)

func main() {

	logger.InitLogger(os.Getenv("LOG_FILE"))

	db, err := sql.Open("sqlite", "db/expense.db")
	if err != nil {
		logger.Logger.Fatal(err)
	}

	t := template.Must(template.ParseGlob("web/templates/*.html"))
	fs := http.FileServer(http.Dir("web"))
	m := http.NewServeMux()

	m.Handle("/web/", http.StripPrefix("/web/", fs))
	m.HandleFunc("/", handler.IndexHandler(db, t))
	m.HandleFunc("POST /add_expense", handler.AddExpenseHandler(db, t))
	m.HandleFunc("POST /login", handler.LoginHandler(db, t))
	m.HandleFunc("POST /logout", handler.LogoutHandler(db, t))
	m.HandleFunc("GET /expense_distribution", handler.ExpenseDistributionHandler(db))

	server := http.Server{
		Addr:    "0.0.0.0:443",
		Handler: m,
	}

	logger.Logger.Println("starting redirect listener")
	go http.ListenAndServe("0.0.0.0:80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirectUrl := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
	}))

	logger.Logger.Println("starting server")
	err = server.ListenAndServeTLS(
		os.Getenv("CERT_PATH"),
		os.Getenv("PRIVATE_KEY_PATH"),
	)
	logger.Logger.Fatal(err)
}
