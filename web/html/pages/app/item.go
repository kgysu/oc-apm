package htmlapp

import (
	"fmt"
	htmlTemplates "github.com/kgysu/oc-apm/web/html/templates"
	"github.com/kgysu/oc-wrapper/appitem"
	"github.com/kgysu/oc-wrapper/application"
)

func CreateAppItemDetailsHtml(app *application.Application, item appitem.AppItem) string {
	return htmlTemplates.CreateHtml("Item", fmt.Sprint(
		CreateItemMetadataTable(app, item),
		htmlTemplates.CreateBreaks(2),
		CreateItemDetails(item),
		htmlTemplates.CreateBreaks(3),
	))
}

func CreateItemMetadataTable(app *application.Application, item appitem.AppItem) string {
	return htmlTemplates.CreateTable("table table-bordered", fmt.Sprint(
		htmlTemplates.CreateTableBody(fmt.Sprint(
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("Name"),
				htmlTemplates.CreateTableData(item.GetName()),
			)),
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("Kind"),
				htmlTemplates.CreateTableData(item.GetKind()),
			)),
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("File"),
				htmlTemplates.CreateTableData(item.GetFileName()),
			)),
			htmlTemplates.CreateTableRow(fmt.Sprint(
				htmlTemplates.CreateTableHeadData("Actions"),
				htmlTemplates.CreateTableData(createAppItemActions(app, item)),
			)),
		))))
}

func CreateItemDetails(item appitem.AppItem) string {
	itemContent, err := item.ToYaml()
	if err != nil {
		itemContent = err.Error()
	}
	return fmt.Sprint(
		htmlTemplates.CreateHeader(3, fmt.Sprintf("%s (%s)", item.GetName(), item.GetFileName())),
		htmlTemplates.CreateBreaks(1),
		htmlTemplates.CreateDiv("yaml-content",
			htmlTemplates.CreateCodeBlock(itemContent)),
		//fmt.Sprint(
		//"form-control-plaintext",
		//"id"+item.GetName(),
		//"content",
		//"20",
		//"readonly",
		//itemContent,
		//)),
	)
}
