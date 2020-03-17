package input_output_middleware

import (
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"log"
	"net/http"
	"time"
)

const transactionKey = "transaction"

func InputOutputMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Print(fmt.Sprintf("Incoming request %s %s", r.Method, r.URL))
		if r.Header.Get("Content-Type") != "" {
			value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
			if value != "application/json" {
				msg := "Content-Type header is not application/json"
				http.Error(w, msg, http.StatusUnsupportedMediaType)
				log.Print(fmt.Sprintf("Invalid Content-Type (%s)", value))
				elapsed := time.Since(start)
				log.Print(fmt.Sprintf("Outgoing response - %s %s %s", r.Method, r.URL, elapsed))
				return
			}
		}
		w.Header().Add("Content-Type", "application/json")
		h.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Print(fmt.Sprintf("Outgoing response - %s %s %s", r.Method, r.URL, elapsed))

	})
}
