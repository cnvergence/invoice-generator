package invoice

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"gopkg.in/yaml.v3"
)

//New returns Invoice struct loaded with values from YAML and prepares PDF struct.
func New(file []byte) *Invoice {
	invoice := &Invoice{}
	if err := yaml.Unmarshal(file, &invoice); err != nil {
		return nil
	}

	invoice.pdf = pdf.NewMaroto(consts.Portrait, consts.A4)
	invoice.setPDFLayout()

	return invoice
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

func (i *Invoice) setPDFLayout() {
	i.pdf.SetFirstPageNb(1)
	i.pdf.AddUTF8Font("Montserrat", consts.Normal, "fonts/Montserrat-Regular.ttf")
	i.pdf.AddUTF8Font("Montserrat", consts.Italic, "fonts/Montserrat-Italic.ttf")
	i.pdf.AddUTF8Font("Montserrat", consts.Bold, "fonts/Montserrat-Bold.ttf")
	i.pdf.AddUTF8Font("Montserrat", consts.BoldItalic, "fonts/Montserrat-BoldItalic.ttf")
	i.pdf.SetDefaultFontFamily("Montserrat")
	i.pdf.SetPageMargins(10, 15, 10)

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
}

//SaveToPdf saves PDF to a file and closes it.
func (i *Invoice) SaveToPdf(outputPath string) error {
	err := i.pdf.OutputFileAndClose(outputPath)

	return err
}

//Save saves PDF to a file and closes it.
func (i *Invoice) SaveAsBytes() ([]byte, error) {
	bytes, err := i.pdf.Output()

	return bytes.Bytes(), err
}
