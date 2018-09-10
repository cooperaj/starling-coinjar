package middleware

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/cooperaj/starling-coinjar/pkg/response"
)

// SignatureValidationMiddleware Provides the webhook authentication that Starling uses
type SignatureValidationMiddleware struct {
	Secret string
}

// Middleware The middleware handler
func (s *SignatureValidationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hash := r.Header.Get("X-Hook-Signature")
		if hash == "" {
			response := response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "No signature specified",
			}
			response.ServeHTTP(w, r)
			return
		}

		// read and restore the request body since the ReadAll call drains the buffer
		body, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// encode our own version of the signature so we can check it against the
		// provided one.
		shaHash := sha512.New()
		shaHash.Write([]byte(s.Secret + string(body)))

		encodedString := base64.StdEncoding.EncodeToString(shaHash.Sum(nil))

		if hash != encodedString {
			response := response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid signature specified",
			}
			response.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
