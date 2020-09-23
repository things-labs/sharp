package qrcode

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/thinkgos/go-core-package/extos"
	"github.com/thinkgos/go-core-package/extstr"
	"github.com/thinkgos/go-core-package/lib/algo"
	"golang.org/x/image/bmp"
)

// 扩展名,支持四种图片二维码生成
const (
	ExtJPG = ".jpg"
	ExtPNG = ".png"
	ExtGIF = ".gif"
	ExtBMP = ".bmp"
)

// MetaInfo 元信息
type MetaInfo struct {
	Content string // 内容
	Level   qr.ErrorCorrectionLevel
	Mode    qr.Encoding
	Width   int
	Height  int
	Ext     string // 命名扩展名
}

// Generate generate to barcode.Barcode
func (sf *MetaInfo) Generate() (barcode.Barcode, error) {
	code, err := qr.Encode(sf.Content, sf.Level, sf.Mode)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(code, sf.Width, sf.Height)
}

// GenerateToBytes generate to byte
func (sf *MetaInfo) GenerateToBytes() ([]byte, string, error) {
	if !extstr.ContainsFold([]string{ExtJPG, ExtPNG, ExtGIF, ExtBMP}, sf.Ext) {
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
	if extos.IsExist(dst) {
		return filename, nil
	}
	return filename, ioutil.WriteFile(dst, data, 0664)
}

// GenerateToImageB64 generate QR code to base images string, return base64 image string and filename
// base64 images string format like:
// data: image/png;base64,xxxxxxxxxx
func (sf *MetaInfo) GenerateToImageB64() (string, string, error) {
	b64, filename, err := sf.GenerateToBase64()
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("data:image/%s;base64,%s",
		strings.TrimLeft(sf.Ext, "."), b64), filename, nil
}

// GenerateToBase64 generate QR code to base64 string, return base64 string and filename
func (sf MetaInfo) GenerateToBase64() (string, string, error) {
	data, filename, err := sf.GenerateToBytes()
	if err != nil {
		return "", "", err
	}
	return base64.StdEncoding.EncodeToString(data), filename, nil
}
