package server

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"golang.org/x/crypto/bcrypt"
)

var router *gin.Engine

func init() {
	router = initRouter()
}

func makeAPICall(httpMethod string, route string, data url.Values) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	req, err := http.NewRequest(httpMethod, route, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	return w
}

func TestRouter(t *testing.T) {
	t.Run("test_health", func(t *testing.T) {
		w := makeAPICall("GET", "/", nil)

		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.String(), "pong")
	})

	t.Run("test_valid_input", func(t *testing.T) {
		formData := url.Values{}
		formData.Add("password", "securestring")

		w := makeAPICall("POST", "/hash?timer=false", formData)
		assert.Equal(t, w.Code, http.StatusOK)

		result := w.Body.String()
		actual, err := base64.StdEncoding.DecodeString(result)
		if err != nil {
			t.Fatal(err)
		}

		err = bcrypt.CompareHashAndPassword(actual, []byte("securestring"))

		assert.Equal(t, nil, err)
	})

	t.Run("test_valid_input_with_alg_choice", func(t *testing.T) {
		formData := url.Values{}
		formData.Add("password", "securestring")

		w := makeAPICall("POST", "/hash?alg=SHA512&timer=false", formData)

		expected := "8qZk9CXChWG3k63kB2L3Iwl8vPXgpK99lgvebvOXxfyoT1J9SCnPzBxUorEYZsAe+vqArWdOAMChEZR3ng6jOw=="
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.String(), expected)
	})

	t.Run("test_missing_input_value", func(t *testing.T) {
		formData := url.Values{}
		formData.Add("password", "")

		w := makeAPICall("POST", "/hash", formData)

		expected := "no password provided in request"
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), expected)
	})

	t.Run("test_missing_input_parameter", func(t *testing.T) {
		formData := url.Values{}
		formData.Add("SOMERANDOMPARAMETER", "securestring")

		w := makeAPICall("POST", "/hash", formData)

		expected := "no password provided in request"
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), expected)
	})

	t.Run("test_non_existent_endpoint", func(t *testing.T) {
		w := makeAPICall("GET", "/SOMERANDOMENDPOINT", nil)

		assert.Equal(t, w.Code, http.StatusNotFound)
	})

	t.Run("test_invalid_method", func(t *testing.T) {
		w := makeAPICall("POST", "/", nil)

		assert.Equal(t, w.Code, http.StatusNotFound)
	})
}
