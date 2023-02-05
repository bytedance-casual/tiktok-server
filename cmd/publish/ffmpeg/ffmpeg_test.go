package ffmpeg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVideoShot(t *testing.T) {
	err := GetVideoShot("../test/bear.png", "../test/bear.mp4")
	assert.NoError(t, err)
}
