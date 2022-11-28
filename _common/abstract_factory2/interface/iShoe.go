package abstract_factory

type IShoe interface {
    SetLogo(logo string)
    SetSize(size int)
    GetLogo() string
    GetSize() int
}

