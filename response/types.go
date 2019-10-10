package response

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Ok() Response {
	return New(http.StatusOK, defaultHeader, []byte(""))
}

func Error(errors ...string) Response {
	return New(http.StatusInternalServerError, defaultHeader, []byte(strings.Join(errors, " ")))
}

func NoResponse() Response {
	return New(http.StatusNotFound, defaultHeader, []byte("No response was returned, check your handler for this route"))
}

func NotFound() Response {
	return New(http.StatusNotFound, defaultHeader, []byte("404 - Not Found"))
}

func BadRequest(msg string) Response {
	return New(http.StatusBadRequest, defaultHeader, []byte("400 - Bad request"))
}

func Unauthorized() Response {
	return New(http.StatusUnauthorized, defaultHeader, []byte("401 - Unautherized"))
}

func Forbidden() Response {
	return New(http.StatusForbidden, defaultHeader, []byte("403 - Forbidden"))
}

func Text(text string) Response {
	return New(http.StatusOK, defaultHeader, []byte(text))
}

func Json(d interface{}) Response {
	data, err := json.Marshal(d)

	if err != nil {
		return Error(err.Error())
	}

	return New(http.StatusOK, http.Header{"Content-Type": {"application/json"}}, []byte(data))
}

func Next() Response {
	return New(0, defaultHeader, make([]byte, 0))
}

func Redirect(url string) Response {
	return New(301, defaultHeader, []byte(url))
}

func View(path string, data interface{}) Response {
	t := template.Must(template.ParseFiles(path))
	var buffer bytes.Buffer
	err := t.Execute(&buffer, data)

	if err != nil {
		log.Println(err)
	}

	return New(http.StatusOK, defaultHeader, buffer.Bytes())
}
