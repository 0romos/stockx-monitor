package scraper

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ProductInfo struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

func ScrapeProductInfo(url string) (*ProductInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to fetch the URL")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	product := &ProductInfo{}

	name := doc.Find("h1.chakra-heading[data-component='primary-product-title']").First().Text()
	subName := doc.Find("span.chakra-heading[data-component='secondary-product-title']").First().Text()
	product.Name = strings.TrimSpace(name + " " + subName)

	foundPrice := false
	doc.Find("div.css-1e47tnd").Each(func(i int, s *goquery.Selection) {
		text := s.Find("a.chakra-button.css-1tlej2y p.chakra-text.css-1dy2wii").Text()
		if strings.Contains(text, "Buy for") {
			priceStr := strings.ReplaceAll(strings.Split(text, "€")[1], ",", "")
			priceStr = strings.TrimSpace(priceStr)
			priceFloat, err := strconv.ParseFloat(priceStr, 64)
			if err == nil {
				product.Price = priceFloat
				foundPrice = true
			}
		}
	})
	if !foundPrice {
		return nil, errors.New("price not found")
	}

	imageSrc := doc.Find("img.chakra-image.css-1sh8ayr").AttrOr("src", "")
	if imageSrc == "" {
		return nil, errors.New("image not found")
	}
	product.Image = imageSrc

	return product, nil
}
