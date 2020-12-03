package LinearAlgebra

import (
	"errors"
	"fmt"
	"math"
)

//type RealVector struct {
//	nums []big.Float
//}
//
//func NewRealVector(nums []big.Float) *RealVector {
//	if nums == nil {
//		return &RealVector{[]big.Float{}}
//	}
//	return &RealVector{nums: nums}
//}
//
//func(v RealVector) Get() []big.Float {
//	return v.nums
//}
//
//func (v RealVector) Add(other *RealVector) (*RealVector,error) {
//	if len(v.nums) != len(other.nums) {
//		return nil, errors.New("vectors not in the same set")
//	}
//	newVec := &RealVector{[]big.Float{}}
//
//	for key, val := range v.nums {
//		var temp big.Float
//		newVec.nums = append(newVec.nums, *temp.Add(&other.nums[key],&val))
//	}
//	return newVec,nil
//
//}
//
//func (v RealVector) Minus(other *RealVector) (*RealVector,error) {
//	if len(v.nums) != len(other.nums) {
//		return nil, errors.New("vectors not in the same set")
//	}
//	newVec := &RealVector{[]big.Float{}}
//
//	for key, val := range v.nums {
//		var temp big.Float
//		newVec.nums = append(newVec.nums, *temp.Sub(&val, &other.nums[key]))
//	}
//	return newVec,nil
//
//}
//
//func (v RealVector) Dot(other *RealVector) (float32,error) {
//	if len(v.nums) != len(other.nums) {
//		return 0, errors.New("vectors not in the same set")
//	}
//
//	var prod big.Float
//
//	for key, val := range v.nums {
//		var z
//		prod.Add(prod,)
//	}
//
//	return prod,nil
//
//}
//
////only works on vectors in R^3
//func (v RealVector) Cross(other *RealVector) (*RealVector,error) {
//	if len(v.nums) != 3 || len(other.nums) != 3 {
//		return nil, errors.New("atleast one vector is not in R^3")
//	}
//
//	NewVec := &RealVector{[]float32{}}
//
//	NewVec.nums = append(NewVec.nums, v.nums[1]*other.nums[2] - other.nums[1]*v.nums[2])
//	NewVec.nums = append(NewVec.nums, v.nums[2]*other.nums[0] - other.nums[2]*v.nums[0])
//	NewVec.nums = append(NewVec.nums, v.nums[0]*other.nums[1] - other.nums[0]*v.nums[1])
//
//	return NewVec,nil
//}
//
//func (v RealVector) Equal(other *RealVector) bool {
//	if len(v.nums) != len(other.nums) {
//		return false
//	}
//	for key, val := range v.nums {
//		if val != other.nums[key] {
//			fmt.Println(val," : ",other.nums[key])
//			return false
//		}
//	}
//	return true
//}
//
//func(v RealVector) String() string {
//	sVec := "["
//	for _, val := range v.nums {
//		sVec += fmt.Sprintf(" %f,",val)
//	}
//	sVec = sVec[0:len(sVec)-1]
//	sVec += " ]"
//	return sVec
//}

type RealVector struct {
	nums []float64
}

func NewRealVector(nums []float64) *RealVector {
	if nums == nil {
		return &RealVector{[]float64{}}
	}
	return &RealVector{nums: nums}
}

func(v RealVector) Get() []float64 {
	return v.nums
}

func (v RealVector) Add(other *RealVector) (*RealVector,error) {
	if len(v.nums) != len(other.nums) {
		return nil, errors.New("vectors not in the same set")
	}
	newVec := &RealVector{[]float64{}}

	for key, val := range v.nums {

		newVec.nums = append(newVec.nums, other.nums[key]+ val)
	}
	return newVec,nil

}

func (v RealVector) Minus(other *RealVector) (*RealVector,error) {
	if len(v.nums) != len(other.nums) {
		return nil, errors.New("vectors not in the same set")
	}
	newVec := &RealVector{[]float64{}}

	for key, val := range v.nums {

		newVec.nums = append(newVec.nums, val - other.nums[key])
	}
	return newVec,nil

}

func (v RealVector) Dot(other *RealVector) (float64,error) {
	if len(v.nums) != len(other.nums) {
		return 0, errors.New("vectors not in the same set")
	}

	var prod float64

	for key, val := range v.nums {

		prod += val * other.nums[key]
	}

	return prod,nil

}

//only works on vectors in R^3
func (v RealVector) Cross(other *RealVector) (*RealVector,error) {
	if len(v.nums) != 3 || len(other.nums) != 3 {
		return nil, errors.New("atleast one vector is not in R^3")
	}

	NewVec := &RealVector{[]float64{}}

	NewVec.nums = append(NewVec.nums, v.nums[1]*other.nums[2] - other.nums[1]*v.nums[2])
	NewVec.nums = append(NewVec.nums, v.nums[2]*other.nums[0] - other.nums[2]*v.nums[0])
	NewVec.nums = append(NewVec.nums, v.nums[0]*other.nums[1] - other.nums[0]*v.nums[1])

	return NewVec,nil
}

func (v RealVector) Equal(other *RealVector) bool {
	if len(v.nums) != len(other.nums) {
		return false
	}
	for key, val := range v.nums {
		if !equal(val,other.nums[key]) {
			fmt.Println(val," : ",other.nums[key])
			return false
		}
	}
	return true
}

func(v RealVector) Norm() float64 {
	var norm float64
	for _, val := range v.nums {
		norm += val*val
	}
	return math.Sqrt(norm)
}

func(v RealVector) String() string {
	sVec := "["
	for _, val := range v.nums {
		sVec += fmt.Sprintf(" %f,",val)
	}
	sVec = sVec[0:len(sVec)-1]
	sVec += " ]"
	return sVec
}

