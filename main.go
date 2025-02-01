package main

import (
	"birdton/configurations"
	"birdton/handlers"
	"net/http"

	"github.com/kkdai/twitter"
)

func main() {

	config := configurations.GetConfiguration()

	xClient := twitter.NewServerClient(config.ConsumerKey, config.ConsumerSecret)

	client := handlers.NewServerClient(xClient, config)

	http.HandleFunc("/maketoken", client.GetTwitterToken)
	http.HandleFunc("/request", client.RedirectUserToTwitter)
	http.HandleFunc("/follow", client.GetFollower)
	http.HandleFunc("/followids", client.GetFollowerIDs)
	http.HandleFunc("/time", client.GetTimeLine)
	http.HandleFunc("/user", client.GetUserDetail)

	http.ListenAndServe(config.ServeAddr, nil)

}
