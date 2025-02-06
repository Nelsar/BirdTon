package main

import (
	"birdton/configurations"
	"birdton/handlers"
	"birdton/services"
	"net/http"
)

func main() {

	config := configurations.GetConfiguration()

	xService := services.NewTwitterService(config.TwitterToken)

	handler := handlers.NewServerHandler(xService)

	http.HandleFunc("/following", handler.GetFollowersUserId)
	http.HandleFunc("/is_following", handler.GetFolowingCheck)

	http.ListenAndServe(config.HostAddr, nil)

}
