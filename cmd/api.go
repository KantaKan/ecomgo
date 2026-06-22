package main

import "net/http"


type application struct {
	config config
	//logger
}


func (app *application) mount() http.Handler {

}

type config struct {
	addr string
	db dbConfig

}

type dbConfig struct {
	dsn string
}