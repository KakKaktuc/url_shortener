package routes

import (
	"net/http"
	"url_shortener/internal/handler"
)

// RegisterRoutes — инициализация всех маршрутов
func RegisterRoutes(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", h.CreateURL)
	mux.HandleFunc("/r/", h.Redirect)

	// можно добавить healthcheck
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return mux
}
