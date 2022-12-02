package abstract_factory

//import . "abstract_factory/interface"


type Adidas struct {
}
/*
func (a *Adidas) makeShoe() IShoe {
      return  &Shoe{
            Logo: "adidas",
            Size: 14,
        }
    
}

func (a *Adidas) makeShirt() IShirt {
      return  &Shirt{
            Logo: "adidas",
            Size: 14,
        }
}
*/

func (a *Adidas) makeShoe() IShoe {
    return &AdidasShoe{
        Shoe: Shoe{
            Logo: "adidas",
            Size: 14,
        },
    }
}

func (a *Adidas) makeShirt() IShirt {
    return &AdidasShirt{
        Shirt: Shirt{
            Logo: "adidas",
            Size: 14,
        },
    }
}

