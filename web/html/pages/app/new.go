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
		htmlTemplates.CreateInputFormGroupRow("form-group row", "name", "Name", fmt.Sprint(
			htmlTemplates.CreateInput("text", "form-control", "name", ""),
			htmlTemplates.CreateInputDescription("label", "App name. Apps with the same name are overwritten if they already exist!"),
		)),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "label", "Label-Selector", fmt.Sprint(
			htmlTemplates.CreateInput("text", "form-control", "label", "app="),
			htmlTemplates.CreateInputDescription("label", "Label Selector or empty for all."),
		)),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "from", "From Existing",
			htmlTemplates.CreateCheckBox("from", "Create from existing openshift app")),
	))
}
