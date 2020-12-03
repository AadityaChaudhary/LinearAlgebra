package LinearAlgebra


const epsilon = 1e-8
//low precision equality: 1e-4
//medium: 1e-8
//high: 1e-12

func equal(a,b float64) bool {
	return (a-b) < epsilon && (b-a) < epsilon
}


