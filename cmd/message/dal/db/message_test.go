package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/internal/conf"
)

func TestQueryMessage(t *testing.T) {
	conf.Init()
	Init()
	//message, err := QueryLatestMessage(context.Background(), 2, 3)
	message, err := QueryLatestMessage(context.Background(), 4, 3)
	assert.NoError(t, err)
	fmt.Printf("%v", message)
}
