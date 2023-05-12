package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	// Import all standard Benthos components
	_ "github.com/benthosdev/benthos/v4/public/components/all"

	// Add the geo plugin you want
	_ "github.com/akhenakh/geo-benthos/h3"
	_ "github.com/akhenakh/geo-benthos/s2"
)

func main() {
	service.RunCLI(context.Background())
}