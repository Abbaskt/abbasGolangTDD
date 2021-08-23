package temperature_test

import (
	"awesomeProject/temperature"
	testifyAssert "github.com/stretchr/testify/assert"
	"testing"
)

func TestTemperatureComparison(t *testing.T) {
	assert := testifyAssert.New(t)
	t.Run("1 degree celsius should be equal to 33.8 ToFahrenheit ", func(t *testing.T) {
		celsius := temperature.NewCelsius(1.)
		fahrenheit := temperature.NewFahrenheit(33.8)

		assert.InDeltaf(float64(celsius), float64(fahrenheit), 0.1, "celsius: %v fahrenheit:%v", celsius, fahrenheit)
	})

	t.Run("100 degree celsius should be equal to 212 ToFahrenheit ", func(t *testing.T) {
		celsius := temperature.NewCelsius(100.)
		fahrenheit := temperature.NewFahrenheit(212.)

		assert.InDeltaf(float64(celsius), float64(fahrenheit), 0.1, "celsius: %v fahrenheit:%v", celsius, fahrenheit)
	})
}

func TestTemperatureConversion(t *testing.T) {
	conversionAssert := testifyAssert.New(t)
	t.Run("33.8 fahrenheit should be 1 degree celsius", func(t *testing.T) {
		celsius := temperature.NewCelsius(1.)
		expectedResultInFahrenheit := 33.8

		valueInFahrenheit := celsius.ToFahrenheit()

		conversionAssert.InDeltaf(valueInFahrenheit, expectedResultInFahrenheit, 0.009, "33.8°F is 1°C")
	})

	t.Run("212 fahrenheit should be 100 degree celsius", func(t *testing.T) {
		celsius := temperature.NewCelsius(100.)
		expectedResultInFahrenheit := 212.

		valueInFahrenheit := celsius.ToFahrenheit()

		conversionAssert.InDeltaf(valueInFahrenheit, expectedResultInFahrenheit, 0.009, "33.8°F is 1°C")
	})

	t.Run("273.15 celsius should be 0 degree kelvin", func(t *testing.T) {
		celsius := temperature.NewCelsius(-273.15)
		expectedResultInFahrenheit := 0.

		valueInKelvin := celsius.ToKelvin()

		conversionAssert.InDeltaf(valueInKelvin, expectedResultInFahrenheit, 0.009, "273.15°C is 0K")
	})

	t.Run("373.15 kelvin should be 100 degree celsius", func(t *testing.T) {
		celsius := temperature.NewCelsius(100.)
		expectedResultInFahrenheit := 373.15

		valueInKelvin := celsius.ToKelvin()

		conversionAssert.InDeltaf(valueInKelvin, expectedResultInFahrenheit, 0.009, "373.15K is 100°C")
	})
}

func TestTemperatureAddition(t *testing.T) {
	t.Run("Expect 2 celsius when 1 celsius is added with 33.8 fahrenheit", func(t *testing.T) {
		celsius := temperature.NewCelsius(1.)
		fahrenheit := temperature.NewFahrenheit(33.8)
		expectedAdditionResult := temperature.NewCelsius(2.)

		celsius.Add(fahrenheit)

		temperatureAssert(t, celsius, expectedAdditionResult, "2 celsius should be returned as addition of 1 celsius and 33.8 fahrenheit")
	})

	t.Run("Expect 2 celsius when 1 celsius is added with 274.15 kelvin", func(t *testing.T) {
		celsius := temperature.NewCelsius(1.)
		kelvin := temperature.NewKelvin(274.15)
		expectedAdditionResult := temperature.NewCelsius(2.)

		celsius.Add(kelvin)

		temperatureAssert(t, celsius, expectedAdditionResult, "2 celsius should be returned as addition of 1 celsius and 274.15 kelvin")
	})
}

func temperatureAssert(t *testing.T, celsius temperature.Temperature, expectedAdditionResult temperature.Temperature, msg string) bool {
	return testifyAssert.InDeltaf(t, float64(celsius), float64(expectedAdditionResult), 0.0009, msg)
}
