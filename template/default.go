package template

import (
	"html/template"

	"model-admin/page"
)

type DefaultTemplate struct {
}

func (d DefaultTemplate) TableTemplate(table *page.Table) template.HTML {
	// TODO implement me
	panic("implement me")
}

func (d DefaultTemplate) TableColumnTemplate(column *page.TableColumn) template.HTML {
	// TODO implement me
	panic("implement me")
}
