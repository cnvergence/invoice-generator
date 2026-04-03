package invoice

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// buildCompanyDetails prepares rows with Buyer and Seller contact details on the invoice.
func (i *Invoice) buildCompanyDetails() {
	i.pdf.AddRows(
		row.New(7).WithStyle(&props.Cell{BackgroundColor: getTealColor()}).Add(
			col.New(3).Add(
				text.New("Seller", props.Text{
					Top:   1.5,
					Size:  9,
					Style: fontstyle.Bold,
					Align: align.Center,
					Color: getWhiteColor(),
				}),
			),
			col.New(4),
			col.New(5).Add(
				text.New("Buyer", props.Text{
					Top:   1.5,
					Size:  9,
					Style: fontstyle.Bold,
					Align: align.Center,
					Color: getWhiteColor(),
				}),
			),
		),
		row.New(10).Add(
			col.New(2).Add(
				text.New("Name:  ", props.Text{
					Top:   2,
					Style: fontstyle.Bold,
					Align: align.Left,
					Color: getTealColor(),
				}),
			),
			col.New(3).Add(
				text.New(i.Company.Seller.Name, props.Text{
					Top:   2,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
			),
			col.New(2),
			col.New(2).Add(
				text.New("Name:  ", props.Text{
					Top:   2,
					Style: fontstyle.Bold,
					Align: align.Left,
					Color: getTealColor(),
				}),
			),
			col.New(3).Add(
				text.New(i.Company.Buyer.Name, props.Text{
					Top:   2,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
			),
		),
		row.New(10).Add(
			col.New(2).Add(
				text.New("Address:  ", props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Align: align.Left,
					Color: getTealColor(),
				}),
			),
			col.New(3).Add(
				text.New(i.Company.Seller.Address, props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
			),
			col.New(2),
			col.New(2).Add(
				text.New("Address:  ", props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Align: align.Left,
					Color: getTealColor(),
				}),
			),
			col.New(3).Add(
				text.New(i.Company.Buyer.Address, props.Text{
					Top:   2,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
			),
		),
		row.New(7).Add(
			col.New(2).Add(
				text.New("VAT Number:  ", props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Align: align.Left,
					Color: getTealColor(),
				}),
			),
			col.New(3).Add(
				text.New(i.Company.Seller.VAT, props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
			),
			col.New(2),
			col.New(2).Add(
				text.New("VAT Number:  ", props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Align: align.Left,
					Color: getTealColor(),
				}),
			),
			col.New(3).Add(
				text.New(i.Company.Buyer.VAT, props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
			),
		),
		row.New(2).Add(col.New(12)),
	)
}
