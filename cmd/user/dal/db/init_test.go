package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/internal/conf"
)

func TestInit(t *testing.T) {
	conf.Init()
	Init()
	fmt.Println(conf.Config)
	assert.NotNil(t, DB)
}
