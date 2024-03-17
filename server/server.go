package server

import (
	"net/http"
)

func StartServer() {
	http.HandleFunc("/api/actors", api.AddActorHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
