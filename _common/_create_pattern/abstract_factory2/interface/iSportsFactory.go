package abstract_factory



type ISportsFactory interface {
    MakeShoe() IShoe
    MakeShirt() IShirt
}



