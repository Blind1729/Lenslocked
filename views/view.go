package views

import "html/template"

func NewView(files ...string) *View {
	// We append footer view to every page
	files = append(files, "views/layouts/footer.gohtml")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
