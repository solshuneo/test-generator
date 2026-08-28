package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"test-generator/generator"

	"github.com/goccy/go-yaml"
)

type NumberResponse struct {
	Data int `yaml:"data"`
}

func TestPostNumber(t *testing.T) {
	router := SetupRouter()
	t.Run("Min = max", func(t *testing.T) {
		w := httptest.NewRecorder()
		number := generator.Variable{
			Min: 10,
			Max: 10,
		}
		data, err := yaml.Marshal(number)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest("POST", "/number", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/yaml")
		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", w.Code)
		}
		if w.Body.String() != "" {
			var res NumberResponse
			err := yaml.Unmarshal(w.Body.Bytes(), &res)
			if err != nil {
				t.Fatal(err)
			}
			guess := res.Data
			if number.Min != guess {
				t.Errorf("expected number to be %d, got %d", number.Min, guess)
			}
		} else {
			t.Errorf("expected status code 200, got %d", w.Code)
		}
	})
	t.Run("Min > Max", func(t *testing.T) {
		w := httptest.NewRecorder()
		number := generator.Variable{
			Min: 10,
			Max: 5,
		}
		data, err := yaml.Marshal(number)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest("POST", "/number", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/yaml")
		router.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", w.Code)
		}
	})
	t.Run("Min < Max", func(t *testing.T) {
		w := httptest.NewRecorder()
		number := generator.Variable{
			Min: 5,
			Max: 10,
		}
		data, err := yaml.Marshal(number)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest("POST", "/number", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/yaml")
		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", w.Code)
		}
		if w.Body.String() != "" {
			var res NumberResponse
			err := yaml.Unmarshal(w.Body.Bytes(), &res)
			if err != nil {
				t.Fatal(err)
			}
			guess := res.Data
			if guess < number.Min || number.Max < guess {
				t.Errorf("expected number to be between %d and %d, got %d", number.Min, number.Max, guess)
			}
		} else {
			t.Errorf("expected status code 200, got %d", w.Code)
		}
	})
}
