package urlshortener

import (
	"net/http"
	"gopkg.in/yaml.v2"
	"encoding/json"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	handlerFn := func (w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if elem, okay := pathsToUrls[path]; okay {
			http.Redirect(w, r, elem, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}

	return handlerFn;
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYml(yml)
	if err != nil {
		return nil, err
	}

	pathsToUrls := createMap(pathUrls)
	
	return MapHandler(pathsToUrls, fallback), nil
}

func JSONHandler(jsonString []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJSON(jsonString)
	if err != nil {
		return nil, err
	}

	pathsToUrls := createMap(pathUrls)
	
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYml(yml []byte) ([]path, error) {
	var pathUrls []path
	if err := yaml.Unmarshal(yml, &pathUrls); err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func parseJSON(jsonString []byte) ([]path, error) {
	var pathUrls []path
	if err := json.Unmarshal(jsonString, &pathUrls); err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func createMap(pathUrls []path) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pathUrl := range pathUrls {
		pathsToUrls[pathUrl.Path] = pathUrl.URL
	}
	return pathsToUrls
}

// Struct fields are only unmarshalled if they are exported
type path struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}