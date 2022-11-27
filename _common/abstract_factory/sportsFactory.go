package abstract_factory

import "fmt"
import . "abstract_factory/interface"


/*
type ISportsFactory interface {
    makeShoe() IShoe
    makeShirt() IShirt
}
*/

func GetSportsFactory(brand string) (ISportsFactory, error) {

    if brand == "adidas" {
        return &Adidas{}, nil
    }

    if brand == "nike" {
        return &Nike{}, nil
    }

/*
    if brand == "adidas" {
        return new(Adidas), nil
    }

    if brand == "nike" {
        return new(Nike), nil
    }
    */
    return nil, fmt.Errorf("Wrong brand type passed")
}
