package sharp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NormalizeEOL(t *testing.T) {
	data1 := []string{
		"",
		"This text starts with empty lines",
		"another",
		"",
		"",
		"",
		"Some other empty lines in the middle",
		"more.",
		"And more.",
		"Ends with empty lines too.",
		"",
		"",
		"",
	}

	data2 := []string{
		"This text does not start with empty lines",
		"another",
		"",
		"",
		"",
		"Some other empty lines in the middle",
		"more.",
		"And more.",
		"Ends without EOLtoo.",
	}

	buildEOLData := func(data []string, eol string) []byte {
		return []byte(strings.Join(data, eol))
	}

	dos := buildEOLData(data1, "\r\n")
	unix := buildEOLData(data1, "\n")
	mac := buildEOLData(data1, "\r")

	assert.Equal(t, unix, NormalizeEOL(dos))
	assert.Equal(t, unix, NormalizeEOL(mac))
	assert.Equal(t, unix, NormalizeEOL(unix))

	dos = buildEOLData(data2, "\r\n")
	unix = buildEOLData(data2, "\n")
	mac = buildEOLData(data2, "\r")

	assert.Equal(t, unix, NormalizeEOL(dos))
	assert.Equal(t, unix, NormalizeEOL(mac))
	assert.Equal(t, unix, NormalizeEOL(unix))

	assert.Equal(t, []byte("one liner"), NormalizeEOL([]byte("one liner")))
	assert.Equal(t, []byte("\n"), NormalizeEOL([]byte("\n")))
	assert.Equal(t, []byte("\ntwo liner"), NormalizeEOL([]byte("\ntwo liner")))
	assert.Equal(t, []byte("two liner\n"), NormalizeEOL([]byte("two liner\n")))
	assert.Equal(t, []byte{}, NormalizeEOL([]byte{}))

	assert.Equal(t, []byte("mix\nand\nmatch\n."), NormalizeEOL([]byte("mix\r\nand\rmatch\n.")))
}
