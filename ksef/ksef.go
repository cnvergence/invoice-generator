package ksef

import (
	"encoding/xml"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/cnvergence/invoice-generator/invoice"
)

const xmlns = "http://crd.gov.pl/wzor/2025/06/25/13775/"

// Faktura is the root element of the KSeF FA_3 e-invoice.
type Faktura struct {
	XMLName  xml.Name `xml:"Faktura"`
	Xmlns    string   `xml:"xmlns,attr"`
	XmlnsXsi string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd string   `xml:"xmlns:xsd,attr"`
	Naglowek Naglowek `xml:"Naglowek"`
	Podmiot1 Podmiot1 `xml:"Podmiot1"`
	Podmiot2 Podmiot2 `xml:"Podmiot2"`
	Fa       Fa       `xml:"Fa"`
}

// Naglowek – header section.
type Naglowek struct {
	KodFormularza     KodFormularza `xml:"KodFormularza"`
	WariantFormularza int           `xml:"WariantFormularza"`
	DataWytworzeniaFa string        `xml:"DataWytworzeniaFa"`
	SystemInfo        string        `xml:"SystemInfo,omitempty"`
}

// KodFormularza carries fixed attributes and the "FA" text value.
type KodFormularza struct {
	Value        string `xml:",chardata"`
	KodSystemowy string `xml:"kodSystemowy,attr"`
	WersjaSchemy string `xml:"wersjaSchemy,attr"`
}

// Podmiot1 – seller.
type Podmiot1 struct {
	DaneIdentyfikacyjne DaneIdPodmiot1 `xml:"DaneIdentyfikacyjne"`
	Adres               Adres          `xml:"Adres"`
}

// DaneIdPodmiot1 – seller identity (NIP required).
type DaneIdPodmiot1 struct {
	NIP   string `xml:"NIP"`
	Nazwa string `xml:"Nazwa"`
}

// Podmiot2 – buyer.
type Podmiot2 struct {
	DaneIdentyfikacyjne DaneIdPodmiot2 `xml:"DaneIdentyfikacyjne"`
	Adres               *Adres         `xml:"Adres,omitempty"`
	JST                 int            `xml:"JST"`
	GV                  int            `xml:"GV"`
}

// DaneIdPodmiot2 – buyer identity: NIP, EU VAT (KodUE+NrVatUE), or BrakID.
type DaneIdPodmiot2 struct {
	NIP     string `xml:"NIP,omitempty"`
	KodUE   string `xml:"KodUE,omitempty"`
	NrVatUE string `xml:"NrVatUE,omitempty"`
	BrakID  string `xml:"BrakID,omitempty"`
	Nazwa   string `xml:"Nazwa,omitempty"`
}

// Adres – postal address block.
type Adres struct {
	KodKraju string `xml:"KodKraju"`
	AdresL1  string `xml:"AdresL1"`
	AdresL2  string `xml:"AdresL2,omitempty"`
}

// Fa – core invoice data.
type Fa struct {
	KodWaluty string `xml:"KodWaluty"`
	P_1       string `xml:"P_1"`
	P_1M      string `xml:"P_1M,omitempty"` // place of issue
	P_2       string `xml:"P_2"`
	// P_6: sale/delivery date – omitted when equal to issue date
	P_6 string `xml:"P_6,omitempty"`
	// VAT group totals (each pair is an optional XSD sequence)
	P_13_1   string `xml:"P_13_1,omitempty"`   // net  23 %
	P_14_1   string `xml:"P_14_1,omitempty"`   // VAT  23 %
	P_13_2   string `xml:"P_13_2,omitempty"`   // net   8 %
	P_14_2   string `xml:"P_14_2,omitempty"`   // VAT   8 %
	P_13_3   string `xml:"P_13_3,omitempty"`   // net   5 %
	P_14_3   string `xml:"P_14_3,omitempty"`   // VAT   5 %
	P_13_6_1 string `xml:"P_13_6_1,omitempty"` // net 0% domestic (0 KR)
	P_13_6_2 string `xml:"P_13_6_2,omitempty"` // net 0% WDT
	P_13_6_3 string `xml:"P_13_6_3,omitempty"` // net 0% export
	P_13_7   string `xml:"P_13_7,omitempty"`   // value exempt (zw)
	P_13_8   string `xml:"P_13_8,omitempty"`   // np I (not subject)
	P_13_9   string `xml:"P_13_9,omitempty"`   // np II (not subject)
	P_13_10  string `xml:"P_13_10,omitempty"`  // reverse charge (oo)
	// P_15: total gross payable (required)
	P_15          string     `xml:"P_15"`
	Adnotacje     Adnotacje  `xml:"Adnotacje"`
	RodzajFaktury string     `xml:"RodzajFaktury"`
	FaWiersz      []FaWiersz `xml:"FaWiersz,omitempty"`
	Platnosc      *Platnosc  `xml:"Platnosc,omitempty"`
}

