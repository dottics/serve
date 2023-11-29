package openapi

import (
	"embed"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"strings"
)

//go:embed index.html
var f embed.FS

// getYAMLtoJSON reads the doc specifications and returns the specification as
// JSON in a slice of bytes.
func getYAMLtoJSON() ([]byte, error) {
	// read the openapi yaml file
	xb, err := os.ReadFile("./openapi.yaml")
	if err != nil {
		return xb, err
	}

	// unmarshal the yaml file
	var yml = make(map[string]any)
	err = yaml.Unmarshal(xb, yml)
	if err != nil {
		return xb, err
	}

	// marshal the yaml file to json
	xb, err = json.Marshal(yml)
	if err != nil {
		return xb, err
	}
	return xb, err
}

func getHTML() []byte {
	data, _ := f.ReadFile("index.html")
	return data
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var xb []byte
	var err error

	if strings.HasSuffix(r.RequestURI, ".yml") {
		xb, err = getYAMLtoJSON()
		w.Header().Set("content-type", "application/json")
		if err != nil {
			w.WriteHeader(403)
			msg := fmt.Sprintf(`{"detail": "%s"}`, err.Error())
			_, _ = w.Write([]byte(msg))
			return
		}
	} else {
		// we assume it is the html
		xb = getHTML()
		w.Header().Set("content-type", "text/html")
	}

	w.WriteHeader(200)
	_, _ = w.Write(xb)
}
