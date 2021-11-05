package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "cfos",
		Name:   "cFos PowerBrain",
		Sample: `uri: 192.0.2.2:502
id: 1
# an evcc sponsortoken is required for using this charger`,
	}

	registry.Add(template)
}
