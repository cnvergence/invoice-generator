package invoice

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// buildCompanyDetails prepares rows with Buyer and Seller contact details on the invoice.
func (i *Invoice) buildCompanyDetails() {
	i.pdf.Row(7, func() {
		i.pdf.SetBackgroundColor(getTealColor())
		i.pdf.Col(3, func() {
			i.pdf.Text("Seller", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.NewWhite(),
			})
		})
		i.pdf.ColSpace(3)
		i.pdf.Col(3, func() {
			i.pdf.Text("Buyer", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Right,
				Color: color.NewWhite(),
			})
		})
		i.pdf.ColSpace(3)
	})

	i.pdf.SetBackgroundColor(color.NewWhite())
	i.pdf.Row(10, func() {
		i.pdf.Col(3, func() {
			i.pdf.Text("Name:  ", props.Text{
				Top:   2,
				Style: consts.Bold,
				Align: consts.Right,
				Color: getTealColor(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Seller.Name, props.Text{
				Top:   2,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text("Name:  ", props.Text{
				Top:   2,
				Style: consts.Bold,
				Align: consts.Right,
				Color: getTealColor(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Buyer.Name, props.Text{
				Top:   2,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
	})
	i.pdf.Row(10, func() {
		i.pdf.Col(3, func() {
			i.pdf.Text("Address:  ", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Right,
				Color: getTealColor(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Seller.Address, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text("Address:  ", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Right,
				Color: getTealColor(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Buyer.Address, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
	})
	i.pdf.Row(7, func() {
		i.pdf.Col(3, func() {
			i.pdf.Text("VAT Number:  ", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Right,
				Color: getTealColor(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Seller.VAT, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text("VAT Number:  ", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Right,
				Color: getTealColor(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Buyer.VAT, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
	})
	i.pdf.Row(2, func() {
	})

}
