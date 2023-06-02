package huawei

import (
	"errors"

	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.HuaweiDevice, error) {
	devDriver := driver.NewDriver(connection, "\n")
	base := HWDevice{
		Driver:     devDriver,
		Prompt:     "",
		DeviceType: DeviceType,
	}

	switch DeviceType {
	case "huawei_olt":
		return &HuaweiOLT{
			Driver: devDriver,
			Prompt: "",
			base:   &base,
		}, nil
	default:
		return nil, errors.New("unsupported DeviceType: " + DeviceType)
	}

}
