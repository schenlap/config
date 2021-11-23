package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "",
		Name:   "",
		Sample: ``,
	}

	registry.Add(template)
}
