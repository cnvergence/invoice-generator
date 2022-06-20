package invoice

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"gopkg.in/yaml.v3"
)

//New returns Invoice struct loaded with values from YAML and prepares PDF struct.
func New(file []byte) (*Invoice, error) {
	invoice := &Invoice{}
	if err := yaml.Unmarshal(file, &invoice); err != nil {
		return nil, fmt.Errorf("could not unmarshal yaml values: %s", err)
	}

	invoice.pdf = pdf.NewMaroto(consts.Portrait, consts.A4)
	err := invoice.setPDFLayout()
	if err != nil {
		return nil, fmt.Errorf("could not set the invoice layout: %s", err)
	}

	return invoice, nil
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func (i *Invoice) setPDFLayout() error {
	i.pdf.SetFirstPageNb(1)
	i.pdf.SetPageMargins(10, 15, 10)
	err := i.setFonts()
	if err != nil {
		return fmt.Errorf("could not configure fonts: %s", err)
	}
	i.buildHeader()
	i.buildFooter()
	i.buildCompanyDetails()
	i.buildBankDetails()
	i.buildTable()
	i.buildSignature()

	_, height := i.pdf.GetPageSize()
	current := i.pdf.GetCurrentOffset()
	filler := height - current - 60
	i.pdf.Row(filler, func() {
	})
	return nil
}

//SaveToPdf saves Invoice to a PDF file and closes it.
func (i *Invoice) SaveToPdf(outputPath string) error {
	err := i.pdf.OutputFileAndClose(outputPath)
	if err != nil {
		return fmt.Errorf("could not save Invoice to .pdf file: %s", err)
	}
	return err
}

//Save saves Invoice to bytes and closes it.
func (i *Invoice) SaveAsBytes() ([]byte, error) {
	bytes, err := i.pdf.Output()
	if err != nil {
		return nil, fmt.Errorf("could not save Invoice to bytes: %s", err)
	}
	return bytes.Bytes(), err
}
