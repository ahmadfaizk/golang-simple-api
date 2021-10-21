package main

import (
	"golang-simple-api/db"
	"golang-simple-api/handler"
	"golang-simple-api/router"
	"golang-simple-api/store"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	d := db.New()
	db.AutoMigrate(d)
	as := store.NewArticleStore(d)
	h := handler.NewHandler(as)
	h.Register(v1)
	r.Logger.Fatal(r.Start(":8080"))
}
