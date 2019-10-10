package http_test

import (
	"log"
	gohttp "net/http"
	"testing"
	"time"

	"github.com/kvartborg/http"
	"github.com/kvartborg/http/response"
)

func Hello(req *http.Request) http.Response {
	return response.Text("Hello world")
}

func Auth(req *http.Request) http.Response {
	if !req.Query.Has("auth") {
		return response.Unauthorized()
	}

	return response.Next()
}

func TestServer(t *testing.T) {
	server := http.NewServer()
	server.Get("/", Auth, Hello)

	if err := server.Listen(3000); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkGoImplementation(b *testing.B) {
	var failed int
	mux := gohttp.NewServeMux()
	mux.HandleFunc("/", func(w gohttp.ResponseWriter, r *gohttp.Request) {
		r.Body.Close()
		w.Write([]byte("OK"))
	})

	server := &gohttp.Server{
		Addr:         ":40000",
		Handler:      mux,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		server.ListenAndServe()
	}()

	time.Sleep(time.Second)

	for i := 0; i < b.N; i++ {
		client := gohttp.Client{Timeout: time.Second}

		res, err := client.Get("http://localhost:40000/")

		if res != nil {
			res.Body.Close()
		}

		if err != nil {
			failed++
			continue
		}

		if res.StatusCode != 200 {
			b.Fatal("did not receive a status code of 200 ok")
		}
	}

	b.Logf("Failed: %d", failed)
	time.Sleep(time.Second)
	server.Close()
	time.Sleep(time.Second)
}

func BenchmarkNewImplementation(b *testing.B) {
	var failed int
	mux := http.NewServer()
	mux.Get("/", func(req *http.Request) http.Response {
		return response.Text("OK")
	})

	server := &gohttp.Server{
		Addr:         ":40001",
		Handler:      mux,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		server.ListenAndServe()
	}()

	time.Sleep(time.Second)

	for i := 0; i < b.N; i++ {
		client := gohttp.Client{Timeout: time.Second}

		res, err := client.Get("http://localhost:40001/")

		if res != nil {
			res.Body.Close()
		}

		if err != nil {
			failed++
			continue
		}

		if res.StatusCode != 200 {
			b.Fatal("did not receive a status code of 200 ok")
		}
	}

	b.Logf("Failed: %d", failed)
	time.Sleep(time.Second)
	server.Close()
	time.Sleep(time.Second)
}
