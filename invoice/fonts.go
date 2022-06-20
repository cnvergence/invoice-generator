package invoice

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/flopp/go-findfont"
	"github.com/johnfercher/maroto/pkg/consts"
)

func (i *Invoice) setFonts() error {
	if i.Options.FontFamily != "Arial" {
		fontPath, err := findfont.Find(i.Options.FontFamily)
		if err != nil {
			return fmt.Errorf("could not find the installed font: %s", err)
		}
		i.pdf.SetFontLocation(filepath.Dir(fontPath))
		fontlist := findfont.List()
		fonts := filterFonts(fontlist, func(val string) bool {
			return strings.Contains(val, i.Options.FontFamily)
		})
		for _, font := range fonts {
			if strings.Contains(font, "Regular") || strings.EqualFold(font, i.Options.FontFamily) {
				i.pdf.AddUTF8Font(i.Options.FontFamily, consts.Normal, filepath.Base(font))
			}
			if strings.Contains(font, "Italic") {
				i.pdf.AddUTF8Font(i.Options.FontFamily, consts.Italic, filepath.Base(font))
			}
			if strings.Contains(font, "Bold") {
				i.pdf.AddUTF8Font(i.Options.FontFamily, consts.Bold, filepath.Base(font))
			}
			if strings.Contains(font, "BoldItalic") || strings.Contains(font, "Bold Italic") {
				i.pdf.AddUTF8Font(i.Options.FontFamily, consts.BoldItalic, filepath.Base(font))
			}
		}
		i.pdf.SetDefaultFontFamily(i.Options.FontFamily)
	}
	return nil
}

func filterFonts(fonts []string, cond func(string) bool) []string {
	result := []string{}
	for i := range fonts {
		if cond(fonts[i]) {
			result = append(result, fonts[i])
		}
	}
	return result
}