// Adnotacje – annotation flags required by the schema.
type Adnotacje struct {
	P_16                 int                  `xml:"P_16"`
	P_17                 int                  `xml:"P_17"`
	P_18                 int                  `xml:"P_18"`
	P_18A                int                  `xml:"P_18A"`
	Zwolnienie           Zwolnienie           `xml:"Zwolnienie"`
	NoweSrodkiTransportu NoweSrodkiTransportu `xml:"NoweSrodkiTransportu"`
	P_23                 int                  `xml:"P_23"`
	PMarzy               PMarzy               `xml:"PMarzy"`
}

// Zwolnienie – "no VAT exemption" branch (standard invoices).
type Zwolnienie struct {
	P_19N int `xml:"P_19N"`
}

// NoweSrodkiTransportu – "no new means of transport" branch.
type NoweSrodkiTransportu struct {
	P_22N int `xml:"P_22N"`
}

// PMarzy – "no margin procedure" branch.
type PMarzy struct {
	P_PMarzyN int `xml:"P_PMarzyN"`
}

// FaWiersz – a single invoice line item.
type FaWiersz struct {
	NrWierszaFa int    `xml:"NrWierszaFa"`
	P_7         string `xml:"P_7,omitempty"`        // description
	P_8A        string `xml:"P_8A,omitempty"`       // unit of measure
	P_8B        string `xml:"P_8B,omitempty"`       // quantity
	P_9A        string `xml:"P_9A,omitempty"`       // unit price net
	P_11        string `xml:"P_11,omitempty"`       // line net total
	P_12        string `xml:"P_12,omitempty"`       // VAT rate code
	KursWaluty  string `xml:"KursWaluty,omitempty"` // exchange rate
}

// Platnosc – payment terms.
type Platnosc struct {
	FormaPlatnosci  int               `xml:"FormaPlatnosci,omitempty"`
	RachunekBankowy []RachunekBankowy `xml:"RachunekBankowy,omitempty"`
}

// RachunekBankowy – bank account details.
type RachunekBankowy struct {
	NrRB       string `xml:"NrRB"`
	SWIFT      string `xml:"SWIFT,omitempty"`
	NazwaBanku string `xml:"NazwaBanku,omitempty"`
}

// Generate converts an Invoice into a KSeF FA_3 XML document.
// Returns the full XML bytes including the XML declaration header.
func Generate(inv *invoice.Invoice) ([]byte, error) {
	issueDate := parseDate(inv.IssueDate)
	saleDate := parseDate(inv.SaleDate)
	if saleDate == issueDate {
		saleDate = "" // P_6 omitted when equal to P_1
	}

	faktura := Faktura{
		Xmlns:    xmlns,
		XmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsXsd: "http://www.w3.org/2001/XMLSchema",
		Naglowek: Naglowek{
			KodFormularza: KodFormularza{
				Value:        "FA",
				KodSystemowy: "FA (3)",
				WersjaSchemy: "1-0E",
			},
			WariantFormularza: 3,
			// DataWytworzeniaFa must be >= 2025-09-01T00:00:00Z (FA_3 schema constraint)
			DataWytworzeniaFa: time.Now().UTC().Format(time.RFC3339),
		},
		Podmiot1: buildPodmiot1(inv),
		Podmiot2: buildPodmiot2(inv),
		Fa:       buildFa(inv, issueDate, saleDate),
	}

	out, err := xml.MarshalIndent(faktura, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal XML: %w", err)
	}
	return append([]byte(xml.Header), out...), nil
}

func buildPodmiot1(inv *invoice.Invoice) Podmiot1 {
	s := inv.Company.Seller
	return Podmiot1{
		DaneIdentyfikacyjne: DaneIdPodmiot1{
			NIP:   s.VAT,
			Nazwa: s.Name,
		},
		Adres: buildAdres(s.CountryCode, s.AddressLine1, s.Address, s.AddressLine2),
	}
}

func buildPodmiot2(inv *invoice.Invoice) Podmiot2 {
	b := inv.Company.Buyer
	var daneId DaneIdPodmiot2
	switch {
	case b.EUCode != "" && b.EUVatNumber != "":
		daneId = DaneIdPodmiot2{KodUE: b.EUCode, NrVatUE: b.EUVatNumber, Nazwa: b.Name}
	case b.NIPMissing:
		daneId = DaneIdPodmiot2{BrakID: "1", Nazwa: b.Name}
	default:
		daneId = DaneIdPodmiot2{NIP: b.VAT, Nazwa: b.Name}
	}
	addr := buildAdres(b.CountryCode, b.AddressLine1, b.Address, b.AddressLine2)
	return Podmiot2{
		DaneIdentyfikacyjne: daneId,
		Adres:               &addr,
		JST:                 2,
		GV:                  2,
	}
}

