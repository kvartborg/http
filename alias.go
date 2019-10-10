package http

import (
	"net/http"

	"github.com/kvartborg/http/request"
	"github.com/kvartborg/http/response"
)

type Request = request.Request
type Response = response.Response
type Header = http.Header

func Next() Response {
	return response.Next()
}
