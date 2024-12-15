package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func GetListCategoryMinhChau(url string) []string {
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

	// Find all <a> tags and extract the href attribute
	doc.Find("div.title_bottom.title-sp a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Println(href)
		}
	})

	return listUrl
}

func GetListProductMinhChau(url string) []string {
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

	// Find all <a> tags and extract the href attribute
	doc.Find("div.itemsanpham").Each(func(i int, s *goquery.Selection) {
		// Extract the link from the <a> tag inside <div class="img">
		link, exists := s.Find("div.img a").Attr("href")
		if exists {
			// Print the extracted link
			fmt.Println("Product Link:", link)
		}
	})

	return listUrl
}

func GetDetailProductMinhChau(url string) {
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

	title := doc.Find("title").Text()

	fmt.Println("name", title)

	// category
	doc.Find(".content_main .content_content #breadcrumbs a").Each(func(i int, s *goquery.Selection) {
		// Extract the title attribute of each <a> tag
		category, exists := s.Attr("title")
		if exists {
			// Print the extracted title
			fmt.Println("category:", category)
		}
	})

	// Create a map to store the extracted data
	dataMap := make(map[string]string)

	// Find all rows with class 'row_chitiet' containing a label and value
	doc.Find(".row_chitiet").Each(func(i int, s *goquery.Selection) {
		// Extract label and value text
		label := s.Find(".ct_label").Text()
		value := s.Find(".row_noidung").Text()

		// If both label and value are found, add to map
		if label != "" && value != "" {
			// Clean up the label and value
			label = strings.TrimSpace(label)
			value = strings.TrimSpace(value)

			// Add to map
			dataMap[label] = value
		}

		// If the row contains a list (<ul>), extract the list items
		if s.Find(".row_noidung ul").Length() > 0 {
			var listItems []string
			// Iterate over each <li> item and append to listItems
			s.Find(".row_noidung ul li").Each(func(j int, li *goquery.Selection) {
				listItems = append(listItems, li.Text())
			})
			// Join all list items into one string and use the label 'Description' for the key
			dataMap["Description"] = strings.Join(listItems, "; ")
		}
	})

	// Print the resulting map
	for key, value := range dataMap {
		fmt.Printf("%s: %s\n", key, value)
	}

	doc.Find("div.tab-noidungchitiet table.table tbody tr").Each(func(index int, row *goquery.Selection) {
		// Extract text from the first and second <td> elements in the row
		key := row.Find("td").First().Text()
		value := row.Find("td").Next().Text()

		// Clean up whitespace
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		// Print the result
		fmt.Printf("%s\t%s\n", key, value)
	})

}
