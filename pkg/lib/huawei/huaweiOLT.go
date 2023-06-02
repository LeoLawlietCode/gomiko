package huawei

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

type HuaweiOLT struct {
	Driver driver.IDriver
	Prompt string
	base   types.HuaweiDevice
}

func (d *HuaweiOLT) Connect() error {
	return d.base.Connect()

}

func (d *HuaweiOLT) Disconnect() {
	d.base.Disconnect()

}

func (d *HuaweiOLT) SetMode(mode string) {
	d.base.SetMode(mode)
}

func (d *HuaweiOLT) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *HuaweiOLT) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)
}
