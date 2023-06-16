package main

import (
	"net/http"
	"go-store/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":3333", nil)
}
