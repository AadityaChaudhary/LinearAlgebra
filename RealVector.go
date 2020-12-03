package LinearAlgebra

import "errors"

type RealVector struct {
	nums []float32
}

func NewRealVector(nums []float32) *RealVector {
	if nums == nil {
		return &RealVector{[]float32{}}
	}
	return &RealVector{nums: nums}
}

func(rv RealVector) Get() []float32 {
	return rv.nums
}

func (rv RealVector) Add(other RealVector) (*RealVector,error) {
	if len(rv.nums) != len(other.nums) {
		return nil, errors.New("vectors not in the same set")
	}
	newVec := &RealVector{[]float32{}}

	for key, val := range rv.nums {
		newVec.nums = append(newVec.nums, other.nums[key] + val)
	}
	return newVec,nil

}

func (rv RealVector) Minus(other RealVector) (*RealVector,error) {
	if len(rv.nums) != len(other.nums) {
		return nil, errors.New("vectors not in the same set")
	}
	newVec := &RealVector{[]float32{}}

	for key, val := range rv.nums {
		newVec.nums = append(newVec.nums, val - other.nums[key])
	}
	return newVec,nil

}

func (rv RealVector) Dot(other RealVector) (float32,error) {
	if len(rv.nums) != len(other.nums) {
		return 0, errors.New("vectors not in the same set")
	}

	var prod float32

	for key, val := range rv.nums {
		prod += other.nums[key] * val
	}

	return prod,nil

}

//only works on vectors in R^3
func (rv RealVector) Cross(other RealVector) (*RealVector,error) {
	if len(rv.nums) != 3 || len(other.nums) != 3 {
		return nil, errors.New("atleast one vector is not in R^3")
	}

	NewVec := &RealVector{[]float32{}}

	NewVec.nums = append(NewVec.nums, rv.nums[1]*other.nums[2] - other.nums[1]*rv.nums[2])
	NewVec.nums = append(NewVec.nums, rv.nums[2]*other.nums[0] - other.nums[2]*rv.nums[0])
	NewVec.nums = append(NewVec.nums, rv.nums[0]*other.nums[1] - other.nums[0]*rv.nums[1])

	return NewVec,nil

}

