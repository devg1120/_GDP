package bridge2

import "fmt"
import "_common/bridge2/interface"        


type Windows struct {
    printer bridge2.Printer
}

func (w *Windows) Print() {
    fmt.Println("Print request for windows")
    w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p bridge2.Printer) {
    w.printer = p
}
