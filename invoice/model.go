package invoice

import "github.com/johnfercher/maroto/pkg/pdf"

//Invoice parameters.
type Invoice struct {
	pdf       pdf.Maroto
	Number    string  `yaml:"number"`
	IssueDate string  `yaml:"issueDate"`
	SaleDate  string  `yaml:"saleDate"`
	DueDate   string  `yaml:"dueDate"`
	Notes     string  `yaml:"notes"`
	Company   Company `yaml:"company"`
	Bank      Bank    `yaml:"bank"`
	Items     []*Item `yaml:"items"`
	Currency  string  `yaml:"currency"`
	Signature string  `yaml:"signature"`
	Options   Options `yaml:"options"`
}

//Company details of buyer and seller.
type Company struct {
	Buyer  Buyer  `yaml:"buyer"`
	Seller Seller `yaml:"seller"`
}

//Buyer company details.
type Buyer struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	VAT     string `yaml:"vat"`
}

//Seller company details.
type Seller struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	VAT     string `yaml:"vat"`
}

//Bank details on the invoice.
type Bank struct {
	AccountNumber string `yaml:"accountNumber"`
	Swift         string `yaml:"swift"`
}

//Item parameters.
type Item struct {
	Description string  `yaml:"description"`
	Quantity    float64 `yaml:"quantity"`
	UnitPrice   float64 `yaml:"unitPrice"`
	VATRate     float64 `yaml:"vatRate"`
}

//Options of the PDF document.
type Options struct {
	FontFamily string `yaml:"font" default:"Arial"`
}
