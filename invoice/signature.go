package invoice

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// buildSignature prepares signatures of the receiver and issuer.
func (i *Invoice) buildSignature() {
	signatureColor := &props.Color{Red: 10, Green: 20, Blue: 30}
	i.pdf.AddRows(
		row.New(0.5).WithStyle(&props.Cell{BackgroundColor: getTealColor()}).Add(col.New(12)),
		row.New(15).Add(
			col.New(1).Add(
				text.New("Notes:", props.Text{
					Top:   1,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
					Color: getTealColor(),
				}),
			),
			col.New(3).Add(
				text.New(i.Notes, props.Text{
					Top:   1,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
				}),
			),
		),
		row.New(15).Add(
			col.New(6).Add(
				signature.New("Signature of the receiver", props.Signature{
					FontSize:  12.0,
					FontStyle: fontstyle.BoldItalic,
					FontColor: signatureColor,
				}),
			),
			col.New(3),
			col.New(3).Add(
				text.New(i.Signature, props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Center,
				}),
				signature.New("Signature of the issuer", props.Signature{
					FontSize:  12.0,
					FontStyle: fontstyle.BoldItalic,
					FontColor: signatureColor,
				}),
			),
		),
	)
}
