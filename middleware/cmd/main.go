package main

import (
	"net/http"

	"goprojects/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	logger := middleware.CreateLogger("anaLog")
	http.Handle("/", middleware.Time(logger, hello))
	http.ListenAndServe(":3000", nil)
}
