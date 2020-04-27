package htmllist

import (
	"fmt"
	htmlTemplates "github.com/kgysu/oc-apm/web/html/templates"
	"github.com/kgysu/oc-apm/web/server/handlers/urls"
	"github.com/kgysu/oc-wrapper/appitem"
)

func CreateItemsListHtml(items []appitem.AppItem, labelSelector string, kinds string) string {
	return htmlTemplates.CreateHtml("List", fmt.Sprint(
		htmlTemplates.CreateBreaks(1),
		createParamsForm(labelSelector, kinds),
		htmlTemplates.CreateBreaks(5),
		CreateItemsAccordion(items),
		htmlTemplates.CreateBreaks(5),
	))
}

func createParamsForm(labelSelector string, kinds string) string {
	return htmlTemplates.CreateForm(urls.OcCmdList, "get", "List", fmt.Sprint(
		htmlTemplates.CreateInputFormGroupRow("form-group row", "label", "Label Selector", fmt.Sprint(
			htmlTemplates.CreateInput("text", "form-control", "label", labelSelector),
			htmlTemplates.CreateInputDescription("label", "Label Selector or empty for all."),
		)),
		htmlTemplates.CreateInputFormGroupRow("form-group row", "Kind-List", "Kind", fmt.Sprint(
			htmlTemplates.CreateInput("text", "form-control", "kinds", kinds),
			htmlTemplates.CreateInputDescription("label", "Filter by kind or empty for all. "+
				"Multiple values allowed with any char as separator. Example: Pod, DeploymentConfig, Service, Route"),
		))))
}

func CreateItemsAccordion(items []appitem.AppItem) string {
	itemCards := ""
	for i, item := range items {
		itemContent, err := item.ToYaml()
		if err != nil {
			itemContent = err.Error()
		}
		itemCards = itemCards +
			htmlTemplates.CreateCard(
				fmt.Sprintf("item%d", i),
				"items",
				fmt.Sprintf("itemCard%d", i),
				item.InfoStatusHtml(),
				CreateItemContent(itemContent))
	}
	return htmlTemplates.CreateAccordion("items", itemCards)
}

func CreateItemContent(content string) string {
	//return htmlTemplates.CreateDiv("yaml-content",
	//	htmlTemplates.CreateCodeBlock(content))
	return htmlTemplates.CreateCodeBlock(content)
}
