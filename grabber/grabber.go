package grabber

import (
	. "avito/shared"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func Grab(page AvitoPage) (AvitoPage, error) {

	res, err := http.Get(page.Url)
	if err != nil {
		return page, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return page, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return page, err
	}

	doc.Find("span.title-info-title-text").Each(func(_ int, s *goquery.Selection) {
		page.Title = s.Text()
	})
	doc.Find("span.js-item-price").Each(func(_ int, s *goquery.Selection) {
		page.Price = s.Text()
		page.PriceInt, _ = s.Attr("content")
	})
	doc.Find("span.gallery-img-cover").Each(func(_ int, s *goquery.Selection) {
		var n, _ = s.Attr("style")
		fmt.Println(n)
		page.Image = "https:" + between(n, "('", "')")
	})
	return page, nil
}

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}
