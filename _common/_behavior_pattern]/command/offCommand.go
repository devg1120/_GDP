package command

import . "command/interface"

type OffCommand struct {
    device Device
}

func (c *OffCommand) Execute() {
    c.device.Off()
}
