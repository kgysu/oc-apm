package htmlPages

import (
	"fmt"
	htmlTemplates "github.com/kgysu/oc-apm/web/html/templates"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
)

func CreateGitPageHtml() string {
	return htmlTemplates.CreateHtml("Git", fmt.Sprint(
		htmlTemplates.CreateHeader(3, "Load from Git-Repo"),
		htmlTemplates.CreateBreaks(1),
		createGitForm(),
		htmlTemplates.CreateBreaks(2),
	))
}

func createGitForm() string {
	return htmlTemplates.CreateForm(urls.LoadFromGitUrl, "post", "Load", fmt.Sprint(
		htmlTemplates.CreateInputFormGroupRow("form-group row", "repoUrl", "Repo-URL",
			htmlTemplates.CreateInput("text", "form-control", "repoUrl", "")),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "subPath", "Repo sub directory",
			htmlTemplates.CreateInput("text", "form-control", "subPath", "")),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "tag", "Git Tag",
			htmlTemplates.CreateInput("text", "form-control", "tag", "")),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "branch", "Git Branch",
			htmlTemplates.CreateInput("text", "form-control", "branch", "")),
	))

}
