package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main(){
	fmt.Println("Hello UK")
	godotenv.Load()
	PORTSTRING := os.Getenv("PORT")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))
	
	  v1Router := chi.NewRouter()
	  v1Router.Post("/healthz", handlerReadiness)

	  v1Router.Post("/err", handlerError)

	  router.Mount("/v1", v1Router)


	srv := &http.Server{
		Handler: router,
		Addr: ":" + PORTSTRING,
	}

	log.Printf("server starting on port %v", PORTSTRING)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


	

	if PORTSTRING == "" {
		log.Fatal(" error loading port from the env file")

	}
	fmt.Println("Port:", PORTSTRING)

}