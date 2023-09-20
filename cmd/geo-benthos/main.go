package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	// Import all standard Benthos components
	_ "github.com/benthosdev/benthos/v4/public/components/all"

	// Add the needed needed geo plugins
	_ "github.com/akhenakh/geo-benthos/country"
	_ "github.com/akhenakh/geo-benthos/h3"
	_ "github.com/akhenakh/geo-benthos/randpos"
	_ "github.com/akhenakh/geo-benthos/s2"
	_ "github.com/akhenakh/geo-benthos/tz"
)

func main() {
	service.RunCLI(context.Background())
}