// buildAdres constructs an Adres, falling back to the single-line address
// when addressLine1 is not explicitly set.
func buildAdres(countryCode, addressLine1, addressFallback, addressLine2 string) Adres {
	if countryCode == "" {
		countryCode = "PL"
	}
	line1 := addressLine1
	if line1 == "" {
		line1 = addressFallback
	}
	return Adres{
		KodKraju: countryCode,
		AdresL1:  line1,
		AdresL2:  addressLine2,
	}
}

// vatSums accumulates net and VAT amounts grouped by standard PL VAT rate.
type vatSums struct {
	net23, vat23 float64
	net8, vat8   float64
	net5, vat5   float64
	net0kr       float64 // 0% domestic (0 KR)
	net0wdt      float64 // 0% WDT
	net0ex       float64 // 0% export
	netZw        float64 // exempt (zw)
	netNpI       float64 // np I
	netNpII      float64 // np II
	netOO        float64 // reverse charge (oo)
	totalGross   float64
	hasOO        bool // true if any item uses reverse charge
}

// effectiveVATCode returns the VAT rate code to use for an item,
// preferring the string VATRateCode over the numeric VATRate.
func effectiveVATCode(item *invoice.Item) string {
	if item.VATRateCode != "" {
		return item.VATRateCode
	}
	return vatRateToString(item.VATRate)
}

func computeSums(items []*invoice.Item) vatSums {
	var s vatSums
	for _, item := range items {
		net := round2(item.Quantity * item.UnitPrice)
		code := effectiveVATCode(item)

		switch code {
		case "23", "22":
			vatAmt := round2(net * item.VATRate / 100)
			s.net23 += net
			s.vat23 += vatAmt
			s.totalGross += net + vatAmt
		case "8", "7":
			vatAmt := round2(net * item.VATRate / 100)
			s.net8 += net
			s.vat8 += vatAmt
			s.totalGross += net + vatAmt
		case "5":
			vatAmt := round2(net * item.VATRate / 100)
			s.net5 += net
			s.vat5 += vatAmt
			s.totalGross += net + vatAmt
		case "0 KR":
			s.net0kr += net
			s.totalGross += net
		case "0 WDT":
			s.net0wdt += net
			s.totalGross += net
		case "0 EX":
			s.net0ex += net
			s.totalGross += net
		case "zw":
			s.netZw += net
			s.totalGross += net
		case "np I":
			s.netNpI += net
			s.totalGross += net
		case "np II":
			s.netNpII += net
			s.totalGross += net
		case "oo":
			s.netOO += net
			s.hasOO = true
			s.totalGross += net
		default:
			// For any other numeric rate, compute normally
			vatAmt := round2(net * item.VATRate / 100)
			s.totalGross += net + vatAmt
		}
	}
	s.net23 = round2(s.net23)
	s.vat23 = round2(s.vat23)
	s.net8 = round2(s.net8)
	s.vat8 = round2(s.vat8)
	s.net5 = round2(s.net5)
	s.vat5 = round2(s.vat5)
	s.net0kr = round2(s.net0kr)
	s.net0wdt = round2(s.net0wdt)
	s.net0ex = round2(s.net0ex)
	s.netZw = round2(s.netZw)
	s.netNpI = round2(s.netNpI)
	s.netNpII = round2(s.netNpII)
	s.netOO = round2(s.netOO)
	s.totalGross = round2(s.totalGross)
	return s
}

