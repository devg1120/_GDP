package main

import "fmt"
import bi "_common/bridge2/interface"
import b  "_common/bridge2"

type Linux struct {
    printer bi.Printer
}

func (m *Linux) Print() {
    fmt.Println("Print request for linux")
    m.printer.PrintFile()
}

func (m *Linux) SetPrinter(p bi.Printer) {
    m.printer = p
}


func main() {

    hpPrinter := &b.Hp{}
    epsonPrinter := &b.Epson{}
    macComputer := &b.Mac{}

    macComputer.SetPrinter(hpPrinter)
    macComputer.Print()
    fmt.Println()

    macComputer.SetPrinter(epsonPrinter)
    macComputer.Print()
    fmt.Println()

    winComputer := &b.Windows{}

    winComputer.SetPrinter(hpPrinter)
    winComputer.Print()
    fmt.Println()

    winComputer.SetPrinter(epsonPrinter)
    winComputer.Print()
    fmt.Println()
}
