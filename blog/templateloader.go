package blog

import (
	"html/template"
	"fmt"
	"log"
)

func LoadTemplate(templateName string) *template.Template {	
	completeTemplateName := fmt.Sprintf("templates/%s.html",templateName)
	generatedTemplate, err := template.ParseFiles(completeTemplateName)
	if err != nil {
		log.Fatal("Cant process index template")
	}
	return generatedTemplate
}
