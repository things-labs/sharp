// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package avatar

import (
	"bytes"
	"fmt"
	"image"
	"image/color/palette"

	// Enable PNG support:
	_ "image/jpeg"
	_ "image/png"

	"math/rand"
	"time"

	"github.com/issue9/identicon"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

// AvatarSize returns avatar's size
// const AvatarSize = 290

type Avatar struct {
	Width     uint
	Height    uint
	MaxWidth  int
	MaxHeight int
}

// RandomImage generates and returns a random avatar image unique to input data
// in default size (Min(height,width) height and width).
func (sf *Avatar) RandomImage(data []byte) (image.Image, error) {
	randExtent := len(palette.WebSafe) - 32
	rand.Seed(time.Now().UnixNano())
	colorIndex := rand.Intn(randExtent)
	backColorIndex := colorIndex - 1
	if backColorIndex < 0 {
		backColorIndex = randExtent - 1
	}

	size := sf.Width
	if size > sf.Height {
		size = sf.Height
	}

	// Define size, background, and forecolor
	imgMaker, err := identicon.New(int(size),
		palette.WebSafe[backColorIndex], palette.WebSafe[colorIndex:colorIndex+32]...)
	if err != nil {
		return nil, fmt.Errorf("identicon.New: %v", err)
	}
	return imgMaker.Make(data), nil
}

// Prepare accepts a byte slice as input, validates it contains an image of an
// acceptable format, and crops and resizes it appropriately.
func (sf Avatar) Prepare(data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	imgCfg, _, err := image.DecodeConfig(reader)
	if err != nil {
		return nil, fmt.Errorf("DecodeConfig: %v", err)
	}
	if imgCfg.Width > sf.MaxWidth {
		return nil, fmt.Errorf("Image width is too large: %d > %d", imgCfg.Width, sf.MaxWidth)
	}
	if imgCfg.Height > sf.MaxHeight {
		return nil, fmt.Errorf("Image height is too large: %d > %d", imgCfg.Height, sf.MaxHeight)
	}
	reader.Reset(data)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("Decode: %v", err)
	}

	if imgCfg.Width != imgCfg.Height {
		var newSize, ax, ay int
		if imgCfg.Width > imgCfg.Height {
			newSize = imgCfg.Height
			ax = (imgCfg.Width - imgCfg.Height) / 2
		} else {
			newSize = imgCfg.Width
			ay = (imgCfg.Height - imgCfg.Width) / 2
		}

		img, err = cutter.Crop(img, cutter.Config{
			Width:  newSize,
			Height: newSize,
			Anchor: image.Point{ax, ay},
		})
		if err != nil {
			return nil, err
		}
	}

	img = resize.Resize(sf.Width, sf.Height, img, resize.NearestNeighbor)
	return img, nil
}
