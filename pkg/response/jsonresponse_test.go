package response_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cooperaj/starling-coinjar/pkg/response"
)

var _ = Describe("JsonResponse", func() {
	type Response struct {
		TestValue string `json:"test_value"`
	}

	var handler = &JsonResponse{
		Body: &Response{
			TestValue: "Test",
		},
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

			Expect(rr.Code).To(Equal(http.StatusOK))
		})

		It("should have a body that is correctly formatted json", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			bodyBytes, _ := ioutil.ReadAll(rr.Body)
			Expect(string(bodyBytes)).Should(MatchJSON(`{"test_value":"Test"}`))
		})
	})
})
