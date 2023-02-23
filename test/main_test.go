package test

import (
	"context"
	"testing"
	"tiktok-server/cmd/api/rpc"
)

const TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IklsbFRhbWVyIiwiaWQiOjIsImV4cCI6MTA5MDA1MTQxNjQsImlhdCI6MTY3NzE0MjEyOCwiaXNzIjoiYnl0ZWRhbmNlIn0.eRvQ9KUKFxF4QkQIp7Xf17PFmVEeOT2eqDeQYUChjyg"

var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()
	rpc.InitRPC()
	m.Run()
}
