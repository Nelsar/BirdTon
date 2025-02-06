package handlers

import (
	"birdton/configurations"
	"birdton/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type ServeClient struct {
	client *services.TwitterService
	config *configurations.Configuration
}

func NewServerHandler(client *services.TwitterService) *ServeClient {
	return &ServeClient{
		client: client,
	}
}

func (s ServeClient) GetFollowersUserId(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Query().Get("user_id")

	if userId == "" {
		http.Error(w, "userID обязателен", http.StatusBadRequest)
		return
	}

	following, err := s.client.FollowersById(userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка API: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(following)
}

func (s ServeClient) GetFolowingCheck(w http.ResponseWriter, r *http.Request) {
	sourceID := r.URL.Query().Get("source_id")
	targetID := r.URL.Query().Get("target_id")

	if sourceID == "" || targetID == "" {
		http.Error(w, "sourceID и targetID обязательны", http.StatusBadRequest)
		return
	}

	checkFolowing, err := s.client.FollowinfCheck(sourceID, targetID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка API: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]bool{"is_following": checkFolowing}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
