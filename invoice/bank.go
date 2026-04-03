package invoice

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// buildBankDetails prepares rows with Bank details on the invoice.
func (i *Invoice) buildBankDetails() {
	i.pdf.AddRows(
		row.New(0.5).WithStyle(&props.Cell{BackgroundColor: getTealColor()}).Add(col.New(12)),
		row.New(20).Add(
			col.New(3).Add(
				text.New("Account no:", props.Text{
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
					Color: getTealColor(),
				}),
				text.New(i.Bank.AccountNumber, props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
				}),
			),
			col.New(2).Add(
				text.New("Bank/SWIFT: ", props.Text{
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
					Color: getTealColor(),
				}),
				text.New(i.Bank.Swift, props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
				}),
			),
		),
	)
}
