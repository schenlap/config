package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "garo",
		Name:   "PC Electric Garo",
		Sample: `uri: http://192.0.2.2:8080/servlet
meter: <CENTRAL100|CENTRAL101|INTERNAL|EXTERNAL|TWIN> # Value can be found at http://192.0.2.2:8080/servlet/rest/chargebox/status `,
	}

	registry.Add(template)
}
