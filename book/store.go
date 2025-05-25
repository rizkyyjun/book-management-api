package book

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Store struct {
	books map[string]Book
	mu    sync.RWMutex
}

func NewStore() *Store {
	return &Store{books: make(map[string]Book)}
}

func (s *Store) Create(b Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.books[b.ISBN]; exists {
		return fmt.Errorf("Book with ISBN %s already exists", b.ISBN)
	}
	s.books[b.ISBN] = b
	return nil
}

func (s *Store) GetAll(page, limit int) []Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var books []Book
	for _, b := range s.books {
		books = append(books, b)
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(books) {
		return []Book{}
	}

	if end > len(books) {
		end = len(books)
	}

	return books[start:end]
}

func (s *Store) GetByCriteria(c Criteria) []Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var filtered []Book
	for _, b := range s.books {
		if c.Title != "" && !strings.Contains(strings.ToLower(b.Title), strings.ToLower(c.Title)) {
			continue
		}
		if c.Author != "" && !strings.Contains(strings.ToLower(b.Author), strings.ToLower(c.Author)) {
			continue
		}
		if c.ISBN != "" && b.ISBN != c.ISBN {
			continue
		}
		if c.ReleaseDate != "" && b.ReleaseDate != c.ReleaseDate {
			continue
		}
		filtered = append(filtered, b)
	}

	// Sort
	sortBy := strings.ToLower(c.SortBy)
	order := strings.ToLower(c.Order)

	if sortBy == "" {
		sortBy = "title"
	}
	if order != "desc" {
		order = "asc"
	}

	sort.Slice(filtered, func(i, j int) bool {
		var less bool
		switch sortBy {
		case "title":
			less = strings.ToLower(filtered[i].Title) < strings.ToLower(filtered[j].Title)
		case "author":
			less = strings.ToLower(filtered[i].Author) < strings.ToLower(filtered[j].Author)
		case "isbn":
			less = filtered[i].ISBN < filtered[j].ISBN
		case "release_date":
			less = filtered[i].ReleaseDate < filtered[j].ReleaseDate
		default:
			less = strings.ToLower(filtered[i].Title) < strings.ToLower(filtered[j].Title)
		}
		if order == "desc" {
			return !less
		}
		return less
	})

	// Pagination
	page := c.Page
	limit := c.Limit
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	start := (page - 1) * limit
	end := start + limit

	if start >= len(filtered) {
		return []Book{}
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end]
}

func (s *Store) Get(isbn string) (Book, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.books[isbn]
	return b, ok
}

func (s *Store) Update(isbn string, b Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.books[isbn]; !exists {
		return fmt.Errorf("Book with ISBN %s is not found", isbn)
	}
	s.books[isbn] = b
	return nil
}

func (s *Store) Delete(isbn string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.books[isbn]; !exists {
		return fmt.Errorf("Book with ISBN %s is not found", isbn)
	}
	delete(s.books, isbn)
	return nil
}
