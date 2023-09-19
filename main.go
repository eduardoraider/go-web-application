package main

import (
	"net/http"
	"web-application/routes"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
