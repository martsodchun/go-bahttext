# go-bahttext

[![Build Status](https://travis-ci.org/martsodchun/go-bahttext.svg?branch=main)](https://travis-ci.org/martsodchun/go-bahttext)
[![GoDoc](https://godoc.org/github.com/martsodchun/go-bahttext/bahttext?status.svg)](https://godoc.org/github.com/martsodchun/go-bahttext/bahttext)

A Go library to convert numeric values to Thai word representations.

หลิบไทยบ้านสำหรับแปลงเลขเป็นตัวหนังสือไว้ใช้ในการออกเช็ค หรือใช้ในการออกใบเสร็จ
ใช้ฟรีไปเลยครับหามานานไม่เคยเจอเลย จนเขียนขึ้นมาเอง แล้วก็เลยแชร์ให้คนอื่นได้ใช้กัน อิอิ
เจอบัคเรียกด้วยนะครับ แจ้งได้เลย ขอบคุณครับ

## Usage

```go
import "github.com/martsodchun/go-bahttext"

amount := "5,555.55 บาท"
fmt.Printf("%s => %s\n", amount, bahttext.ConvertToWords(amount))

=> "ห้าพันห้าร้อยห้าสิบห้าบาทห้าสิบห้าสตางค์"
```

## License

MIT License

Copyright (c) 2023 Ponthakorn Sodchun (martsodchun)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.