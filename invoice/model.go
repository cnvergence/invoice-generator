package invoice

import "github.com/johnfercher/maroto/pkg/pdf"

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
}

type Company struct {
	Buyer  Buyer  `yaml:"buyer"`
	Seller Seller `yaml:"seller"`
}

type Buyer struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	VAT     string `yaml:"vat"`
}

type Seller struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	VAT     string `yaml:"vat"`
}

type Bank struct {
	AccountNumber string `yaml:"accountNumber"`
	Swift         string `yaml:"swift"`
}

type Item struct {
	Description string  `yaml:"description"`
	Quantity    float64 `yaml:"quantity"`
	UnitPrice   float64 `yaml:"unitPrice"`
	VATRate     float64 `yaml:"vatRate"`
}
