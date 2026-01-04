package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/CHILLERAN/Snippetbox/internal/models"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql" // New import
)

type Application struct{
	logger 		   *slog.Logger
	snippets 	   *models.SnippetModel
	users          *models.UserModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
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

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
    sessionManager.Store = mysqlstore.New(db)
    sessionManager.Lifetime = 12 * time.Hour

	app := &Application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
		users:          &models.UserModel{DB: db},
		templateCache: templateCache,
		formDecoder: formDecoder,
		sessionManager: sessionManager,
	}

	flag.Parse()

	app.logger.Info("Starting server", "Address", fmt.Sprintf("https://localhost%v", *address))

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
    }

	server := &http.Server{
		Addr: *address,
		Handler: app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig: tlsConfig,

		IdleTimeout:  time.Minute,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
	}

    err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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