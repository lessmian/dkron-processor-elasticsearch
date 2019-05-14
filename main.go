package main

import (
	"github.com/lessmian/dkron-processor-elasticsearch"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		Processor: new(elasticsearch),
	})
}
