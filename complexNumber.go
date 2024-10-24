package main

type ComplexNumber struct {
	Re float64
	Im float64
}

func (a ComplexNumber) Add(b ComplexNumber) ComplexNumber {
	return ComplexNumber{
		a.Re + b.Re,
		a.Im + b.Im,
	}
}

func (a ComplexNumber) Product(b ComplexNumber) ComplexNumber {
	return ComplexNumber{
		a.Re*b.Re - a.Im*b.Im,
		a.Re*b.Im + a.Im*b.Re,
	}
}

func (a ComplexNumber) Magnitude() float64 {
	return a.Re*a.Re + a.Im*a.Im
}
