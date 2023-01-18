package conf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	fmt.Println(Config)
	assert.Nil(t, err)
	assert.NotNil(t, *Config)
}
