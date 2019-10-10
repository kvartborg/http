package http

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kvartborg/http/request"
	"github.com/kvartborg/http/response"
)

type Handler func(*Request) Response

type route struct {
	methods  []string
	handlers []Handler
}

type Server struct {
	http.Handler
	routes map[string]*route
}

func NewServer() *Server {
	return &Server{
		routes: map[string]*route{},
	}
}

func (s *Server) register(methods []string, path string, handlers []Handler) {
	if _, ok := s.routes[path]; ok {
		return
	}

	s.routes[path] = &route{methods, handlers}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res Response
	req := request.New(r)
	route, ok := s.routes[strings.Split(r.RequestURI, "?")[0]]

	if ok {
		for _, method := range route.methods {
			if method != r.Method {
				ok = false
				res = response.NotFound()
			}
		}

		for _, handler := range route.handlers {
			res = handler(req)

			if res.Status() != 0 {
				break
			}
		}
	} else {
		res = response.NotFound()
	}

	if res.Status() >= 300 && res.Status() < 400 {
		http.Redirect(w, r, res.String(), res.Status())
		return
	}

	for key := range res.Header() {
		w.Header().Set(key, res.Header().Get(key))
	}

	w.WriteHeader(res.Status())
	_, err := w.Write(res.Body())

	if err != nil {
		log.Println(err)
	}
}

func (s *Server) Listen(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s)
}
