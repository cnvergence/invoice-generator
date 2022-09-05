package invoice

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// buildHeader prepares header on the invoice.
func (i *Invoice) buildHeader() {
	i.pdf.RegisterHeader(func() {
		i.pdf.Row(30, func() {
			i.pdf.Col(5, func() {
				i.pdf.Text("Invoice", props.Text{
					Size:  30,
					Style: consts.Bold,
					Align: consts.Left,
				})
				i.pdf.Text(i.Number, props.Text{
					Top:   12,
					Size:  30,
					Style: consts.Bold,
				})
			})
			i.pdf.ColSpace(3)
			i.pdf.Col(4, func() {
				i.pdf.Text("Date of issue:", props.Text{
					Size:  8,
					Style: consts.Bold,
					Align: consts.Left,
					Color: getTealColor(),
				})
				i.pdf.Text(i.IssueDate, props.Text{
					Size:  8,
					Style: consts.Bold,
					Align: consts.Center,
				})
				i.pdf.Text("Date of sale:", props.Text{
					Top:   12,
					Size:  8,
					Style: consts.Bold,
					Color: getTealColor(),
				})
				i.pdf.Text(i.SaleDate, props.Text{
					Top:   12,
					Size:  8,
					Style: consts.Bold,
					Align: consts.Center,
				})
				i.pdf.Text("Due date:", props.Text{
					Top:   24,
					Size:  8,
					Style: consts.Bold,
					Color: getTealColor(),
				})
				i.pdf.Text(i.DueDate, props.Text{
					Top:   24,
					Size:  8,
					Style: consts.Bold,
					Align: consts.Center,
				})
			})
		})
	})
}
