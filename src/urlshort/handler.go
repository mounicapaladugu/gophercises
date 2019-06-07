package urlshort

import (
	"net/http"
)

//maphandler function
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return nil
}

//yamlhandler - will parse the provided yaml and then return http.Handlerfunc
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}
