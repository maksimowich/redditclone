package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
)

func ValidateRequest(r *http.Request, dst any) error {
	contentType := r.Header.Get("Content-Type")

	switch {
	case contentType == "application/json":
		if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
			return err
		}
	default:
		if err := r.ParseForm(); err != nil {
			return err
		}
		decoder := schema.NewDecoder()
		if err := decoder.Decode(dst, r.PostForm); err != nil {
			return err
		}
	}

	_, err := govalidator.ValidateStruct(dst)
	return err
}
