package builder


import (
	"testing"
	"fmt"
)

func TestBuilder(t *testing.T)  {
    normalBuilder := getBuilder("normal")
    iglooBuilder := getBuilder("igloo")

    director := newDirector(normalBuilder)
    normalHouse := director.buildHouse()

    fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

    director.setBuilder(iglooBuilder)
    iglooHouse := director.buildHouse()

    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
	

}


