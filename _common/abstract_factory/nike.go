package abstract_factory

import . "abstract_factory/interface"


type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
     return  &Shoe{
            Logo: "nike",
            Size: 14,
        }
}

func (n *Nike) makeShirt() IShirt {
    return   &Shirt{
            Logo: "nike",
            Size: 14,
        }
}

/*
//func (n *Nike) makeShoe() IShoe {
func (n *Nike) makeShoe() *NikeShoe {
    return &NikeShoe{
        Shoe: Shoe{
            Logo: "nike",
            Size: 14,
        },
    }
}

//func (n *Nike) makeShirt() IShirt {
func (n *Nike) makeShirt() *NikeShirt {
    return &NikeShirt{
        Shirt: Shirt{
            Logo: "nike",
            Size: 14,
        },
    }
}
*/
