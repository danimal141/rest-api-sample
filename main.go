package main

import "github.com/danimal141/rest-api-sample/api/router"

func main() {
	r := router.NewRouter()
	r.Logger.Fatal(r.Start(":8080"))
}
