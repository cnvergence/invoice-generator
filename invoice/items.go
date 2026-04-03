package invoice

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// buildTable prepares table with items on the invoice with calculated tax amounts and total gross amounts.
func (i *Invoice) buildTable() {
	header := getHeader()
	items := i.getItems()
	taxes, totals := i.countTax()
	contents := appendItems(items, taxes, totals)
	gridSizes := []int{1, 3, 1, 2, 1, 1, 3}

	i.pdf.AddRows(
		row.New(2).WithStyle(&props.Cell{BackgroundColor: getTealColor()}).Add(col.New(12)),
	)

	headerCols := make([]core.Col, len(header))
	for j, h := range header {
		headerCols[j] = col.New(gridSizes[j]).Add(
			text.New(h, props.Text{
				Style: fontstyle.Bold,
				Size:  8,
				Align: align.Center,
				Color: getTealColor(),
			}),
		)
	}
	i.pdf.AddRows(row.New(8).Add(headerCols...))

	for idx, content := range contents {
		contentCols := make([]core.Col, len(gridSizes))
		for j := range gridSizes {
			var val string
			if j < len(content) {
				val = content[j]
			}
			contentCols[j] = col.New(gridSizes[j]).Add(
				text.New(val, props.Text{
					Style: fontstyle.Normal,
					Size:  10,
					Align: align.Center,
				}),
			)
		}
		contentRow := row.New(8).Add(contentCols...)
		if idx%2 == 0 {
			contentRow = contentRow.WithStyle(&props.Cell{BackgroundColor: getGrayColor()})
		}
		i.pdf.AddRows(contentRow)
	}

	i.pdf.AddRows(
		row.New(10).Add(
			col.New(8),
			col.New(2).WithStyle(&props.Cell{BackgroundColor: getTealColor()}).Add(
				text.New("Total:", props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
					Color: getWhiteColor(),
				}),
			),
			col.New(2).WithStyle(&props.Cell{BackgroundColor: getTealColor()}).Add(
				text.New(fmt.Sprintf("%s %s", calculateInvoiceSum(contents), i.Currency), props.Text{
					Top:   3,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Center,
					Color: getWhiteColor(),
				}),
			),
		),
	)
}

func getHeader() []string {
	return []string{"No", "Description", "Quantity", "Unit net price", "VAT rate", "VAT amount", "Total gross price"}
}

func calculateInvoiceSum(values [][]string) string {
	var sum float64
	for _, value := range values {
		num, _ := strconv.ParseFloat(value[len(value)-1], 64)
		sum = sum + num
	}

	return strconv.FormatFloat(sum, 'f', 2, 64)
}

func (i *Invoice) countTax() ([]float64, []float64) {
	var taxes []float64
	var totals []float64

	for _, item := range i.Items {
		vat := item.VATRate
		price := item.UnitPrice
		quantity := item.Quantity

		tax := quantity * (vat * price / 100)
		total := quantity*price + tax

		taxes = append(taxes, tax)
		totals = append(totals, total)
	}

	return taxes, totals
}

func appendItems(values [][]string, taxes []float64, totals []float64) [][]string {
	number := 1
	for i := range values {
		values[i] = append([]string{strconv.Itoa(number)}, values[i]...)
		values[i] = append(values[i], strconv.FormatFloat(taxes[i], 'f', 2, 64))
		values[i] = append(values[i], strconv.FormatFloat(totals[i], 'f', 2, 64))
		number++
	}

	return values
}

func (i *Invoice) getItems() [][]string {
	var items [][]string

	v := reflect.Indirect(reflect.ValueOf(i.Items))
	if v.Kind() != reflect.Slice {
		return nil
	}

	for i := range make([]struct{}, v.Len()) {
		e := reflect.Indirect(v.Index(i))

		if e.Kind() != reflect.Struct {
			return nil
		}
		var element []string
		for fieldIdx := range make([]struct{}, e.NumField()) {
			element = append(element, fmt.Sprint(e.Field(fieldIdx).Interface()))
		}

		items = append(items, element)
	}

	return items
}
