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

func TestGetBalance(t *testing.T) {
	t.Run("Get user balance", func(t *testing.T) {
		req, err := http.NewRequest("GET", config.API_VERSION_ONE+"account/balance", nil)
		if err != nil {
			t.Fatal(err)
		}

		//fake gorilla/mux vars
		vars := map[string]string{
			"user_id": "b7392d0b-8a6f-4436-869a-037d054ea7d5",
		}

		req = mux.SetURLVars(req, vars)
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.GetBalance)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusOK, resp.Body.String())
		}
	})

	t.Run("Get unregistered user balance", func(t *testing.T) {
		req, err := http.NewRequest("GET", config.API_VERSION_ONE+"account/balance", nil)
		if err != nil {
			t.Fatal(err)
		}

		//fake gorilla/mux vars
		vars := map[string]string{
			"user_id": "b7392d0b-8a6f-4436-869a-037d054ea7d9",
		}

		req = mux.SetURLVars(req, vars)
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.GetBalance)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusNotFound, resp.Body.String())
		}
	})
}

func TestDeposit(t *testing.T) {
	t.Run("Deposit to user's account", func(t *testing.T) {
		f := "amount=88.0&user_id=b7392d0b-8a6f-4436-869a-037d054ea7d5"
		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"account/deposit", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.Deposit)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusOK, resp.Body.String())
		}
	})

	t.Run("Missing user id param", func(t *testing.T) {
		f := "amount=88.0"
		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"account/deposit", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.Deposit)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusBadRequest, resp.Body.String())
		}
	})

	t.Run("Deposit by unregistered user account", func(t *testing.T) {
		f := "amount=88.0&user_id=b7392d0b-8a6f-4436-869a-037d054ea7d9"
		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"account/deposit", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.Deposit)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusUnprocessableEntity, resp.Body.String())
		}
	})
}

func TestTransfer(t *testing.T) {
	t.Run("Transfer between two different users account", func(t *testing.T) {
		f := "amount=88.0&user_id=b7392d0b-8a6f-4436-869a-037d054ea7d5&recipient=d78002d0b-9a6f-4436-999a-237d054ea7d5"
		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"account/transfer", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.Transfer)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusOK, resp.Body.String())
		}
	})

	t.Run("Transfer to an unregistered user", func(t *testing.T) {
		f := "amount=88.0&user_id=b7392d0b-8a6f-4436-869a-037d054ea7d5&recipient=d78002d0b-9a6f-4436-999a-237d054ea7d9"
		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"account/transfer", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.Transfer)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusUnprocessableEntity, resp.Body.String())
		}
	})

	t.Run("Transfer more amount than in account", func(t *testing.T) {
		f := "amount=808.0&user_id=b7392d0b-8a6f-4436-869a-037d054ea7d5&recipient=d78002d0b-9a6f-4436-999a-237d054ea7d5"
		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"account/transfer", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.Transfer)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusUnprocessableEntity, resp.Body.String())
		}
	})
}

func TestWithdrawal(t *testing.T) {
	t.Run("Withdraw from user's account", func(t *testing.T) {
		f := "amount=88.0&user_id=b7392d0b-8a6f-4436-869a-037d054ea7d5"
		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"account/withdraw", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.Withdraw)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusOK, resp.Body.String())
		}
	})

	t.Run("Withdraw more than in account", func(t *testing.T) {
		f := "amount=880.0&user_id=b7392d0b-8a6f-4436-869a-037d054ea7d5"
		req, err := http.NewRequest("POST", config.API_VERSION_ONE+"account/withdraw", strings.NewReader(f))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.Withdraw)

		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v : data %v",
				status, http.StatusUnprocessableEntity, resp.Body.String())
		}
	})
}
