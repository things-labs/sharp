package paginator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParam(t *testing.T) {
	require.Equal(t, Param{1, DefaultPageSize}, DefaultParam())

	param := Param{}
	assert.EqualValues(t, &Param{1, 20}, param.Inspect(20))
	param = Param{}
	assert.Equal(t, &Param{1, DefaultPageSize}, param.Inspect())
}
