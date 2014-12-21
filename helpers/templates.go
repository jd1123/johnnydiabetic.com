package helpers

import (
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/jd1123/johnnydiabetic.com/config"
)

func Template(name string) string {
	return path.Join(config.CONFIG["templateDir"], name)
}

func RunTemplate(w http.ResponseWriter, templateName, baseTemplateName string, data interface{}) {
	tmpl := template.Must(template.New(templateName).ParseFiles(baseTemplateName, templateName))
	fmt.Println("Executing a template")
	tmpl.ExecuteTemplate(w, "base", data)
}
