package eva

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	eva := Eva{}

	result, err := eva.eval(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, result)
}
