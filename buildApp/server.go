package server

import (
	"fmt"
	"net/http"
	"strings"
)
type PlayerServer struct {
	store PlayerStore
}
func (p * PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

func NewPlayerServer(p PlayerStore) *PlayerServer{
	return &PlayerServer{p}
}