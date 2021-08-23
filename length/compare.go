package length

type length float64

const (
	Centimeter length = 1
	Inch              = 2.54 * Centimeter
	Feet              = 12 * Inch
	Meter             = 100 * Centimeter
	Kilometer         = 1000 * Meter
)

func (l length) Centimeter() float64 {
	return float64(l)
}

func (l length) Meter() float64 {
	return float64(l / Meter)
}

func (l length) Kilometer() float64 {
	return float64(l / Kilometer)
}

func (l length) Feet() float64 {
	return float64(l / Feet)
}

func (l length) Inch() float64 {
	return float64(l / Inch)
}
