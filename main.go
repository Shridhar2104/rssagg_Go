package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)
func main(){

	fmt.Println("Hello World")
	godotenv.Load()

	portString:=os.Getenv("PORT")

	if portString ==""{
		log.Fatal("Port not found")
	}
	fmt.Println(portString)

	r:= chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders:	[]string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	v1r:=chi.NewRouter()
	v1r.Get("/healthz", handlerReadiness)
	
	v1r.Get("/err", handleError)



	r.Mount("/v1", v1r)

	srv:=&http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", portString),
	}
	err:=srv.ListenAndServe()

	if err!=nil{
		log.Fatal(err)

	}




}