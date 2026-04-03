package invoice

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// buildHeader prepares header on the invoice.
func (i *Invoice) buildHeader() error {
	return i.pdf.RegisterHeader(
		row.New(30).Add(
			col.New(5).Add(
				text.New("Invoice", props.Text{
					Size:  30,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
				text.New(i.Number, props.Text{
					Top:   12,
					Size:  30,
					Style: fontstyle.Bold,
				}),
			),
			col.New(3),
			col.New(4).Add(
				text.New("Date of issue:", props.Text{
					Size:  8,
					Style: fontstyle.Bold,
					Align: align.Left,
					Color: getTealColor(),
				}),
				text.New(i.IssueDate, props.Text{
					Size:  8,
					Style: fontstyle.Bold,
					Align: align.Center,
				}),
				text.New("Date of sale:", props.Text{
					Top:   12,
					Size:  8,
					Style: fontstyle.Bold,
					Color: getTealColor(),
				}),
				text.New(i.SaleDate, props.Text{
					Top:   12,
					Size:  8,
					Style: fontstyle.Bold,
					Align: align.Center,
				}),
				text.New("Due date:", props.Text{
					Top:   24,
					Size:  8,
					Style: fontstyle.Bold,
					Color: getTealColor(),
				}),
				text.New(i.DueDate, props.Text{
					Top:   24,
					Size:  8,
					Style: fontstyle.Bold,
					Align: align.Center,
				}),
			),
		),
	)
}
