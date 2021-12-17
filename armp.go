package main

import (
    "io"
    "fmt"
    "net/http"
)

// create a handler struct
type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "Hello")

    fmt.Fprint(res, " World! ")

    res.Write([]byte("❤️"))
}

func main() {
    // create a new handler
    handler := HttpHandler{}    // listen and serve
    http.ListenAndServe(":9000", handler)
}
