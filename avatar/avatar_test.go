package avatar

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RandomImage(t *testing.T) {
	avt := Avatar{
		290, 290,
		4096, 4096,
	}
	img, err := avt.RandomImage([]byte("gogs@local"))
	require.NoError(t, err)
	assert.Equal(t, 290, img.Bounds().Max.X)
	assert.Equal(t, 290, img.Bounds().Max.Y)

	avt.Height = 280
	img, err = avt.RandomImage([]byte("gogs@local"))
	require.NoError(t, err)
	assert.Equal(t, 280, img.Bounds().Max.X)
	assert.Equal(t, 280, img.Bounds().Max.Y)

	avt.Width = 0
	avt.Height = 0
	_, err = avt.RandomImage([]byte("gogs@local"))
	require.Error(t, err)
}

func Test_PrepareWithPNG(t *testing.T) {
	avt := Avatar{
		290, 290,
		4096, 4096,
	}

	data, err := ioutil.ReadFile("testdata/avatar.png")
	require.NoError(t, err)

	img, err := avt.Prepare(data)
	require.NoError(t, err)

	assert.Equal(t, 290, img.Bounds().Max.X)
	assert.Equal(t, 290, img.Bounds().Max.Y)
}

func Test_PrepareWithJPEG(t *testing.T) {
	avt := Avatar{
		290, 290,
		4096, 4096,
	}
	data, err := ioutil.ReadFile("testdata/avatar.jpeg")
	require.NoError(t, err)

	imgPtr, err := avt.Prepare(data)
	require.NoError(t, err)

	assert.Equal(t, 290, imgPtr.Bounds().Max.X)
	assert.Equal(t, 290, imgPtr.Bounds().Max.Y)
}

func Test_PrepareWithInvalidImage(t *testing.T) {
	avt := Avatar{
		290, 290,
		5, 5,
	}
	_, err := avt.Prepare([]byte{})
	require.EqualError(t, err, "DecodeConfig: image: unknown format")
}

func Test_PrepareWithInvalidImageSize(t *testing.T) {
	avt := Avatar{
		290, 290,
		5, 5,
	}
	data, err := ioutil.ReadFile("testdata/avatar.png")
	require.NoError(t, err)

	_, err = avt.Prepare(data)
	assert.EqualError(t, err, "Image width is too large: 10 > 5")

	avt.MaxWidth = 4095

	_, err = avt.Prepare(data)
	assert.EqualError(t, err, "Image height is too large: 10 > 5")
}
