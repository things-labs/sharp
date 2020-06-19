package qrcode

import (
	"bytes"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/thinkgos/sharp"
	"github.com/thinkgos/sharp/algo"
	"github.com/thinkgos/strext"
	"golang.org/x/image/bmp"
)

// 支持四种图片二维码生成
const (
	ExtJPG = ".jpg"
	ExtPNG = ".png"
	ExtGIF = ".gif"
	ExtBMP = ".bmp"
)

// 元信息
type MetaInfo struct {
	Content string
	Level   qr.ErrorCorrectionLevel
	Mode    qr.Encoding
	Width   int
	Height  int
	Ext     string
}

func (sf *MetaInfo) Generate() (barcode.Barcode, error) {
	code, err := qr.Encode(sf.Content, sf.Level, sf.Mode)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(code, sf.Width, sf.Height)
}

func (sf *MetaInfo) GenerateToBytes() ([]byte, string, error) {
	if !strext.ContainsFold([]string{ExtJPG, ExtPNG, ExtGIF, ExtBMP}, sf.Ext) {
		return nil, "", fmt.Errorf("not support image format: %s", sf.Ext)
	}

	code, err := sf.Generate()
	if err != nil {
		return nil, "", err
	}

	buf := &bytes.Buffer{}
	switch strings.ToLower(sf.Ext) {
	case ExtJPG:
		err = jpeg.Encode(buf, code, nil)
	case ExtPNG:
		err = png.Encode(buf, code)
	case ExtGIF:
		err = gif.Encode(buf, code, nil)
	case ExtBMP:
		err = bmp.Encode(buf, code)
	}
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), algo.SHA1(sf.Content) + sf.Ext, nil
}

// GenerateToFile generate QR code, return filename
func (sf *MetaInfo) GenerateToFile(path string) (string, error) {
	data, filename, err := sf.GenerateToBytes()
	if err != nil {
		return "", err
	}
	dst := filepath.Join(path, filename)
	if sharp.IsExist(dst) {
		return filename, nil
	}
	return filename, ioutil.WriteFile(dst, data, 0664)
}
