package main

import (
	"context"
	"log"

	"simpleAPI/core/config"
	"simpleAPI/core/db"
	"simpleAPI/core/server"
	"simpleAPI/internal/middleware"
	"simpleAPI/internal/service"

	pmods "simpleAPI/internal/models/payments"
	"simpleAPI/internal/models/users"
	"simpleAPI/internal/views/auth"
	pviews "simpleAPI/internal/views/payments"

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

	// users models settings
	uc := users.New(dbs)

	// auth views settings
	au := auth.New(
		cfg.Keys.SecretKey,
		cfg.Keys.RefreshKey,
		uc,
	)

	// payments settings
	db := pmods.New(dbs)
	svc := service.New(db)
	p := pviews.New(svc)

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

	srv := server.New(cfg.Server.BindAddr)
	log.Printf("[*] Starting server on %s\n", cfg.Server.BindAddr)
	srv.ListenAndServe(ctx, r)
}
