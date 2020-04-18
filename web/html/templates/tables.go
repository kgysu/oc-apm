package htmlTemplates

import "fmt"

func CreateTable(class, content string) string {
	return fmt.Sprintf(TableTemplate, class, content)
}

const TableTemplate = `<table class="%s">
%s
</table>
`

func CreateTableHeader(class, content string) string {
	return fmt.Sprintf(TableHeaderTemplate, class, content)
}

const TableHeaderTemplate = `<thead class="%s">
%s
</thead>
`

func CreateTableBody(content string) string {
	return fmt.Sprintf(TableBodyTemplate, content)
}

const TableBodyTemplate = `<tbody>
%s
</tbody>
`

func CreateTableRow(content string) string {
	return fmt.Sprintf(TableRowTemplate, content)
}

const TableRowTemplate = `<tr>
%s
</tr>
`

func CreateTableHeadData(content string) string {
	return fmt.Sprintf(TableHeadDataTemplate, content)
}

const TableHeadDataTemplate = `<th>%s</th>`

func CreateTableData(content string) string {
	return fmt.Sprintf(TableDataTemplate, content)
}

const TableDataTemplate = `<td>%s</td>`
