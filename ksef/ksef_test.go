package ksef

import (
	"strings"
	"testing"

	"github.com/cnvergence/invoice-generator/invoice"
)

func TestGenerate_NormalizesVATIdentifiers(t *testing.T) {
	inv := &invoice.Invoice{
		Number:    "1/07/2026",
		IssueDate: "31-07-2026",
		SaleDate:  "31-07-2026",
		Currency:  "EUR",
		Company: invoice.Company{
			Seller: invoice.Seller{
				Name:        "Seller",
				Address:     "Street 1, City",
				VAT:         "PL5842784571",
				CountryCode: "PL",
			},
			Buyer: invoice.Buyer{
				Name:        "Buyer GmbH",
				Address:     "Strasse 2, Stadt",
				VAT:         "DE306148241",
				CountryCode: "PL",
			},
		},
		Items: []*invoice.Item{{
			Description: "Software development",
			Quantity:    1,
			UnitPrice:   100,
			VATRateCode: "np II",
		}},
	}

	out, err := Generate(inv)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	xml := string(out)

	for _, want := range []string{
		"<NIP>5842784571</NIP>",
		"<PrefiksPodatnika>PL</PrefiksPodatnika>",
		"<KodUE>DE</KodUE>",
		"<NrVatUE>306148241</NrVatUE>",
		"<KodKraju>DE</KodKraju>",
		"<P_13_9>100.00</P_13_9>",
		"<P_18>1</P_18>",
	} {
		if !strings.Contains(xml, want) {
			t.Errorf("missing %s in generated xml", want)
		}
	}
	for _, reject := range []string{
		"RodzajFormularza",
		"<NIP>PL5842784571</NIP>",
		"<NIP>DE306148241</NIP>",
	} {
		if strings.Contains(xml, reject) {
			t.Errorf("unexpected %s in generated xml", reject)
		}
	}
}
