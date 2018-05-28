package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cooperaj/starling-coinjar/middleware"
)

var _ = Describe("Signature", func() {
	const authenticatedEndpoint = "/transactions"
	const middlewareSecret = "testSecret"

	var (
		signatureMiddleware SignatureValidationMiddleware
	)

	BeforeEach(func() {
		signatureMiddleware = SignatureValidationMiddleware{
			Secret: middlewareSecret,
		}
	})

	Describe("Accessing an endpoint that requires authentication", func() {
		Context("Without providing a signature", func() {
			It("should fail with a 401 response", func() {
				req := httptest.NewRequest(http.MethodPost, authenticatedEndpoint, nil)
				rr := httptest.NewRecorder()

				handler := signatureMiddleware.Middleware(GetTestHandler())
				handler.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(401))
			})
		})

		Context("Providing an incorrect signature", func() {
			It("should fail with a 401 response", func() {
				body := []byte("{}")

				req := httptest.NewRequest("POST", authenticatedEndpoint, bytes.NewBuffer(body))
				req.Header.Set("X-Hook-Signature", "incorrectSignature")

				rr := httptest.NewRecorder()

				handler := signatureMiddleware.Middleware(GetTestHandler())
				handler.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(401))
			})
		})

		Context("Providing a correct signature", func() {
			It("should succeed with a 200 response", func() {
				body := []byte("{}")

				req := httptest.NewRequest("POST", authenticatedEndpoint, bytes.NewBuffer(body))
				req.Header.Set("X-Hook-Signature", "vwknFNvo7NoRocsxJHwJBE6lqf1v6N650/I/0u1hgwdKj87L4DpdykvWnUP45n2H+VeM3dOsQ97surGvvN5imQ==")

				rr := httptest.NewRecorder()

				handler := signatureMiddleware.Middleware(GetTestHandler())
				handler.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(200))
			})
		})
	})
})
