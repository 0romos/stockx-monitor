package monitor

import (
	"stockx-monitor/config"
	"stockx-monitor/internal/scraper"
	"sync"
)

type ProductData struct {
	URL     string               `json:"url"`
	Details *scraper.ProductInfo `json:"details"`
}

func GetProductDetails() ([]ProductData, error) {
	var products []ProductData
	var mu sync.Mutex
	var wg sync.WaitGroup
	var err error

	for _, url := range config.ProductURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			productInfo, scrapeErr := scraper.ScrapeProductInfo(url)
			if scrapeErr != nil {
				err = scrapeErr
				return
			}
			mu.Lock()
			products = append(products, ProductData{URL: url, Details: productInfo})
			mu.Unlock()
		}(url)
	}

	wg.Wait()

	return products, err
}
