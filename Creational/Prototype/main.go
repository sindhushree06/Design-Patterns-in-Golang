package prototype

import "fmt"

type Prototype struct {
	Value string
}

func (p Prototype) Clone() Prototype {
	return p
}

func main() {
	//creating a new object by copying an existing object
	obj := Prototype{Value: "original"}
	cloneObj := obj.Clone()

	cloneObj.Value = "clone"
	fmt.Println(obj.Value)
	fmt.Println(cloneObj.Value)

}

//value gets printed differently for both original and clone object because they are different objects in memory, even though they have the same value initially.
//this pattern is usefull when we have object with 50 fields and we have to create one more object of similar kind
// In this case we can just use clone
