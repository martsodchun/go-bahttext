package go_bahttext

import (
	"fmt"
)

func main() {
	exampleAmount := "7,532.73 บาท"
	result, err := ConvertToText(exampleAmount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s => %s\n", exampleAmount, result)

}
