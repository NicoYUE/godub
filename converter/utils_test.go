package converter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStrSet(t *testing.T) {
	p, _ := os.Getwd()
	print(p)
	bin, err := AvailableBinary()
	if err != nil {
		print(err)
	}
	assert.Equal(t, bin, "")
}
