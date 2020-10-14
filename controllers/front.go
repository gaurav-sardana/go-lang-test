package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc1 := newUserController()

	http.Handle("/users", *uc1)
	http.Handle("/users/", *uc1)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
