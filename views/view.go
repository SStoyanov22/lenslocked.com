package views
import "http/template"

func NewView (files ...string) *View{
	files = append(files, "views/layouts/footer.gohtml")

	t,err = template.ParseFiles(files...)
	if err!= nil{
		panic(err)
	}

	return &View{
		Template : t,
	}
}

type struct View {
	Template *template.Template
}