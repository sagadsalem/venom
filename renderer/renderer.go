package renderer

import (
	"github.com/unrolled/render"
	"io"
	"os"
	"projects/starter/helpers"
)

var (
	Rnd *render.Render
)

func init() {
	//init the renderer
	Rnd = NewRenderer("templates", "layouts/base", Development())
}

// NewRenderer instance
func NewRenderer(templatesPath string, layoutPath string, development bool) *render.Render {
	return render.New(render.Options{
		Directory:     templatesPath,
		Layout:        layoutPath,
		Extensions:    []string{".html"},
		IsDevelopment: development,
		Funcs:         helpers.AppHelpers(),
	})
}

// Check if We in Development Mode
func Development() bool {
	if os.Getenv("APP_LOCAL") == "dev" {
		return true
	} else {
		return false
	}
}

// HTML RESPONSE
func HTML(w io.Writer, status int, name string, binding interface{}) error {
	if err := Rnd.HTML(w, status, name, binding); err != nil {
		return err
	}
	return nil
}

// JSON RESPONSE
func JSON(w io.Writer, status int, v interface{}) error {
	if err := Rnd.JSON(w, status, v); err != nil {
		return err
	}
	return nil
}
