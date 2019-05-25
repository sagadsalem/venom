package utils

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	// Tpl is template singleton
	Tpl *template.Template
	// err for handling errors
	err error
)

func init() {
	Tpl, err = FindAndParseTemplates("templates")
	if err != nil {
		panic(err.Error())
	}
}

// Message function to create the response struct
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond function to create the response JSON
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		os.Exit(0)
	}
}

// FindAndParseTemplates function parse all the templates
func FindAndParseTemplates(rootDir string) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name)
			t, e2 = t.Funcs(template.FuncMap{
				//you can parse all your functions here
				"APP_NAME": func() string {
					return os.Getenv("APP_NAME")
				},
			}).Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}
