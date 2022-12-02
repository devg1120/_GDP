package builder

type IBuilder interface {
    SetWindowType()
    SetDoorType()
    SetNumFloor()
    GetHouse() House
}

