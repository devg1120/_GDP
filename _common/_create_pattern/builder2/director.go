package builder

import . "builder/interface"

type Director struct {
    builder IBuilder
}

func newDirector(b IBuilder) *Director {
    return &Director{
        builder: b,
    }
}

func (d *Director) setBuilder(b IBuilder) {
    d.builder = b
}

func (d *Director) buildHouse() House {
    d.builder.SetDoorType()
    d.builder.SetWindowType()
    d.builder.SetNumFloor()
    return d.builder.GetHouse()
}
