package main

import (
	"context"
	"log"

	"simpleAPI/internal/config"
	"simpleAPI/internal/core/db"
	"simpleAPI/internal/core/server"
	"simpleAPI/internal/moneycount/middleware"
	"simpleAPI/internal/moneycount/service"

	"simpleAPI/internal/moneycount/handlers/auth"
	pviews "simpleAPI/internal/moneycount/handlers/payments"
	"simpleAPI/internal/moneycount/models/database"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.New(".", "settings", "yaml")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// db settings
	dbs, err := db.New(ctx, cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer dbs.Close()

	db := database.New(dbs)

	// payments settings
	svc := service.New(db)
	p := pviews.New(svc)

	// auth views settings
	au := auth.New(
		cfg.Keys.SecretKey,
		cfg.Keys.RefreshKey,
		svc,
	)

	// middleware settings
	mvs := middleware.New(cfg.Keys.SecretKey)

	// routing settings
	r := mux.NewRouter()
	r.Use(*mvs.MiddleFuncs...)
	r.HandleFunc("/api/sign_in", au.SignIn)
	r.HandleFunc("/api/sign_up", au.SignUp)
	r.HandleFunc("/api/refresh", au.Refresh)
	r.HandleFunc("/api/payments", p.Create).Methods("POST")
	r.HandleFunc("/api/payments", p.List).Methods("GET")
	r.HandleFunc("/api/payments/{id}", p.Retrieve).Methods("GET")
	r.HandleFunc("/api/payments/{id}", p.Update).Methods("PUT")
	r.HandleFunc("/api/payments/{id}", p.Destroy).Methods("DELETE")
	r.HandleFunc("/api/payments-group/by-month", p.GroupByMonth)

	srv := server.New(cfg.Server.BindAddr)
	log.Printf("[*] Starting server on %s\n", cfg.Server.BindAddr)
	srv.ListenAndServe(ctx, r)
}
