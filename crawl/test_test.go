package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"testing"
)

func TestGoQuery(t *testing.T) {
	html := `
	<div>
		<script type="application/ld+json">
		{
			"@context": "https://schema.org",
			"@type": "Product",
			"name": "Gel Chống Nắng ANESSA Perfect UV Dưỡng Da Ẩm Mịn SPF50+/PA++++ (Chai 90g)",
			"additionalProperty": [
				{
					"@type": "PropertyValue",
					"name": "Tên sản phẩm",
					"value": "ANESSA PERFECT UV SUNSCREEN SKINCARE GEL"
				},
				{
					"@type": "PropertyValue",
					"name": "Danh mục",
					"value": "Kem chống nắng dành cho mặt"
				},
				{
					"@type": "PropertyValue",
					"name": "Công dụng",
					"value": "Thành phần thiên nhiên chiết xuất lá trà xanh, hoa anh đào, rễ cây Potentilla dưỡng da mềm mịn, góp phần bảo vệ làn da trước tia cực tím và môi trường ô nhiễm. 50% chiết xuất từ hoa hồng tây, collagen cá biển sâu, lô hội và super Hyaluronic Acid giúp da mịn màng, cung cấp đủ độ ẩm nhưng không tạo cảm giác nhờn rít."
				},
				{
					"@type": "PropertyValue",
					"name": "Nhà sản xuất",
					"value": "Shiseido"
				},
				{
					"@type": "PropertyValue",
					"name": "Quy cách",
					"value": "90g"
				},
				{
					"@type": "PropertyValue",
					"name": "Lưu ý",
					"value": "Mọi thông tin trên đây chỉ mang tính chất tham khảo. Đọc kỹ hướng dẫn sử dụng trước khi dùng"
				}
			]
		}
		</script>
	</div>`

	// Sử dụng goquery để tìm <script type="application/ld+json">
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// Tìm đoạn JSON trong thẻ <script>
	var jsonData string
	doc.Find("script[type='application/ld+json']").Each(func(i int, s *goquery.Selection) {
		jsonData = s.Text()
	})

	// Giải mã JSON
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
}
