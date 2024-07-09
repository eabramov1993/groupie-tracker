package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Обработчик ошибки 404
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 Not Found: %s", r.URL.Path)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Not Found")
}

// Обработчик ошибки 500
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("500 Internal Server Error: %s - %v", r.URL.Path, err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "500 Internal Server Error")
}

// Обработчик ошибки 405 Method Not Allowed
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("405 Method Not Allowed: %s %s", r.Method, r.URL.Path)
	w.Header().Set("Allow", "GET")
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Обработчик ошибки 400 Bad Request
func BadRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("400 Bad Request: %s - %v", r.URL.Path, err)
	http.Error(w, "Bad Request", http.StatusBadRequest)
}
