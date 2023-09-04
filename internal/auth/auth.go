package auth

import "net/http"

// Extracts api key from the header of http request
// Example -> Authorization {actual api key}

func GetAPIKey(headers http.Header) (string, error) {

}
