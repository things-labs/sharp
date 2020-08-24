package main

import (
	"log"

	"github.com/boombuler/barcode/qr"

	"github.com/thinkgos/sharp/qrcode"
)

func main() {
	meta := qrcode.MetaInfo{
		Content: "http://www.baidu.com",
		Level:   qr.M,
		Mode:    qr.Auto,
		Width:   100,
		Height:  100,
		Ext:     qrcode.ExtJPG,
	}
	data, name, err := meta.GenerateToImageB64()
	if err != nil {
		panic(err)
	}
	log.Println(name)
	log.Println(data)
}
