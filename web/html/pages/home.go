package htmlPages

import htmlTemplates "github.com/kgysu/oc-apm/web/html/templates"

// Home/Main Page
func CreateHomePage() string {
	return htmlTemplates.CreateHtml("Home", "main page")
}
