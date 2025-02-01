package handlers

import (
	"birdton/configurations"
	"fmt"
	"net/http"

	"github.com/kkdai/twitter"
)

type ServeClient struct {
	client *twitter.ServerClient
	config *configurations.Configuration
}

func NewServerClient(client *twitter.ServerClient, config *configurations.Configuration) *ServeClient {
	return &ServeClient{
		client: client,
		config: config,
	}
}

func (s *ServeClient) GetTwitterToken(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")

	s.client.CompleteAuth(tokenKey, verificationCode)
	timelineURL := fmt.Sprintf("http://%s/time", r.Host)

	http.Redirect(w, r, timelineURL, http.StatusTemporaryRedirect)
}

func (s ServeClient) RedirectUserToTwitter(w http.ResponseWriter, r *http.Request) {

	url, _ := s.client.GetAuthURL(s.config.CallbackURL)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (s ServeClient) GetFollower(w http.ResponseWriter, r *http.Request) {
	followers, bits, _ := s.client.QueryFollower(10)
	fmt.Println("Followers=", followers)
	fmt.Fprintf(w, "The item is: "+string(bits))
}

func (s ServeClient) GetFollowerIDs(w http.ResponseWriter, r *http.Request) {
	followers, bits, _ := s.client.QueryFollowerIDs(10)
	fmt.Println("Follower IDs=", followers)
	fmt.Fprintf(w, "The item is: "+string(bits))
}

func (s ServeClient) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	followers, bits, _ := s.client.QueryFollowerById(2244994945)
	fmt.Println("Follower Detail of =", followers)
	fmt.Fprintf(w, "The item is: "+string(bits))
}

func (s ServeClient) GetTimeLine(w http.ResponseWriter, r *http.Request) {
	timeline, bits, _ := s.client.QueryTimeLine(1)
	fmt.Println("TimeLine=", timeline)
	fmt.Fprintf(w, "The item is: "+string(bits))

}
