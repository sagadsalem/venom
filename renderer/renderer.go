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

const (
	TEMPLATES = "templates"
	LAYOUT    = "layouts/base"
)

func init() {
	//init the renderer
	Rnd = NewRenderer(TEMPLATES, LAYOUT, Development())
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
func HTML(w io.Writer, status int, name string, binding interface{}, options ...render.HTMLOptions) error {
	if err := Rnd.HTML(w, status, name, binding, prepareHTMLOptions(options)); err != nil {
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

// prepareHTMLOptions
func prepareHTMLOptions(options []render.HTMLOptions) render.HTMLOptions {
	if len(options) > 0 {
		return render.HTMLOptions{
			Layout: options[0].Layout,
		}
	}

	return render.HTMLOptions{Layout: LAYOUT}
}
