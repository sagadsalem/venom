package helpers

import (
	"html/template"
	"os"
)

// HelperFunc all helpers goes here
func AppHelpers() []template.FuncMap {
	var HelperFunc []template.FuncMap
	appName := template.FuncMap{
		"appName": func() string { return os.Getenv("APP_NAME") },
	}
	HelperFunc = append(HelperFunc, appName)
	return HelperFunc
}
