package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveLineBreak(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"hello\nhello",
			"hello hello",
		},
		{
			"hello \nhello",
			"hello hello",
		},
		{
			"hello\n hello",
			"hello hello",
		},
		{
			"    hello    hello   \n   hello  hello   ",
			"    hello    hello hello  hello   ",
		},
		{
			"    hello   ",
			"    hello   ",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, removeLineBreak(tt.input))
	}
}

func TestRemoteImagesDisabledByDefault(t *testing.T) {
	r := NewRenderer(40, 0)
	assert.False(t, r.allowRemoteImages)

	reader, err := r.imageFromDestination("https://example.com/image.png")
	assert.Error(t, err)
	assert.Nil(t, reader)
}

func TestWithRemoteImages(t *testing.T) {
	r := NewRenderer(40, 0, WithRemoteImages(true))
	assert.True(t, r.allowRemoteImages)
}
