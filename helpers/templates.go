package helpers

import (
	"net/http"
	"path"
	"text/template"

	"github.com/jd1123/johnnydiabetic.com/config"
)

func Template(name string) string {
	return path.Join(config.CONFIG["templateDir"], name)
}

func RunTemplate(w http.ResponseWriter, templateName, baseTemplateName string, data interface{}) {
	tp := Template(templateName)
	bstp := Template(baseTemplateName)
	tmpl := template.Must(template.New(tp).ParseFiles(bstp, tp))
	tmpl.ExecuteTemplate(w, "base", data)
}
