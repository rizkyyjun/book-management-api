package book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func MakeHandler(store *Store) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var b Book
			if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
				http.Error(w, "invalid body", http.StatusBadRequest)
				return
			}
			if err := store.Create(b); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			Log("Book created: " + b.ISBN)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Book successfully created",
				"book":    b,
			})
		case http.MethodGet:
			page, _ := strconv.Atoi(r.URL.Query().Get("page"))
			limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

			if page < 1 {
				page = 1
			}
			if limit < 1 {
				limit = 10
			}

			books := store.GetAll(page, limit)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(books)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		isbn := strings.TrimPrefix(r.URL.Path, "/books/")
		switch r.Method {
		case http.MethodGet:
			b, ok := store.Get(isbn)
			if !ok {
				http.Error(w, fmt.Sprintf("Book with ISBN %s is not found", isbn), http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(b)
		case http.MethodPut:
			var b Book
			if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
				http.Error(w, "invalid body", http.StatusBadRequest)
				return
			}
			if err := store.Update(isbn, b); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Book successfully updated",
				"book":    b,
			})
		case http.MethodDelete:
			if err := store.Delete(isbn); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			Log("Book deleted: " + isbn)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": fmt.Sprintf("Book with ISBN %s successfully deleted", isbn),
			})
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/books/get-by-criteria", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var criteria Criteria
		if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		results := store.GetByCriteria(criteria)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	return mux
}
