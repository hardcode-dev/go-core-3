package main

import (
	"go-dev-v3/13-api/2-api/pkg/api"
	"go-dev-v3/18-microservices/movie-reviews/pkg/queue"
	"go-dev-v3/18-microservices/movie-reviews/pkg/reviews"

	"github.com/gorilla/mux"
)

type Server struct {
	api   *api.API
	queue *queue.Interface
	db    *reviews.DB
}

func New() *Server {
	s := Server{
		api:   api.New(mux.NewRouter()),
		queue: nil,
		db:    nil,
	}
	return &s
}

func main() {
	srv := New()
	srv.run()
}

func (s *Server) run() {

}
