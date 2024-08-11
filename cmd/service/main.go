package main

import (
	"os"

	"github.com/amnestia/xyz-multifinance/internal/api/server"
)

func main() {
	os.Exit(server.New().Run())
}
