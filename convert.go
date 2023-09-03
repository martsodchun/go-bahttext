package go_bahttext

import (
	"errors"
	"fmt"
	"regexp"
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
	sliceStr := strings.Split(intPart, "")
	resultSet := make([]string, len(sliceStr)/6+1)
	sets := make([]string, len(sliceStr)/6+1)
	setI := len(sets) - 1
	i := 0

	for j := len(sliceStr) - 1; j >= 0; j-- {
		sets[setI] = sliceStr[j] + sets[setI]
		i++
		if i >= 6 {
			i = 0
			setI--
		}
	}

	for i, set := range sets {
		for j := 0; j < len(set); j++ {
			digit := int(set[j] - '0')
			if digit != 0 {
				if j == len(set)-1 && digit == 1 {
					resultSet[i] += "เอ็ด"
				} else if j == len(set)-2 && digit == 2 {
					resultSet[i] += "ยี่"
				} else if j == len(set)-2 && digit == 1 {
					resultSet[i] += ""
				} else {
					resultSet[i] += units[digit]
				}
				resultSet[i] += positions[len(set)-j-1]
			}
		}

		if resultSet[0] == "เอ็ด" && len(resultSet) == 1 {
			resultSet[0] = "หนึ่ง"
		}
	}
	strResult := strings.Join(resultSet[:], "ล้าน")

	return strResult
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

func ConvertToText(amount string) (string, error) {
	cleanAmount := cleanInput(amount)
	if !isValidNumber(cleanAmount) {
		return "", errors.New("Wrong Format")
	}

	return converting(cleanAmount)
}

func ConvertFloatToText(famount float64) (string, error) {
	cleanAmount := fmt.Sprintf("%f", famount)
	return converting(cleanAmount)
}

func converting(cleanAmount string) (string, error) {
	bahtPart, satangPart := splitInto2(fmt.Sprintf("%s", cleanAmount))

	if len(satangPart) > 2 {
		return "", errors.New("Decimal point more than 2 point")
	}

	intText := convertIntegerPart(bahtPart)
	decText := convertDecimalPart(satangPart)

	if intText == "" && decText == "ถ้วน" {
		return "ศูนย์บาทถ้วน", nil
	}

	if intText == "" {
		intText = "ศูนย์"
	}

	wordRepresentation := intText + "บาท" + decText
	return wordRepresentation, nil
}

func splitInto2(s string) (string, string) {
	index := strings.Index(s, ".")
	if index == -1 {
		return s, ""
	}

	sPart := strings.TrimRight(s[index+1:], "0")
	if len(sPart) == 1 {
		sPart += "0"
	}
	return s[:index], sPart
}

func isValidNumber(s string) bool {
	pattern := `^\d*\.?\d*$`

	matched, _ := regexp.MatchString(pattern, s)

	return matched
}
