package http

import (
	"net/http"

	"github.com/kvartborg/http/response"
)

func (s *Server) Get(path string, handlers ...Handler) {
	s.register([]string{http.MethodGet}, path, handlers)
}

func (s *Server) Post(path string, handlers ...Handler) {
	s.register([]string{http.MethodPost}, path, handlers)
}

func (s *Server) Put(path string, handlers ...Handler) {
	s.register([]string{http.MethodPut}, path, handlers)
}

func (s *Server) Delete(path string, handlers ...Handler) {
	s.register([]string{http.MethodDelete}, path, handlers)
}

func (s *Server) Any(path string, handlers ...Handler) {
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodTrace,
	}
	s.register(methods, path, handlers)
}

func (s *Server) View(path, view string) {
	s.Get(path, func(*Request) Response {
		return response.View(view, nil)
	})
}
