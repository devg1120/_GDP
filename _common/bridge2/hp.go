package bridge2

import "fmt"

type Hp struct {
}

func (p *Hp) PrintFile() {
    fmt.Println("Printing by a HP Printer")
}
