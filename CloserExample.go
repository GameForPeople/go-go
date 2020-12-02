package main

import (
	"fmt"
)

type Vector3 struct {
	x int32
	y int32
	z int32
}

func (v Vector3) Length() int32 {
    return v.x + v.y + v.z
}

func (v Vector3) DoMulti() {
    v.x = v.x * v.Length()
}

func (v *Vector3) DoMulti_Ptr() {
    v.x = v.x * v.Length()
}

func (v Vector3) DoDivi() {
    v.y = v.y / v.Length()
}

func (v Vector3) DoDivi_Ptr() {
    v.y = v.y / v.Length()
}

func main() {
	tempVector := Vector3{ 10, 20, 30 }
	
	fmt.Println("Vector : ", tempVector, "Vector의 길이 값은?  : ", tempVector.Length() )

	tempVector.DoMulti()

	fmt.Println("Vector : ", tempVector, "Vector의 길이 값은?  : ", tempVector.Length() )

	tempVector.DoDivi()

	fmt.Println("Vector : ", tempVector, "Vector의 길이 값은?  : ", tempVector.Length() )

	tempVector.DoMulti_Ptr()

	fmt.Println("Vector : ", tempVector, "Vector의 길이 값은?  : ", tempVector.Length() )

	tempVector.DoDivi_Ptr()

	fmt.Println("Vector : ", tempVector, "Vector의 길이 값은?  : ", tempVector.Length() )
}