package go_bahttext

import (
	"fmt"
)

func main() {
	exampleAmount := "7,532.73 บาท"
	fmt.Printf("%s => %s\n", exampleAmount, ConvertToText(exampleAmount))
}
