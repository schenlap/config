package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Huawei SUN2000-12KTL (Pv Meter)",
		Sample: `power:
  source: modbus
  timeout: 2s
  id: 1
  # chose either locally attached:
  device: /dev/ttyUSB0 # serial port
  baudrate: 19200
  comset: 8N1
  # or via TCP:
  uri: 192.0.2.2:502
  rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter
  # register details
  register:
    address: 32080 # Active Huawei Solar generation power
    type: holding
    decode: int32`,
	}

	registry.Add(template)
}
