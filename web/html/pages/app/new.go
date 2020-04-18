package htmlapp

import (
	"fmt"
	htmlTemplates "github.com/kgysu/oc-apm/web/html/templates"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
)

func CreateAppNewHtml() string {
	return htmlTemplates.CreateHtml("New App", fmt.Sprint(
		CreateNewAppConfigForm(),
		htmlTemplates.CreateBreaks(1),
	))
}

func CreateNewAppConfigForm() string {
	return htmlTemplates.CreateForm(urls.CmdAppNewUrl, "post", "Create", fmt.Sprint(
		htmlTemplates.CreateInputFormGroupRow("form-group row", "name", "Name",
			htmlTemplates.CreateInput("text", "form-control", "name", "")),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "label", "Label-Selector",
			htmlTemplates.CreateInput("text", "form-control", "label", "app=")),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "from", "From Existing",
			htmlTemplates.CreateCheckBox("from", "Create from existing openshift app")),
	))
}
