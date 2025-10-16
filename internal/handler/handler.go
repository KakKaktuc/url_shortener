package handler

import (
	"net/http"
	"encoding/json"
	"url_shortener/internal/model"
    "url_shortener/internal/database"
	//"url_shortener/internal/transport"

)

type Handler struct {
    repo *database.Database 
}
//Конструктор - создаёт handler с доступом к бд
func NewHandler(repo *database.Database) *Handler {
    return &Handler{repo:repo}
}

//POST /shorten - создаёт короткую ссылку
func (h *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Original string `json:"original_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	url := &model.URL{
		Original: req.Original,
		Short:    model.GenerateShortCode(6),
		Visits:   0,
	}

	if err := h.repo.Create(r.Context(), url); err != nil {
		http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"short_url": url.Short,
		"original":  url.Original,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
 

//Get /r/{short} - редирект по короткому коду
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
    short := r.URL.Path[len("/r/"):]
    original, err := h.repo.GetOriginal(r.Context(), short)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    http.Redirect(w, r, original, http.StatusFound)
}