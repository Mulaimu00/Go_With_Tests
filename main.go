package main
//Contains code for build application chapter
import (
	"log"
	"net/http"
	server "learnGoWithTests/buildApp"
)
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	s := server.NewPlayerServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":5000", s))
}