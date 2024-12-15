package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetListCategory(url string) []string {
	var listUrl []string
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error: Status code %d\n", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("li.list-none").Each(func(index int, li *goquery.Selection) {
		aTag := li.Find("a")
		href, exists := aTag.Attr("href")
		if exists {
			listUrl = append(listUrl, href)
		}
	})

	return listUrl
}

type Price struct {
	ID              int     `json:"id"`
	MeasureUnitCode int     `json:"measureUnitCode"`
	MeasureUnitName string  `json:"measureUnitName"`
	IsSellDefault   bool    `json:"isSellDefault"`
	Price           float64 `json:"price"`
	CurrencySymbol  string  `json:"currencySymbol"`
	IsDefault       bool    `json:"isDefault"`
	Inventory       int     `json:"inventory"`
	IsInventory     bool    `json:"isInventory"`
}

type Category struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ParentName string `json:"parentName"`
	Slug       string `json:"slug"`
	Level      int    `json:"level"`
	IsActive   bool   `json:"isActive"`
}

type Product struct {
	Sku           string     `json:"sku"`
	Name          string     `json:"name"`
	WebName       string     `json:"webName"`
	Image         string     `json:"image"`
	Category      []Category `json:"category"`
	Price         Price      `json:"price"`
	Slug          string     `json:"slug"`
	Ingredients   string     `json:"ingredients"`
	DosageForm    string     `json:"dosageForm"`
	Brand         string     `json:"brand"`
	Specification string     `json:"specification"`
	Prices        []Price    `json:"prices"`
}

type Response struct {
	Products []Product `json:"products"`
}

func GetListProduct(path string, skipCount int, maxResultCount int) []string {
	// ngoài url thì có thể lấy thêm được các data khác
	url := "https://api.nhathuoclongchau.com.vn/lccus/search-product-service/api/products/ecom/product/search/cate"
	method := "POST"
	pathReplace := strings.Replace(path, "/", "", -1)
	payload := fmt.Sprintf(`{
		"skipCount": %d,
		"maxResultCount": %d,
		"codes": [
			"productTypes",
			"objectUse",
			"priceRanges",
			"prescription",
			"skin",
			"flavor",
			"manufactor",
			"indications",
			"brand",
			"brandOrigin"
		],
		"sortType": 4,
		"category": [
			"%s"
		]
	}`, skipCount, maxResultCount, pathReplace)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}
	var listUrl []string
	// Lọc với các product dung ham getMoreInfo de get them thong tin
	for _, product := range response.Products {
		GetMoreInfoProduct(product.Slug)
	}
	return listUrl
}

func GetMoreInfoProduct(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error: Status code %d\n", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Lấy giá trị từ thẻ <span data-test="price">
	price := doc.Find(`span[data-test="price"]`).Text()
	fmt.Printf("Price: %s\n", price)

	// Lấy các thông tin từ <tr class="content-container">
	doc.Find("tr.content-container").Each(func(i int, tr *goquery.Selection) {
		// Lấy tiêu đề (như "Xuất xứ thương hiệu", "Nhà sản xuất", "Nước sản xuất")
		title := tr.Find("p.css-1c4fxto").Text()
		// Lấy giá trị tương ứng (như "Việt Nam", "CÔNG TY CỔ PHẦN BIGFA")
		value := tr.Find("div.css-1e2qim1").Text()
		fmt.Printf("%s: %s\n", title, value)
	})

	//Description
	doc.Find(`[id^=detail-content-]`).Each(func(index int, section *goquery.Selection) {
		// Extract the section title
		title := section.Find("h2").Text()
		fmt.Println(title)

		// Handle table data if present
		section.Find("table tbody tr").Each(func(i int, row *goquery.Selection) {
			component := row.Find("td:first-child").Text()
			amount := row.Find("td:last-child").Text()
			fmt.Printf("%s %s\n", strings.TrimSpace(component), strings.TrimSpace(amount))
		})

		// Handle paragraph content
		section.Find("div > p").Each(func(i int, paragraph *goquery.Selection) {
			text := paragraph.Text()
			fmt.Println(strings.TrimSpace(text))
		})
		fmt.Println() // Separate sections with a blank line
	})
}
