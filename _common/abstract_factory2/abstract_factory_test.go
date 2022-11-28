package abstract_factory

import (
	"testing"
	"fmt"
)

import . "abstract_factory/interface"

func TestCarFactory(t *testing.T)  {
	
    adidasFactory, _ := GetSportsFactory("adidas")
    nikeFactory, _ := GetSportsFactory("nike")

    nikeShoe := nikeFactory.MakeShoe()
    nikeShirt := nikeFactory.MakeShirt()

    adidasShoe := adidasFactory.MakeShoe()
    adidasShirt := adidasFactory.MakeShirt()

    printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)

}

func printShoeDetails(s IShoe) {
    fmt.Printf("Logo: %s", s.GetLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.GetSize())
    fmt.Println()
}

func printShirtDetails(s IShirt) {
    fmt.Printf("Logo: %s", s.GetLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.GetSize())
    fmt.Println()
}

