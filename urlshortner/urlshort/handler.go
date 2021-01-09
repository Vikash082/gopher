package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler takes map[string]string as input and forward the Request
// to original destination. Incase, the path is not found in pathToUrls
// it forward those request to default handler which is ServerMux in this
// case
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if redPath, ok := pathToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, redPath, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler take slice of bytes as input which is actually the array of
// path & url. It decodes the byte into slice of pYaml struct and call
// MapHandler for further processing.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

type pYaml struct {
	Path    string `yaml:"path"`
	OrigURL string `yaml:"url"`
}

func parseYaml(inYaml []byte) ([]pYaml, error) {
	var out []pYaml
	err := yaml.Unmarshal(inYaml, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func buildMap(parsedYaml []pYaml) map[string]string {
	m1 := make(map[string]string)
	for _, m := range parsedYaml {
		m1[m.Path] = m.OrigURL
	}
	return m1
}
