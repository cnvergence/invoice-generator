package invoice

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

//buildSignature prepares signatures of the receiver and issuer.
func (i *Invoice) buildSignature() {
	i.pdf.SetBackgroundColor(getTealColor())
	i.pdf.Line(0.5)
	i.pdf.SetBackgroundColor(color.NewWhite())

	i.pdf.Row(15, func() {
		i.pdf.Col(1, func() {
			i.pdf.Text("Notes:", props.Text{
				Top:   1,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Left,
				Color: getTealColor(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Notes, props.Text{
				Top:   1,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Left,
			})
		})
	})

	i.pdf.Row(15, func() {
		i.pdf.Col(6, func() {
			i.pdf.Signature("Signature of the receiver", props.Font{
				Size:  12.0,
				Style: consts.BoldItalic,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
			})
		})
		i.pdf.ColSpace(3)
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Signature, props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
			i.pdf.Signature("Signature of the issuer", props.Font{
				Size:  12.0,
				Style: consts.BoldItalic,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
			})
		})
	})

}
