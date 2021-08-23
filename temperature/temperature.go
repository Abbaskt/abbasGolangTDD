package temperature

type Temperature float64

func NewCelsius(celsius float64) Temperature {
	return Temperature(celsius)
}

func NewFahrenheit(fahrenheit float64) Temperature {
	return Temperature((fahrenheit - 32.) * (5. / 9.))
}

func NewKelvin(kelvin float64) Temperature {
	return Temperature(kelvin - 273.15)
}

func (temperature Temperature) ToFahrenheit() float64 {
	return (9 * float64(temperature) / 5) + 32
}

func (temperature Temperature) ToKelvin() float64 {
	return float64(temperature) + 273.15
}

func (temperature *Temperature) Add(temp2 Temperature) {
	*temperature += temp2
}
