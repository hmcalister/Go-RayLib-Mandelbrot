package main

type ComplexNumber struct {
	Re float64
	Im float64
}

func ComplexAddition(a, b ComplexNumber) ComplexNumber {
	return ComplexNumber{
		a.Re + b.Re,
		a.Im + b.Im,
	}
}

func ComplexProduct(a, b ComplexNumber) ComplexNumber {
	return ComplexNumber{
		a.Re*b.Re - a.Im*b.Im,
		a.Re*b.Im + a.Im*b.Re,
	}
}

func ComplexMagnitude(a ComplexNumber) float64 {
	return a.Re*a.Re + a.Im*a.Im
}
