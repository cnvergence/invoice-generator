package invoice

import (
	"fmt"
	"strings"

	maroto "github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"gopkg.in/yaml.v3"
)

// New returns Invoice struct loaded with values from YAML and prepares PDF struct.
func New(file []byte) (*Invoice, error) {
	invoice := &Invoice{}
	if err := yaml.Unmarshal(file, &invoice); err != nil {
		return nil, fmt.Errorf("could not unmarshal yaml values: %s", err)
	}

	// Auto-split address into addressLine1/addressLine2 when not explicitly set.
	splitAddress(&invoice.Company.Seller.AddressLine1, &invoice.Company.Seller.AddressLine2, invoice.Company.Seller.Address)
	splitAddress(&invoice.Company.Buyer.AddressLine1, &invoice.Company.Buyer.AddressLine2, invoice.Company.Buyer.Address)

	cfgBuilder := config.NewBuilder().
		WithPageSize(pagesize.A4).
		WithLeftMargin(10).
		WithTopMargin(15).
		WithRightMargin(10).
		WithPageNumber(props.PageNumber{
			Pattern: "Page {current} of {total}",
			Place:   props.LeftBottom,
			Style:   fontstyle.BoldItalic,
			Size:    8,
			Color:   getTealColor(),
		})

	var err error
	cfgBuilder, err = invoice.configureFonts(cfgBuilder)
	if err != nil {
		return nil, fmt.Errorf("could not configure fonts: %s", err)
	}

	invoice.pdf = maroto.New(cfgBuilder.Build())
	if err := invoice.setPDFLayout(); err != nil {
		return nil, fmt.Errorf("could not set the invoice layout: %s", err)
	}

	return invoice, nil
}

// splitAddress populates line1/line2 from a single address string (split on
// comma) when they are not already set.
func splitAddress(line1, line2 *string, address string) {
	if *line1 != "" || address == "" {
		return
	}
	parts := strings.SplitN(address, ",", 2)
	*line1 = strings.TrimSpace(parts[0])
	if len(parts) == 2 && *line2 == "" {
		*line2 = strings.TrimSpace(parts[1])
	}
}

func getTealColor() *props.Color {
	return &props.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getWhiteColor() *props.Color {
	return &props.Color{
		Red:   255,
		Green: 255,
		Blue:  255,
	}
}

func (i *Invoice) setPDFLayout() error {
	if err := i.buildHeader(); err != nil {
		return fmt.Errorf("could not build header: %s", err)
	}
	if err := i.buildFooter(); err != nil {
		return fmt.Errorf("could not build footer: %s", err)
	}
	i.buildCompanyDetails()
	i.buildBankDetails()
	i.buildTable()
	i.buildSignature()
	return nil
}

// SaveToPdf saves Invoice to a PDF file.
func (i *Invoice) SaveToPdf(outputPath string) error {
	doc, err := i.pdf.Generate()
	if err != nil {
		return fmt.Errorf("could not generate Invoice: %s", err)
	}
	if err := doc.Save(outputPath); err != nil {
		return fmt.Errorf("could not save Invoice to .pdf file: %s", err)
	}
	return nil
}

// SaveAsBytes saves Invoice to bytes.
func (i *Invoice) SaveAsBytes() ([]byte, error) {
	doc, err := i.pdf.Generate()
	if err != nil {
		return nil, fmt.Errorf("could not generate Invoice: %s", err)
	}
	return doc.GetBytes(), nil
}
