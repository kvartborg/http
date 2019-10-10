# http

> **This is an experiment**, the Go standard implementation is great and should
always be used instead of this!

http is a reinterpretation of the http package shipped with Go standard
library. The core reason behind this experiment is to remove common boilerplate
and allow for better structure and reuse of handlers across projects.

The primary change is the new handler signature which is a function that receives
a request and outputs a response.
```go
type Handler func(*Request) Response
```

### Example

```go
package main

import (
  "github.com/kvartborg/http"
  "github.com/kvartborg/http/response"
)

func Hello(*http.Request) http.Response {
  return response.Text("Hello world!")
}

func Authenticate(req *http.Request) http.Response {
  if req.Query.Has("auth") {
    return response.Next()
  }

  return response.UnAuthorized()
}

func main() {
    server := http.NewServer()
    server.Get("/", Authenticate, Hello)
    server.Listen(3000)
}
```

### License
This project is licensed under the [MIT License](https://github.com/kvartborg/http/blob/master/LICENSE).
