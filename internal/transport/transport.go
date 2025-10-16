package transport

type URLResponse struct {
	ShortURL string `json:"short_url"`
	Original string `json:"original_url"`
}