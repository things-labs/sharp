package qrcode

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
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

// Encode encode to barcode.Barcode
func (sf *MetaInfo) Encode() (barcode.Barcode, error) {
	code, err := qr.Encode(sf.Content, sf.Level, sf.Mode)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(code, sf.Width, sf.Height)
}

// Filename filename
func (sf *MetaInfo) Filename() string {
	return algo.SHA1(sf.Content) + sf.Ext
}

// Write wirte QR code
func (sf *MetaInfo) Write(w io.Writer) error {
	if !extstr.ContainsFold([]string{ExtJPG, ExtPNG, ExtGIF, ExtBMP}, sf.Ext) {
		return fmt.Errorf("not support image format: %s", sf.Ext)
	}

	code, err := sf.Encode()
	if err != nil {
		return err
	}
	switch strings.ToLower(sf.Ext) {
	case ExtJPG:
		err = jpeg.Encode(w, code, nil)
	case ExtPNG:
		err = png.Encode(w, code)
	case ExtGIF:
		err = gif.Encode(w, code, nil)
	case ExtBMP:
		err = bmp.Encode(w, code)
	}
	return err
}

// WriteFile write QR code to file, return filename
func (sf *MetaInfo) WriteFile(path string) (string, error) {
	filename := sf.Filename()
	dstFilename := filepath.Join(path, filename)
	if extos.IsExist(dstFilename) {
		return filename, nil
	}

	f, err := os.OpenFile(dstFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return filename, sf.Write(f)
}

// ToBytes generate to byte
func (sf *MetaInfo) ToBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := sf.Write(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToBase64 generate QR code to base64 string, return base64 string
func (sf *MetaInfo) ToBase64() (string, error) {
	data, err := sf.ToBytes()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// ToImageB64 generate QR code to base images string, return base64 image string
// base64 images string format like:
// data: image/png;base64,xxxxxxxxxx
func (sf *MetaInfo) ToImageB64() (string, error) {
	b64, err := sf.ToBase64()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("data:image/%s;base64,%s", strings.TrimLeft(sf.Ext, "."), b64), nil
}
