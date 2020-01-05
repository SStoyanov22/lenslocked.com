package views
import (
	"html/template"
	"path/filepath"
	"net/http"
)
var (
	LayoutDir string  = "views/layouts/"
	TemplateExt string = ".gohtml"
	TemplateDir string = "views/"
)

func NewView (layout string, files... string) *View{
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t,err := template.ParseFiles(files...)
	if err!= nil{
		panic(err)
	}

	return &View{
		Template : t,
		Layout : layout,
	}
}

type View struct{
	Template *template.Template
	Layout string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if err := v.Render(w, nil); err!=nil{
		panic(err)
	}
}

//Render is used to render the view with the predifined layouy
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html") 
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//returns a slice of strings representing
//the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	return files
}

//Takes in a slice of strings
//representing  file paths for templates, and it prepends
// the TemplateDir directory to each string in the slice

//Eg: the input {"home"} would result in the output
//{"views/home"} if TemplateDir = "views"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

//Takes in a slice of strings
//representing  file paths for templates, and it prepends
// the TemplateExt extension to each string in the slice
//
//Eg: the input {"home"} would result in the output
//{"home.gohtml"} if TemplateExt = ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}