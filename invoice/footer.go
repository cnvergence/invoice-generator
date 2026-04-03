package invoice

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// buildFooter prepares footer on the invoice.
// Page numbers are handled automatically via config.WithPageNumber.
func (i *Invoice) buildFooter() error {
	return i.pdf.RegisterFooter(
		row.New(6).Add(
			col.New(12).Add(
				text.New("github.com/cnvergence/invoice-generator", props.Text{
					Top:   1,
					Style: fontstyle.BoldItalic,
					Size:  8,
					Align: align.Left,
					Color: getTealColor(),
				}),
			),
		),
	)
}
