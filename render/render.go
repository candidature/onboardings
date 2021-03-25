package render

import (
	"fmt"
	"github.com/candidature/onboardings/config"
	"github.com/candidature/onboardings/models"
	"log"
	"net/http"
)



func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	parsedTemplateCache := config.GetConfig().TemplateCache

	//PrintCache()
	t, ok := parsedTemplateCache[tmpl]
	if !ok {
		log.Fatal("Template not found")
		return
	}
	err := t.Execute(w,templateData)
	if err!= nil {
		fmt.Printf("Error executing template ", err)
		return
	}
}