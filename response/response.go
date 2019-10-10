package response

import (
	"net/http"
)

var defaultHeader http.Header = http.Header{}

func SetDefaultHeaders(header http.Header) {
	defaultHeader = header
}

type Response struct {
	status int
	header http.Header
	body   []byte
}

func New(status int, header http.Header, body []byte) Response {
	return Response{status, header, body}
}

func (r *Response) Status() int {
	return r.status
}

func (r *Response) Header() http.Header {
	return r.header
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.body)
}
