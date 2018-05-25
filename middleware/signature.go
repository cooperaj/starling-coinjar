package middleware

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"net/http"

	"github.com/cooperaj/starling-coinjar/util"
)

type SignatureValidationMiddleware struct {
	Secret string
}

func (s *SignatureValidationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hash := r.Header.Get("X-Hook-Signature")
		if hash == "" {
			response := util.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "No signature specified",
			}
			response.ServeHTTP(w, r)
			return
		}

		// encode our own version of the signature so we can check it against the
		// provided one.
		var b bytes.Buffer
		b.WriteString(s.Secret)
		b.ReadFrom(r.Body)

		shaHash := sha512.New()
		shaHash.Write(b.Bytes())

		encodedString := base64.StdEncoding.EncodeToString(shaHash.Sum(nil))

		if hash != encodedString {
			response := util.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid signature specified",
			}
			response.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
