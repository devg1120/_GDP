package abstract_factory

type IShirt interface {
    SetLogo(logo string)
    SetSize(size int)
    GetLogo() string
    GetSize() int
}

