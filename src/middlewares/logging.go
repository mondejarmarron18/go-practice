package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		requestTime := time.Now().Format("2006-01-02 15:04:05")

		if urlPath == "/favicon.ico" {
			urlPath = "/"
		}

		fmt.Printf("%s %s %s\n", requestTime, r.Method, urlPath)

		next.ServeHTTP(w, r)
	})
}
