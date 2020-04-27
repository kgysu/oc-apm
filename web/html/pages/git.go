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
		htmlTemplates.CreateInputFormGroupRow("form-group row", "repoUrl", "Repo-URL", fmt.Sprint(
			htmlTemplates.CreateInput("text", "form-control", "repoUrl", ""),
			htmlTemplates.CreateInputDescription("repoUrl", "URL to repository, example: https://github.com/user/repo"),
		)),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "subPath", "Repo sub directory", fmt.Sprint(
			htmlTemplates.CreateInput("text", "form-control", "subPath", ""),
			htmlTemplates.CreateInputDescription("subPath", "Sub directory to clone from repository. Or none to clone all. (slash is required) Example: /apps"),
		)),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "tag", "Git Tag", fmt.Sprint(
			htmlTemplates.CreateInput("text", "form-control", "tag", ""),
			htmlTemplates.CreateInputDescription("tag", "Git Repository Tag to clone. Example: v1.0.0"),
		)),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "branch", "Git Branch", fmt.Sprint(
			htmlTemplates.CreateInput("text", "form-control", "branch", ""),
			htmlTemplates.CreateInputDescription("branch", "Git Repository Branch to clone. ONLY IF TAG IS EMPTY! Example: feature/one"),
		))))

}
