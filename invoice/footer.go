package invoice

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

func (i *Invoice) BuildFooter() {
	i.pdf.RegisterFooter(func() {
		i.pdf.SetAliasNbPages("{nbs}")
		currentPage := strconv.Itoa(i.pdf.GetCurrentPage())
		base64StringImage, err := loadImage("img/go.png")
		if err != nil {
			fmt.Println("Could not load image:", err)
			os.Exit(1)
		}
		i.pdf.Row(6, func() {
			i.pdf.Col(12, func() {
				i.pdf.Text("Generated with", props.Text{
					Top:   1,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: getTealColor(),
				})

			})
		})
		i.pdf.Row(8, func() {
			i.pdf.Col(15, func() {
				_ = i.pdf.Base64Image(base64StringImage, consts.Png, props.Rect{
					Left:    5,
					Top:     1,
					Center:  false,
					Percent: 100,
				})
			})
		})
		i.pdf.Row(6, func() {
			i.pdf.Col(12, func() {
				i.pdf.Text(fmt.Sprintf("Page %s of {nbs}", currentPage), props.Text{
					Top:   1,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: getTealColor(),
				})
			})
		})
	})

}

func loadImage(path string) (string, error) {
	img, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(img), nil
}
