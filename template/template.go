package template

import (
	"html/template"

	"model-admin/page"
)

type Template interface {
	TableTemplate(table *page.Table) template.HTML
	TableColumnTemplate(column *page.TableColumn) template.HTML
}
