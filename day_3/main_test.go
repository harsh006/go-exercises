package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
//Set Gin to Test Mode
gin.SetMode(gin.TestMode)

// Run the other tests
os.Exit(m.Run())
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}
func TestGetUsers(t *testing.T) {
	//req, _ := http.NewRequest("GET", "https://localhost:8080/user-api/user", nil)
	//http.ListenAndServe(req)
	response, _ := http.Get("http://localhost:8080/user-api/user")
	fmt.Println(response)
	try1, _ := io.ReadAll(response.Body)
	fmt.Println(string(try1))
	//testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
	//	// Test that the http status code is 200
	//	statusOK := w.Code == http.StatusOK
	//
	//	// Test that the page title is "Home Page"
	//	// You can carry out a lot more detailed tests using libraries that can
	//	// parse and process HTML pages
	//	p, err := ioutil.ReadAll(w.Body)
	//	pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0
	//
	//	return statusOK && pageOK
	//})
}