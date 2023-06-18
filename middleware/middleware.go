package middleware

import "net/http"

func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
