package tests

import (
	"HelloGo/Handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedBody := `{status: up}`
	assertResponse(t, req, Handlers.HealthCheckHandler, http.StatusOK, expectedBody)
}

func TestHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/whoohoo", nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedBody := `Hello, whoohoo!`
	assertResponse(t, req, Handlers.TalkBackHandler, http.StatusOK, expectedBody)
}

func assertResponse(t *testing.T, req *http.Request, handlerToTest http.HandlerFunc, expectedStatusCode int, expectedBody string) {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.TalkBackHandler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handlerToTest.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != expectedStatusCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatusCode)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}
