package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/httprate"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()
	app.middlewares(r)

	// MVP 1.0
	// ----------------- General -----------------
	r.Group(func(r chi.Router) {
		r.Get("/v1/health", app.healthcheckHandler)
		r.Get("/v1/stats", app.statsHandler)
	})

	r.Group(func(r chi.Router) {
		// r.Use(jwtauth.Verifier(tokenAuth))
		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		// r.Use(jwtauth.Authenticator)
		// MVP 1.0
		// ----------------- Payments -----------------
		r.Route("/v1/payments", func(r chi.Router) {
			r.Get("/", app.listPaymentHandler)
			r.Post("/", app.createPaymentHandler)
			r.Get("/{id}", app.viewPaymentHandler)
			r.Delete("/{id}", app.deletePaymentHandler)
		})

		// MVP 2.0
		// ----------------- Payouts -----------------
		r.Route("/v1/payouts", func(r chi.Router) {
			r.Post("/", app.createPayoutHandler)
			r.Get("/", app.listPayoutHandler)
			r.Get("/{id}", app.viewPayoutHandler)
		})

		// MVP 3.0
		// ----------------- Invoices -----------------
		r.Route("/v1/invoices", func(r chi.Router) {
			r.Post("/", app.createInvoiceHandler)
			r.Get("/", app.listInvoiceHandler)
			r.Get("/{id}", app.viewPaymentHandler)
			r.Delete("/{id}", app.deleteInvoiceHandler)
		})

		// MVP 4.0
		// ----------------- Wallets -----------------
		r.Route("/v1/wallets", func(r chi.Router) {
			r.Post("/", app.createWalletHandler)
			r.Get("/", app.listWalletHandler)
			r.Get("/{id}", app.viewWalletHandler)
			r.Put("/{id}", app.updateWalletHandler)
			r.Delete("/{id}", app.deleteWalletHandler)
		})

		// MVP 5.0
		// ----------------- Users -----------------
		r.Route("/v1/users", func(r chi.Router) {
			r.Get("/{id}/stats", app.healthcheckHandler)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		app.ResponseErrorNotFound(w, r, nil, nil)
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		app.ResponseErrorMethodNotAllowed(w, r, nil, nil)
	})
	// Return the chi instance.
	return r
}

func (app *application) middlewares(r *chi.Mux) {
	logger := httplog.NewLogger("stedigate", httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          false,
		RequestHeaders:   true,
		JSON:             false,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": app.config.App.Version,
			"env":     app.config.App.Env,
		},
		QuietDownRoutes: []string{
			"/",
			"/v1/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		TimeFieldFormat: time.RFC3339Nano,
		// SourceFieldName: "source",
	})

	r.Use(
		httplog.RequestLogger(logger),
		middleware.Recoverer,
		middleware.GetHead,
		middleware.CleanPath,
		middleware.RealIP,
		middleware.RequestID,
		middleware.Timeout(60),
		middleware.RedirectSlashes,
		middleware.StripSlashes,
		middleware.AllowContentEncoding("deflate", "gzip"),
		middleware.AllowContentType("application/json"),
		middleware.ContentCharset("UTF-8", "Latin-1", ""),
		middleware.NoCache,
		cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
		middleware.Heartbeat("/v1/ping"),
		httprate.LimitByIP(100, 1*time.Minute),
	)
}
