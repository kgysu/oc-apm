package htmlapp

import (
	"fmt"
	htmlTemplates "github.com/kgysu/oc-apm/web/html/templates"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
	"github.com/kgysu/oc-wrapper/appitem"
	"github.com/kgysu/oc-wrapper/application"
	"strconv"
)

func CreateAppDetailsHtml(app *application.Application) string {
	return htmlTemplates.CreateHtml("App", fmt.Sprint(
		createProjcetMetadataTable(app),
		htmlTemplates.CreateBreaks(2),
		htmlTemplates.CreateHeader(2, "Items"),
		htmlTemplates.CreateBreaks(1),
		createItemsTable(app),
		htmlTemplates.CreateBreaks(2),
	))
}

func createProjcetMetadataTable(app *application.Application) string {
	return htmlTemplates.CreateTable("table table-bordered", fmt.Sprint(
		htmlTemplates.CreateTableBody(fmt.Sprint(
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("Name"),
				htmlTemplates.CreateTableData(app.Name),
			)),
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("Label"),
				htmlTemplates.CreateTableData(app.Label),
			)),
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("Items"),
				htmlTemplates.CreateTableData(fmt.Sprintf("%d", len(app.Items))),
			)),
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("Controls"),
				htmlTemplates.CreateTableData(CreateAppControls(app, "")),
			)),
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("Actions"),
				htmlTemplates.CreateTableData(CreateAppActions(app, "")),
			)),
		)),
	))
}

func createItemsTable(app *application.Application) string {
	// Header
	header := htmlTemplates.CreateTableHeader("thead-light", fmt.Sprint(
		htmlTemplates.CreateTableHeadData("Nr."),
		htmlTemplates.CreateTableHeadData("Kind"),
		htmlTemplates.CreateTableHeadData("Name"),
		htmlTemplates.CreateTableHeadData("Details"),
		htmlTemplates.CreateTableHeadData("Controls"),
		htmlTemplates.CreateTableHeadData("Actions"),
	))
	// Items
	var itemsRows string
	for i, item := range app.Items {
		appRow := htmlTemplates.CreateTableRow(
			fmt.Sprint(
				htmlTemplates.CreateTableData(strconv.FormatInt(int64(i+1), 10)),
				htmlTemplates.CreateTableData(item.GetKind()),
				htmlTemplates.CreateTableData(item.GetName()),
				htmlTemplates.CreateTableData(
					htmlTemplates.CreateLinkButton("btn btn-outline-secondary", urls.CreateAppItemDetailsPageUrl(app.Name, item.GetKind(), item.GetName()), "Details")),
				htmlTemplates.CreateTableData(createAppItemControls(app, item)),
				htmlTemplates.CreateTableData(createAppItemActions(app, item)),
			))
		itemsRows = itemsRows + appRow
	}

	return htmlTemplates.CreateTable("table table-striped",
		fmt.Sprint(header, htmlTemplates.CreateTableBody(itemsRows)))
}

func createAppItemControls(app *application.Application, item appitem.AppItem) string {
	content := ""
	if item.IsScalable() {
		content = fmt.Sprint(
			htmlTemplates.CreateLinkButton("btn btn-sm btn-outline-success",
				urls.CreateCmdAppItemStartUrl(app.Name, item.GetKind(), item.GetName()), "START"),
			" ",
			htmlTemplates.CreateLinkButton("btn btn-sm btn-outline-danger",
				urls.CreateCmdAppItemStopUrl(app.Name, item.GetKind(), item.GetName()), "STOP"),
			" ",
		)
	}
	return content
}

func createAppItemActions(app *application.Application, item appitem.AppItem) string {
	return fmt.Sprint(
		htmlTemplates.CreateLinkButton("btn btn-sm btn-outline-success",
			urls.CreateCmdAppItemCreateUrl(app.Name, item.GetKind(), item.GetName()), "Create"),
		" ",
		htmlTemplates.CreateLinkButton("btn btn-sm btn-outline-primary",
			urls.CreateCmdAppItemUpdateUrl(app.Name, item.GetKind(), item.GetName()), "Update"),
		" ",
		htmlTemplates.CreateLinkButton("btn btn-sm btn-outline-danger",
			urls.CreateCmdAppItemDeleteUrl(app.Name, item.GetKind(), item.GetName()), "Delete"),
	)
}
