package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	testMux = mux.NewRouter()
	endpoints(testMux)
	m.Run()
}

func Test_mainHandler(t *testing.T) {
	data := []int{}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/Name", bytes.NewBuffer(payload))
	req.Header.Add("content-type", "plain/text")

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)

	//=========================================================

	req = httptest.NewRequest(http.MethodGet, "/Name", nil)
	req.Header.Add("content-type", "plain/text")

	rr = httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)
}
