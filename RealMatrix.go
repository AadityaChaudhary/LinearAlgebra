package LinearAlgebra

import "errors"

type RealMatrix struct {
	vecs []*RealVector
}

func NewRealMatrix(vectors ...*RealVector) (*RealMatrix, error) {
	rows := 0
	for _, v := range vectors {
		if rows == 0 {
			rows = len(v.nums)
		}
		if rows != len(v.nums) {
			return nil, errors.New("vectors must all have the same size")
		}
	}

	return &RealMatrix{vectors}, nil
}

func(m RealMatrix) GetCols() int {
	return len(m.vecs)
}

func(m RealMatrix) GetRows() int {
	if len(m.vecs) > 0 {
		return len(m.vecs[0].nums)
	}
	return 0
}

func(m RealMatrix) Add(other *RealMatrix) (*RealMatrix,error){

	if m.GetCols() != other.GetCols() || m.GetRows() == other.GetRows() {
		return nil, errors.New("matricies need to have the same dimensions")
	}

	newMatrix := &RealMatrix{}

	for key, val := range m.vecs {
		v, err := val.Add(other.vecs[key])
		if err != nil {
			//do something
			return nil, errors.New("encountered errors when adding vectors")
		}

		newMatrix.vecs = append(newMatrix.vecs, v)
	}
	return newMatrix, nil
}

func(m RealMatrix) Minus(other *RealMatrix) (*RealMatrix,error){

	if m.GetCols() != other.GetCols() || m.GetRows() == other.GetRows() {
		return nil, errors.New("matricies need to have the same dimensions")
	}

	newMatrix := &RealMatrix{}

	for key, val := range m.vecs {
		v, err := val.Minus(other.vecs[key])
		if err != nil {
			//do something
			return nil, errors.New("encountered errors when subtracting vectors")
		}

		newMatrix.vecs = append(newMatrix.vecs, v)
	}
	return newMatrix, nil
}


