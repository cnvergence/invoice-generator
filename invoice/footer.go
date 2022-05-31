package invoice

import (
	"fmt"
	"strconv"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// buildFooter prepares footer on the invoice.
func (i *Invoice) buildFooter() {
	i.pdf.RegisterFooter(func() {
		i.pdf.SetAliasNbPages("{nbs}")
		currentPage := strconv.Itoa(i.pdf.GetCurrentPage())
		i.pdf.Row(6, func() {
			i.pdf.Col(12, func() {
				i.pdf.Text(fmt.Sprintf("Page %s of {nbs}", currentPage), props.Text{
					Top:   1,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: getTealColor(),
				})
			})
		})
		i.pdf.Row(6, func() {
			i.pdf.Col(12, func() {
				i.pdf.Text("github.com/cnvergence/invoice-generator", props.Text{
					Top:   1,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: getTealColor(),
				})
			})
		})
	})
}
