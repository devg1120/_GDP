package up_layer

import "fmt"
import . "up_layer/interface"

type Api_v1 struct {

}

type getData struct {
   Get bool
   No int
   Id    string
   Name  string

}


func (a *Api_v1) Put(p Iup_layer_put_data) {

	fmt.Printf("Put(%s)\n", p.Name)

}

func (a *Api_v1) Get() Iup_layer_get_data {
	fmt.Printf("Ger\n")
        return &getData {
		Get:true,
		No:1,
		Id:"id_001",
		Name:"test name",
	}
}
