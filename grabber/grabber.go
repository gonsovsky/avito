package grabber

import (
	. "avito/shared"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Grab(url string) (AvitoPage, error) {
	x := AvitoPage{}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("span.title-info-title-text").Each(func(_ int, s *goquery.Selection) {
		x.Title = s.Text()
	})
	doc.Find("span.js-item-price").Each(func(_ int, s *goquery.Selection) {
		x.Price = s.Text()
		x.PriceInt, _ = s.Attr("content")
	})
	doc.Find("span.gallery-img-cover").Each(func(_ int, s *goquery.Selection) {
		var n, _ = s.Attr("style")
		fmt.Println(n)
		x.Image = "https:" + between(n, "('", "')")
	})
	x.URL = url
	return x, nil
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
