package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOffsetFromPage(t *testing.T) {

	t.Run("OffsetFromPage should return 0 when page is 0", func(t *testing.T) {
		offset := OffsetFromPage(0, 1)
		assert.Equal(t, 0, offset)
	})

	t.Run("OffsetFromPage should return 0 when limit is 10", func(t *testing.T) {
		offset := OffsetFromPage(2, 0)
		assert.Equal(t, 10, offset)
	})

	t.Run("OffsetFromPage should return 2 when page 2 and limit is 2", func(t *testing.T) {
		offset := OffsetFromPage(2, 2)
		assert.Equal(t, 2, offset)
	})
}
