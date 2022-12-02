package command

import . "command/interface"

type OnCommand struct {
    device Device
}

func (c *OnCommand) Execute() {
    c.device.On()
}
