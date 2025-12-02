package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/CHILLERAN/Snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql" // New import
)

type Application struct{
	logger *slog.Logger
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	address := flag.String("addr", ":4000", "HTTP Network Address")
	cs := flag.String("connectionstring", "web:password@/snippetbox?parseTime=true","MySQL data source name") // cs means connectionstring

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(loggerHandler)

		db, err := openDB(*cs)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }

	defer db.Close()

	templateCache, err := newTemplateCache()
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }

	app := &Application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	flag.Parse()

	app.logger.Info("Starting server", "Address", fmt.Sprintf("http://localhost%v", *address))

	err = http.ListenAndServe(*address, app.routes())

	app.logger.Error(err.Error())
	os.Exit(1)
}

func openDB(cs string) (*sql.DB,error) {
	db, err := sql.Open("mysql", cs)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}