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

func RunTemplate(w http.ResponseWriter, templateName, headerTemplateName, footerTemplateName string, data interface{}) {
	tp := Template(templateName)
	bstp := Template(headerTemplateName)
	fttp := Template(footerTemplateName)
	tmpl := template.Must(template.ParseFiles(tp, bstp, fttp))
	//	log.Println("RunTemplate() called on", tp, "with", bstp, "and", fttp, "with data", data)
	tmpl.Execute(w, data)
}

func RunTemplateBase(w http.ResponseWriter, templateName string, data interface{}) {
	base := Template("base.html")
	content := Template(templateName)
	t := template.New("base.html")
	t = template.Must(t.ParseFiles(base, content))
	t.Execute(w, data)
}
