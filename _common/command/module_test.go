package command

import "testing"
//import "fmt"



func TestCommand(t *testing.T) {

    tv := &Tv{}

    onCommand := &OnCommand{
        device: tv,
    }

    offCommand := &OffCommand{
        device: tv,
    }

    onButton := &Button{
        command: onCommand,
    }
    onButton.press()

    offButton := &Button{
        command: offCommand,
    }
    offButton.press()
}
