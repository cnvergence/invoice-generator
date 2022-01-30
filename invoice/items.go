package invoice

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

func (i *Invoice) BuildTable() {
	backgroundColor := getGrayColor()
	header := getHeader()
	items := i.getItems()
	taxes, totals := i.countTax()
	contents := appendItems(items, taxes, totals)

	i.pdf.SetBackgroundColor(getTealColor())
	i.pdf.Row(2, func() {
		i.pdf.Col(12, func() {
		})
	})
	i.pdf.SetBackgroundColor(color.NewWhite())
	i.pdf.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Style:     consts.Normal,
			Size:      8,
			GridSizes: []uint{1, 3, 1, 2, 1, 1, 3},
			Color:     getTealColor(),
		},
		ContentProp: props.TableListContent{
			Style:     consts.Normal,
			Size:      10,
			GridSizes: []uint{1, 3, 1, 2, 1, 1, 3},
		},
		Align:                consts.Center,
		AlternatedBackground: &backgroundColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	i.pdf.Row(10, func() {
		i.pdf.ColSpace(8)
		i.pdf.SetBackgroundColor(getTealColor())
		i.pdf.Col(2, func() {
			i.pdf.Text("Total:", props.Text{
				Top:   3,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
				Color: color.NewWhite(),
			})
		})
		i.pdf.Col(2, func() {
			i.pdf.Text(fmt.Sprintf("%s %s", calculateInvoiceSum(contents), i.Currency), props.Text{
				Top:   3,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
				Color: color.NewWhite(),
			})
		})
	})

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

	for j := 0; j < len(i.Items); j++ {
		vat := i.Items[j].VATRate
		price := i.Items[j].UnitPrice
		quantity := i.Items[j].Quantity

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

	for j := 0; j < v.Len(); j++ {
		e := reflect.Indirect(v.Index(j))

		if e.Kind() != reflect.Struct {
			return nil
		}
		var element []string
		for j := 0; j < e.NumField(); j++ {
			element = append(element, fmt.Sprint(e.Field(j).Interface()))
		}

		items = append(items, element)
	}

	return items
}
