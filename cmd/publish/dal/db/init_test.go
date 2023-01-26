package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/internal/conf"
)

func TestInit(t *testing.T) {
	_ = conf.Init()
	err := Init()
	fmt.Println(DB.Name())
	assert.NoError(t, err)
	assert.NotNil(t, DB)
}
