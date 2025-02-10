package utils

import (
	"net/url"
	"strings"
)

func ExtractDomainFromUrl(link string) (string, error) {
	parser, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(parser.Hostname(), "www."), nil
}
