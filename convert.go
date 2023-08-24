package go_bahttext

import (
	"strings"
)

var units = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า", "สิบ"}
var positions = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}

func cleanInput(amount string) string {
	amount = strings.ReplaceAll(amount, "บาท", "")
	amount = strings.ReplaceAll(amount, ",", "")
	amount = strings.ReplaceAll(amount, " ", "")

	return amount
}

func convertIntegerPart(intPart string) string {
	var result string
	for i := 0; i < len(intPart); i++ {
		digit := int(intPart[i] - '0')
		if digit != 0 {
			if i == len(intPart)-1 && digit == 1 {
				result += "เอ็ด"
			} else if i == len(intPart)-2 && digit == 2 {
				result += "ยี่"
			} else if i == len(intPart)-2 && digit == 1 {
				result += ""
			} else {
				result += units[digit]
			}
			result += positions[len(intPart)-i-1]
		}
	}

	if result == "เอ็ด" {
		return "หนึ่ง"
	}

	return result
}

func convertDecimalPart(decimalPart string) string {
	var result string
	for i := 0; i < len(decimalPart); i++ {
		digit := int(decimalPart[i] - '0')
		if digit != 0 {
			if i == len(decimalPart)-1 && digit == 1 {
				result += "เอ็ด"
			} else if i == len(decimalPart)-2 && digit == 2 {
				result += "ยี่"
			} else if i == len(decimalPart)-2 && digit == 1 {
				result += ""
			} else {
				result += units[digit]
			}
			result += positions[len(decimalPart)-i-1]
		}
	}

	if result == "" {
		return "ถ้วน"
	}

	if result == "เอ็ด" {
		return "หนึ่งสตางค์"
	}

	return result + "สตางค์"
}

func ConvertToText(amount string) string {
	parts := strings.Split(cleanInput(amount), ".")
	if len(parts) > 2 {
		return "ทศนิยมหลายตัวนะจ๊ะ"
	}

	intText := convertIntegerPart(parts[0])
	decText := convertDecimalPart(parts[1])

	if intText == "" && decText == "ถ้วน" {
		return "ศูนย์บาทถ้วน"
	}

	if intText == "" {
		intText = "ศูนย์"
	}

	wordRepresentation := intText + "บาท" + decText
	return wordRepresentation
}
