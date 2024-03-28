package handler

import (
	"fmt"
	"net/http"
	"tiny_url/internal/counter"
	"tiny_url/internal/storage"
	"tiny_url/pkg/utils"
)

type URLHandler struct {
	Storage storage.Storage
	Counter *counter.Counter
}
// NewURLHandler creates a new URLHandler with the given storage.
func NewURLHandler(storage storage.Storage, couter *counter.Counter) *URLHandler {
	return &URLHandler{
		Storage: storage,
		Counter: couter,
	}
}

// ShortenURL handles requests to shorten a URL.
func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form and retrieve the URL.
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	id := h.Counter.Value()
	if id == -1 {
		http.Error(w, "ID range exhausted", http.StatusInternalServerError)
		return
	}
	shortURL := utils.EncodeToBase62(id)
	fullShortURL := fmt.Sprintf("http://localhost:8080/%s", shortURL)

	// Save the mapping and return the shortened URL.
	h.Storage.Save(fullShortURL, originalURL)
	fmt.Fprintf(w, "Shortened URL: %s\n", fullShortURL)
}

// Redirect handles redirection from a shortened URL to the original URL.
func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := "http://localhost:8080" + r.URL.Path
	originalURL, exists := h.Storage.Load(shortURL)
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}
