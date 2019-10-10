package request

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type Request struct {
	Method        string
	URL           *url.URL
	Query         Query
	Param         map[string]string
	Header        http.Header
	Body          []byte
	ContentLength int64
	Host          string
	Form          url.Values
	PostForm      url.Values
	MultipartForm *multipart.Form
	RemoteAddr    string
	RequestURI    string
	TLS           *tls.ConnectionState
	Raw           *http.Request
}

func New(r *http.Request) *Request {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("could not read request body")
	}

	return &Request{
		Method:        r.Method,
		URL:           r.URL,
		Query:         Query(r.URL.Query()),
		Param:         mux.Vars(r),
		Header:        r.Header,
		Body:          body,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		Form:          r.Form,
		PostForm:      r.PostForm,
		MultipartForm: r.MultipartForm,
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
		TLS:           r.TLS,
		Raw:           r,
	}
}
