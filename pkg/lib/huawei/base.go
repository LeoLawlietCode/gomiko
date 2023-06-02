package huawei

import (
	"errors"
	"strings"

	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
)

type HWDevice struct {
	Driver     driver.IDriver
	DeviceType string
	Prompt     string
	Mode       string
}

func (d *HWDevice) Connect() error {

	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt(`[\]#>]`, "#|>")
	if err != nil {
		return err
	}
	d.Prompt = prompt
	return d.sessionPreparation()

}

func (d *HWDevice) Disconnect() {

	d.Driver.Disconnect()

}

func (d *HWDevice) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *HWDevice) SendConfigSet(cmds []string) (string, error) {

	results, err := d.Driver.SendCommandsSet(cmds, d.Prompt)

	return results, err

}

func (d *HWDevice) sessionPreparation() error {
	// Disable interactive mode in Huawei OLT
	_, err := d.Driver.SendCommand("undo smart\n\n", d.Prompt)
	if err != nil {
		return errors.New("gomiko: failed to disable interactive")
	}
	// Disable pager in Huawei OLT
	_, err = d.Driver.SendCommand("scroll\n\n", d.Prompt)
	if err != nil {
		return errors.New("gomiko: failed to disable pager")
	}
	// Change mode to 'enable', 'config' o in global mode by default
	if d.Mode != "" {
		switch d.Mode {
		case "enable":
			out, _ := d.Driver.SendCommand("enable", "#")
			d.Prompt = "#"

			if !strings.Contains(out, "#") {
				return errors.New("gomiko: failed to enter enable mode")
			}
		case "config":
			out, _ := d.Driver.SendCommand("enable", "#")
			d.Prompt = "#"

			if !strings.Contains(out, "#") {
				return errors.New("gomiko: failed to enter enable mode")
			}

			out, _ = d.Driver.SendCommand("config", "#")
			if !strings.Contains(out, "#") {
				return errors.New("gomiko: failed to enter config mode")
			}
		default:
			return errors.New("gomiko: '" + d.Mode + "' mode doesn't exist")
		}
	}
	return err
}

func (d *HWDevice) SetMode(mode string) {
	d.Mode = mode
}
