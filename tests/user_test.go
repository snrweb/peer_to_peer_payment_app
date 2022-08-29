package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/snrweb/peer_to_peer_payment_app/config"
	"github.com/snrweb/peer_to_peer_payment_app/controllers"
)

func TestUserCreation(t *testing.T) {
	t.Run("Add new user", func(t *testing.T) {
		f := "first_name=John&last_name=Doe&email=john_doe@yahoo.com&password=testtest"

		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"user/add", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.CreateUser)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusOK, resp.Body.String())
		}
	})

	t.Run("Add new user when first_name is absent", func(t *testing.T) {
		f := "last_name=Doe&email=john_doe@yahoo.com&password=testtest"

		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"user/add", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.CreateUser)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusBadRequest, resp.Body.String())
		}
	})
}

func TestGetUsers(t *testing.T) {
	t.Run("Get all users", func(t *testing.T) {
		req, err := http.NewRequest("GET", config.API_VERSION_ONE+"user/all", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.GetUsers)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusOK, resp.Body.String())
		}
	})
}

func TestGetUser(t *testing.T) {
	t.Run("Get one user", func(t *testing.T) {
		req, err := http.NewRequest("GET", config.API_VERSION_ONE+"user", nil)
		if err != nil {
			t.Fatal(err)
		}

		//fake gorilla/mux vars
		vars := map[string]string{
			"id": "b7392d0b-8a6f-4436-869a-037d054ea7d5",
		}

		req = mux.SetURLVars(req, vars)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.GetUser)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusOK, resp.Body.String())
		}
	})

	t.Run("Get unregistered user", func(t *testing.T) {
		req, err := http.NewRequest("GET", config.API_VERSION_ONE+"user", nil)
		if err != nil {
			t.Fatal(err)
		}

		//fake gorilla/mux vars
		vars := map[string]string{
			"id": "b7392d0b-8a6f-4436-869a-037d054ea7d9",
		}

		req = mux.SetURLVars(req, vars)

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.GetUser)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusNotFound, resp.Body.String())
		}
	})
}
