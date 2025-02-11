package utils

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func MinifyHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	doc.Find("*").Each(func(index int, item *goquery.Selection) {
		var str string
		str, err = item.Html()
		str = strings.TrimSpace(str)
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, "\t", "")
		if err == nil {
			item.SetHtml(str)
		}
	})

	htmlStr, err := doc.Html()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(htmlStr), nil
}

func ResolveURL(base, ref string) (string, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	refURL, err := url.Parse(ref)
	if err != nil {
		return "", err
	}
	return baseURL.ResolveReference(refURL).String(), nil
}
