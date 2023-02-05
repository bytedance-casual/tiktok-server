package oss

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/internal/conf"
)

func TestInit(t *testing.T) {
	conf.Init()
	Init()
	assert.NotNil(t, *bucket)
}
