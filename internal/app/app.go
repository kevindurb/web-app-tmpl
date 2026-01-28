package app

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevindurb/web-app-tmpl/internal/database"
	"github.com/kevindurb/web-app-tmpl/internal/sqlc"
	ghttp "maragu.dev/gomponents/http"
)

type App struct {
	Mux       *http.ServeMux
	queries   *sqlc.Queries
	validator *validator.Validate
	decoder   *schema.Decoder
}

func New(pool *pgxpool.Pool) *App {
	database.Up(pool)
	mux := http.NewServeMux()

	a := App{
		Mux:       mux,
		queries:   sqlc.New(pool),
		validator: validator.New(),
		decoder:   schema.NewDecoder(),
	}

	mux.Handle("GET /static/", http.StripPrefix("/static/", a.Static()))
	mux.HandleFunc("GET /healthcheck", a.healthcheck)
	mux.HandleFunc("GET /{$}", ghttp.Adapt(a.indexShow))

	return &a
}
