package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Huawei SUN2000-12KTL (PV Meter)",
		Sample: `power:
  source: modbus
  timeout: 2s
  id: 1
  # chose either locally attached:
  device: /dev/ttyUSB0 # serial port
  baudrate: 9600 # Huawei default
  comset: 8N1
  # or via TCP:
  uri: 192.0.2.2:502
  rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter
  # register details
  register:
    address: 37113 # Grid import export power
    type: holding
    decode: int32

currents:
- source: modbus # Huawei phase A grid current
  timeout: 2s
  id: 1
  # chose either locally attached:
  device: /dev/ttyUSB0 # serial port
  baudrate: 9600 # Huawei default
  comset: 8N1
  # or via TCP:
  uri: 192.0.2.2:502
  rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter
  # register details
  register: # manual non-sunspec register configuration
    address: 37107 # Huawei phase A grid current
    type: holding
    decode: int32

- source: modbus # Huawei phase B grid current
  timeout: 2s
  id: 1
  # chose either locally attached:
  device: /dev/ttyUSB0 # serial port
  baudrate: 9600 # Huawei default
  comset: 8N1
  # or via TCP:
  uri: 192.0.2.2:502
  rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter
  # register details
  register: # manual non-sunspec register configuration
    address: 37109 # Huawei phase B grid current
    type: holding
    decode: int32

- source: modbus # Huawei phase C grid current
  timeout: 2s
  id: 1
  # chose either locally attached:
  device: /dev/ttyUSB0 # serial port
  baudrate: 9600 # Huawei default
  comset: 8N1
  # or via TCP:
  uri: 192.0.2.2:502
  rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter
  # register details
  register: # manual non-sunspec register configuration
    address: 37111 # Huawei phase C grid current
    type: holding
    decode: int32`,
	}

	registry.Add(template)
}
