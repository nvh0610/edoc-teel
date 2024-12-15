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

func GetListPharmacyCategory(url string) []string {
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

	doc.Find(".swiper.category-playlist.w-full .swiper-slide a").Each(func(i int, s *goquery.Selection) {
		// Lấy giá trị của thuộc tính href
		href, exists := s.Attr("href")
		if exists {
			fmt.Println(href)
		}
	})

	return listUrl
}

func GetPharmacyProduct(path string) []string {
	param := strings.Split(path, "?")[0]
	pathReplace := strings.Replace(param, "/", "", -1)
	url := fmt.Sprintf("https://api-gateway.pharmacity.vn/pmc-ecm-product/api/public/search/index?platform=1&index=1&limit=20&total=0&refresh=true&page=category&page_slug=%s", pathReplace)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
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
	fmt.Println(string(body))

	return nil
}

type ItemListElement struct {
	Name string `json:"name"`
}

type BreadcrumbList struct {
	ItemListElement []ItemListElement `json:"itemListElement"`
}

type PropertyValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProductPharmacy struct {
	AdditionalProperty []PropertyValue `json:"additionalProperty"`
	AtType             string          `json:"@type"`
}

func GetMorePharmacyProduct(url string) {
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

	//Get category
	var jsonData string
	doc.Find("script[type='application/ld+json']").Each(func(i int, s *goquery.Selection) {
		jsonData = s.Text()
		var product ProductPharmacy
		err = json.Unmarshal([]byte(jsonData), &product)
		if err != nil {
			log.Fatal(err)
		}

		// Tạo map giữa name và value
		propertyMap := make(map[string]string)
		for _, property := range product.AdditionalProperty {
			propertyMap[property.Name] = property.Value
		}

		// In map giữa name và value
		for name, value := range propertyMap {
			fmt.Printf("Name: %s, Value: %s\n", name, value)
		}
	})

	// Giải mã JSON
	var breadcrumb BreadcrumbList
	err = json.Unmarshal([]byte(jsonData), &breadcrumb)
	if err != nil {
		log.Fatal(err)
	}

	// In các giá trị `name`
	for _, item := range breadcrumb.ItemListElement {
		fmt.Println(item.Name)
	}

	// description
	var content string
	doc.Find("div.pmc-content-html p").Each(func(i int, s *goquery.Selection) {
		content += s.Text() + "\n"
	})

	// In ra tất cả nội dung với đúng định dạng
	fmt.Println("Tất cả nội dung:\n", content)
}
