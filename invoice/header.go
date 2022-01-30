package invoice

import (
	"fmt"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

func getCurrentDate() string {
	t := time.Now().Format("01/2006")
	return t
}

func (i *Invoice) BuildHeader() {
	i.pdf.RegisterHeader(func() {
		i.pdf.Row(30, func() {
			i.pdf.Col(5, func() {
				i.pdf.Text("Invoice", props.Text{
					Size:  30,
					Style: consts.Bold,
					Align: consts.Left,
				})
				i.pdf.Text(fmt.Sprintf("%s/%s", i.Number, getCurrentDate()), props.Text{
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
				i.pdf.Text("Date of issue:", props.Text{
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
