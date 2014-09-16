package main

type Measure interface {
	convertTo(Measure) *Measure
	convertFrom(Measure) *Measure
}

type BaseMeasure struct {
	Name    string
	Symbol  string
	SiValue float32
}

type Kilograms struct {
	Base BaseMeasure
}

func NewKilograms() *Kilograms {
	return &Kilograms{
		Base: BaseMeasure{
			Name:    "Kilograms",
			Symbol:  "kg",
			SiValue: 1.0,
		},
	}
}
