package views

import (
	"html/template"
	"path/filepath"
)

const (
	LayoutDir         = "views/layouts/"
	TemplateExtension = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	layoutFiles := layoutFiles()

	for file := range layoutFiles {
		files = append(files, layoutFiles[file])
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

// Get all the layout files present under the layouts dir
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExtension)
	if err != nil {
		panic(err)
	}
	return files
}
