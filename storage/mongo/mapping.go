package mongo

import (
	"github.com/Revazashvili/easer/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AsDbModel(t models.Template, id primitive.ObjectID) Template {
	return Template{
		Id:           id,
		Name:         t.Name,
		Description:  t.Description,
		Owner:        t.Owner,
		TemplateBody: t.TemplateBody,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Options: Options{
			DisableExternalLinks: t.Options.DisableExternalLinks,
			DisableInternalLinks: t.Options.DisableInternalLinks,
			Dpi:                  t.Options.Dpi,
			DisplayHeaderFooter:  t.Options.DisplayHeaderFooter,
			EnableForms:          t.Options.EnableForms,
			Format:               t.Options.Format,
			Grayscale:            t.Options.Grayscale,
			NoImages:             t.Options.NoImages,
			NoBackground:         t.Options.NoBackground,
			Orientation:          t.Options.Orientation,
			PrintBackground:      t.Options.PrintBackground,
			Margin: Margin{
				Top:    t.Options.Margin.Top,
				Bottom: t.Options.Margin.Bottom,
				Left:   t.Options.Margin.Left,
				Right:  t.Options.Margin.Left,
			},
			HeaderFooterOptions: HeaderAndFooterOptions{
				FooterCenter:   t.Options.HeaderFooterOptions.FooterCenter,
				FooterHTML:     t.Options.HeaderFooterOptions.FooterHTML,
				FooterLeft:     t.Options.HeaderFooterOptions.FooterLeft,
				FooterLine:     t.Options.HeaderFooterOptions.FooterLine,
				FooterFontName: t.Options.HeaderFooterOptions.FooterFontName,
				FooterFontSize: t.Options.HeaderFooterOptions.FooterFontSize,
				FooterRight:    t.Options.HeaderFooterOptions.FooterRight,
				FooterSpacing:  t.Options.HeaderFooterOptions.FooterSpacing,
				HeaderCenter:   t.Options.HeaderFooterOptions.HeaderCenter,
				HeaderHTML:     t.Options.HeaderFooterOptions.HeaderHTML,
				HeaderLeft:     t.Options.HeaderFooterOptions.HeaderLeft,
				HeaderLine:     t.Options.HeaderFooterOptions.HeaderLine,
				HeaderFontName: t.Options.HeaderFooterOptions.HeaderFontName,
				HeaderFontSize: t.Options.HeaderFooterOptions.HeaderFontSize,
				HeaderRight:    t.Options.HeaderFooterOptions.HeaderRight,
				HeaderSpacing:  t.Options.HeaderFooterOptions.HeaderSpacing,
			},
		},
	}
}

func AsDomainList(ts []Template) []models.Template {
	out := make([]models.Template, len(ts))

	for i, t := range ts {
		out[i] = AsDomain(t)
	}
	return out
}

func AsDomain(t Template) models.Template {
	return models.Template{
		Id:           t.Id.Hex(),
		Name:         t.Name,
		Description:  t.Description,
		Owner:        t.Owner,
		TemplateBody: t.TemplateBody,
		Options: models.Options{
			DisableExternalLinks: t.Options.DisableExternalLinks,
			DisableInternalLinks: t.Options.DisableInternalLinks,
			Dpi:                  t.Options.Dpi,
			DisplayHeaderFooter:  t.Options.DisplayHeaderFooter,
			EnableForms:          t.Options.EnableForms,
			Format:               t.Options.Format,
			Grayscale:            t.Options.Grayscale,
			NoImages:             t.Options.NoImages,
			NoBackground:         t.Options.NoBackground,
			Orientation:          t.Options.Orientation,
			PrintBackground:      t.Options.PrintBackground,
			Margin: models.Margin{
				Top:    t.Options.Margin.Top,
				Bottom: t.Options.Margin.Bottom,
				Left:   t.Options.Margin.Left,
				Right:  t.Options.Margin.Left,
			},
			HeaderFooterOptions: models.HeaderAndFooterOptions{
				FooterCenter:   t.Options.HeaderFooterOptions.FooterCenter,
				FooterHTML:     t.Options.HeaderFooterOptions.FooterHTML,
				FooterLeft:     t.Options.HeaderFooterOptions.FooterLeft,
				FooterLine:     t.Options.HeaderFooterOptions.FooterLine,
				FooterFontName: t.Options.HeaderFooterOptions.FooterFontName,
				FooterFontSize: t.Options.HeaderFooterOptions.FooterFontSize,
				FooterRight:    t.Options.HeaderFooterOptions.FooterRight,
				FooterSpacing:  t.Options.HeaderFooterOptions.FooterSpacing,
				HeaderCenter:   t.Options.HeaderFooterOptions.HeaderCenter,
				HeaderHTML:     t.Options.HeaderFooterOptions.HeaderHTML,
				HeaderLeft:     t.Options.HeaderFooterOptions.HeaderLeft,
				HeaderLine:     t.Options.HeaderFooterOptions.HeaderLine,
				HeaderFontName: t.Options.HeaderFooterOptions.HeaderFontName,
				HeaderFontSize: t.Options.HeaderFooterOptions.HeaderFontSize,
				HeaderRight:    t.Options.HeaderFooterOptions.HeaderRight,
				HeaderSpacing:  t.Options.HeaderFooterOptions.HeaderSpacing,
			},
		},
	}
}
