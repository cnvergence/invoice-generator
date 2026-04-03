package invoice

import (
	"fmt"
	"strings"

	findfont "github.com/flopp/go-findfont"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/repository"
)

func (i *Invoice) configureFonts(builder config.Builder) (config.Builder, error) {
	if i.Options.FontFamily == "" {
		return builder, nil
	}

	_, err := findfont.Find(i.Options.FontFamily)
	if err != nil {
		return nil, fmt.Errorf("could not find font %s installed: %s", i.Options.FontFamily, err)
	}

	fontlist := findfont.List()
	fonts := filterFonts(fontlist, func(val string) bool {
		return strings.Contains(val, i.Options.FontFamily)
	})

	repo := repository.New()
	for _, font := range fonts {
		if strings.Contains(font, "BoldItalic") || strings.Contains(font, "Bold Italic") {
			repo = repo.AddUTF8Font(i.Options.FontFamily, fontstyle.BoldItalic, font)
		} else if strings.Contains(font, "Bold") {
			repo = repo.AddUTF8Font(i.Options.FontFamily, fontstyle.Bold, font)
		} else if strings.Contains(font, "Italic") {
			repo = repo.AddUTF8Font(i.Options.FontFamily, fontstyle.Italic, font)
		} else if strings.Contains(font, "Regular") || strings.EqualFold(strings.TrimSuffix(font, ".ttf"), i.Options.FontFamily) {
			repo = repo.AddUTF8Font(i.Options.FontFamily, fontstyle.Normal, font)
		}
	}

	customFonts, err := repo.Load()
	if err != nil {
		return nil, fmt.Errorf("could not load fonts: %s", err)
	}

	return builder.
		WithCustomFonts(customFonts).
		WithDefaultFont(&props.Font{Family: i.Options.FontFamily}), nil
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
