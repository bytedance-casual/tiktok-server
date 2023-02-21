package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/internal/conf"
)

func TestMCheckFavorite(t *testing.T) {
	conf.Init()
	Init()
	favorite, err := MCheckFavorite(3, []int64{0, 1, 2, 3}, context.Background())
	assert.NoError(t, err)
	fmt.Printf("%v", favorite)
}
