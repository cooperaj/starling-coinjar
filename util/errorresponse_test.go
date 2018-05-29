package util_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cooperaj/starling-coinjar/util"
)

var _ = Describe("ErrorResponse", func() {
	var handler = &ErrorResponse{
		Message: "An error message",
		Code:    http.StatusExpectationFailed,
	}

	Context("When returning a JSON response", func() {
		It("should have an appropriate mimetype", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			Expect(rr.HeaderMap.Get("content-type")).To(Equal("application/json"))
		})

		It("should be a 200 response code", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusExpectationFailed))
		})

		It("should have a body that is correctly formatted json", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			bodyBytes, _ := ioutil.ReadAll(rr.Body)
			Expect(string(bodyBytes)).Should(MatchJSON(`{"error":"An error message"}`))
		})
	})
})
