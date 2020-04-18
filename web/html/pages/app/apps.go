package htmlapp

import (
	"fmt"
	htmlTemplates "github.com/kgysu/oc-apm/web/html/templates"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
	"github.com/kgysu/oc-wrapper/application"
	"strconv"
)

func CreateAppListHtml(apps []*application.Application) string {
	return htmlTemplates.CreateHtml("Apps", fmt.Sprint(
		htmlTemplates.CreateLinkButton("btn btn-primary float-right", urls.AppNewPageUrl, "+ New App"),
		htmlTemplates.CreateBreaks(4),
		CreateAppsTable(apps),
	))
}

func CreateAppsTable(apps []*application.Application) string {
	// Header
	header := htmlTemplates.CreateTableHeader("thead-light", fmt.Sprint(
		htmlTemplates.CreateTableHeadData("Nr."),
		htmlTemplates.CreateTableHeadData("Name"),
		htmlTemplates.CreateTableHeadData("Status"),
		htmlTemplates.CreateTableHeadData("Details"),
		htmlTemplates.CreateTableHeadData("Controls"),
		htmlTemplates.CreateTableHeadData("Actions"),
	))
	// Apps
	var appRows string
	for i, app := range apps {
		appRow := htmlTemplates.CreateTableRow(
			fmt.Sprint(
				htmlTemplates.CreateTableData(strconv.FormatInt(int64(i), 10)),
				htmlTemplates.CreateTableData(app.Name),
				htmlTemplates.CreateTableData("unknown"),
				htmlTemplates.CreateTableData(htmlTemplates.CreateLinkButton("btn btn-outline-secondary", "/app/"+app.Name, "Details")),
				htmlTemplates.CreateTableData(CreateAppControls(app, "-outline")),
				htmlTemplates.CreateTableData(CreateAppActions(app, "-outline")),
			))

		appRows = appRows + appRow
	}
	return htmlTemplates.CreateTable("table table-striped",
		fmt.Sprint(header, htmlTemplates.CreateTableBody(appRows)))
}

func CreateAppControls(app *application.Application, outline string) string {
	return fmt.Sprint(
		htmlTemplates.CreateLinkButton("btn btn-sm btn"+outline+"-success", urls.CmdAppStartUrl+app.Name, "START"),
		" ",
		htmlTemplates.CreateLinkButton("btn btn-sm btn"+outline+"-danger", urls.CmdAppStopUrl+app.Name, "STOP"),
		" ",
	)
}

func CreateAppActions(app *application.Application, outline string) string {
	return fmt.Sprint(
		htmlTemplates.CreateLinkButton("btn btn-sm btn"+outline+"-success", urls.CmdAppCreateUrl+app.Name, "Create"),
		" ",
		htmlTemplates.CreateLinkButton("btn btn-sm btn"+outline+"-primary", urls.CmdAppUpdateUrl+app.Name, "Update"),
		" ",
		htmlTemplates.CreateLinkButton("btn btn-sm btn"+outline+"-danger", urls.CmdAppDeleteUrl+app.Name, "Delete"),
	)
}
