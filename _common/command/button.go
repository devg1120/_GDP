package command

import . "command/interface"

type Button struct {
    command Command
}

func (b *Button) press() {
    b.command.Execute()
}