func buildFa(inv *invoice.Invoice, issueDate, saleDate string) Fa {
	sums := computeSums(inv.Items)

	fa := Fa{
		KodWaluty:     inv.Currency,
		P_1:           issueDate,
		P_1M:          inv.IssuePlace,
		P_2:           inv.Number,
		P_6:           saleDate,
		P_15:          fmtAmt(sums.totalGross),
		Adnotacje:     buildAdnotacje(sums),
		RodzajFaktury: "VAT",
	}

	if sums.net23 > 0 || sums.vat23 > 0 {
		fa.P_13_1 = fmtAmt(sums.net23)
		fa.P_14_1 = fmtAmt(sums.vat23)
	}
	if sums.net8 > 0 || sums.vat8 > 0 {
		fa.P_13_2 = fmtAmt(sums.net8)
		fa.P_14_2 = fmtAmt(sums.vat8)
	}
	if sums.net5 > 0 || sums.vat5 > 0 {
		fa.P_13_3 = fmtAmt(sums.net5)
		fa.P_14_3 = fmtAmt(sums.vat5)
	}
	if sums.net0kr > 0 {
		fa.P_13_6_1 = fmtAmt(sums.net0kr)
	}
	if sums.net0wdt > 0 {
		fa.P_13_6_2 = fmtAmt(sums.net0wdt)
	}
	if sums.net0ex > 0 {
		fa.P_13_6_3 = fmtAmt(sums.net0ex)
	}
	if sums.netZw > 0 {
		fa.P_13_7 = fmtAmt(sums.netZw)
	}
	if sums.netNpI > 0 {
		fa.P_13_8 = fmtAmt(sums.netNpI)
	}
	if sums.netNpII > 0 {
		fa.P_13_9 = fmtAmt(sums.netNpII)
	}
	if sums.netOO > 0 {
		fa.P_13_10 = fmtAmt(sums.netOO)
	}

	for i, item := range inv.Items {
		w := FaWiersz{
			NrWierszaFa: i + 1,
			P_7:         item.Description,
			P_8A:        item.Unit,
			P_8B:        fmtQty(item.Quantity),
			P_9A:        fmtPrice(item.UnitPrice),
			P_11:        fmtAmt(round2(item.Quantity * item.UnitPrice)),
			P_12:        effectiveVATCode(item),
		}
		if item.ExchangeRate > 0 {
			w.KursWaluty = trimZeros(fmt.Sprintf("%.6f", item.ExchangeRate))
		}
		fa.FaWiersz = append(fa.FaWiersz, w)
	}

	if p := buildPlatnosc(inv); p != nil {
		fa.Platnosc = p
	}
	return fa
}

func buildAdnotacje(sums vatSums) Adnotacje {
	p18 := 2
	if sums.hasOO {
		p18 = 1 // reverse charge: "odwrotne obciążenie"
	}
	return Adnotacje{
		P_16:                 2,
		P_17:                 2,
		P_18:                 p18,
		P_18A:                2,
		Zwolnienie:           Zwolnienie{P_19N: 1},
		NoweSrodkiTransportu: NoweSrodkiTransportu{P_22N: 1},
		P_23:                 2,
		PMarzy:               PMarzy{P_PMarzyN: 1},
	}
}

func buildPlatnosc(inv *invoice.Invoice) *Platnosc {
	hasAccount := inv.Bank.AccountNumber != ""
	if !hasAccount {
		return nil
	}
	p := &Platnosc{FormaPlatnosci: 6} // 6 = Przelew (bank transfer)
	p.RachunekBankowy = []RachunekBankowy{{
		NrRB:       inv.Bank.AccountNumber,
		SWIFT:      inv.Bank.Swift,
		NazwaBanku: inv.Bank.BankName,
	}}
	return p
}

// parseDate converts DD-MM-YYYY → YYYY-MM-DD (as required by KSeF date fields).
func parseDate(s string) string {
	parts := strings.Split(s, "-")
	if len(parts) != 3 || len(parts[2]) != 4 {
		return s // return as-is if not recognisable
	}
	return fmt.Sprintf("%s-%s-%s", parts[2], parts[1], parts[0])
}

// vatRateToString maps a numeric VAT rate to the KSeF TStawkaPodatku code.
func vatRateToString(rate float64) string {
	switch rate {
	case 23:
		return "23"
	case 22:
		return "22"
	case 8:
		return "8"
	case 7:
		return "7"
	case 5:
		return "5"
	case 4:
		return "4"
	case 3:
		return "3"
	case 0:
		return "0 KR"
	case -1:
		return "zw"
	case -2:
		return "oo"
	default:
		return fmt.Sprintf("%.0f", rate)
	}
}

// round2 rounds a float64 to 2 decimal places.
func round2(x float64) float64 {
	return math.Round(x*100) / 100
}

// fmtAmt formats a monetary amount as TKwotowy (2 dp).
func fmtAmt(x float64) string {
	return fmt.Sprintf("%.2f", x)
}

// fmtQty formats a quantity as TIlosci (up to 6 dp, trailing zeros trimmed).
func fmtQty(x float64) string {
	return trimZeros(fmt.Sprintf("%.6f", x))
}

// fmtPrice formats a unit price as TKwotowy2 (up to 8 dp, trailing zeros trimmed).
func fmtPrice(x float64) string {
	return trimZeros(fmt.Sprintf("%.8f", x))
}

// trimZeros removes trailing decimal zeros and the decimal point if unneeded.
func trimZeros(s string) string {
	if !strings.Contains(s, ".") {
		return s
	}
	s = strings.TrimRight(s, "0")
	return strings.TrimRight(s, ".")
}
