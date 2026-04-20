package main
//Contains code for build application chapter
import (
	"log"
	"net/http"
	server "learnGoWithTests/buildApp"
)

func main() {
	s := server.NewPlayerServer(server.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", s))
}
