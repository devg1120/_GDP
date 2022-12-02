package bridge2

import "fmt"
import "_common/bridge2/interface"


type Mac struct {
    printer bridge2.Printer
}

func (m *Mac) Print() {
    fmt.Println("Print request for mac")
    m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p bridge2.Printer) {
    m.printer = p
}
