package test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/ryanadiputraa/ggen-template/app/healthcheck"
	"github.com/ryanadiputraa/ggen-template/pkg/respwr"
)

func TestHealthcheckHandler(t *testing.T) {
	db, _ := newMockDB(t)
	server := newServer(db)
	defer server.Close()

	srvErr := make(chan error, 1)
	runServer(server, srvErr)

	resp, err := http.Get("http://localhost:8080/healthcheck")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	expected := respwr.ResponseData[healthcheck.Healthcheck]{
		Data: healthcheck.Healthcheck{
			Status: "ok",
		},
	}

	var response respwr.ResponseData[healthcheck.Healthcheck]
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}

	if !reflect.DeepEqual(expected, response) {
		t.Errorf("expected response body to be %v; got %v", expected, response)
	}
}
