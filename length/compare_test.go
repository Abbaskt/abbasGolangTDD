package length

import (
	testifyAssert "github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthConversion(t *testing.T) {
	t.Run("Compare Meters and Centimeters", func(t *testing.T) {
		centimeter := 100 * Centimeter
		meter := Meter

		assertEquals(t, centimeter, meter)
	})

	t.Run("Compare Kilometer and Centimeter", func(t *testing.T) {
		centimeter := 100000 * Centimeter
		kilometer := Kilometer

		assertEquals(t, centimeter, kilometer)
	})

	t.Run("Compare Inches and Feet", func(t *testing.T) {
		inches := 12 * Inch
		feet := Feet

		assertEquals(t, inches, feet)
	})

	t.Run("Compare Inches and Meter", func(t *testing.T) {
		inches := 39.3701 * Inch
		meter := Meter

		assertEquals(t, inches, meter)
	})
}

func assertEquals(t *testing.T, expected length, actual length) bool {
	return testifyAssert.InDeltaf(t, float64(expected), float64(actual), 0.0001, "")
}

func TestConvert(t *testing.T) {
	assert := testifyAssert.New(t)
	t.Run("Convert meter to 100 centimeter", func(t *testing.T) {
		var meter = Meter
		expectedCentimeter := 100.

		resultInCentimeter := meter.Centimeter()

		assert.Equal(expectedCentimeter, resultInCentimeter, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert kilometer to 100000 centimeter", func(t *testing.T) {
		var kilometer = Kilometer
		expectedCentimeter := 100000.

		resultInCentimeter := kilometer.Centimeter()

		assert.Equal(expectedCentimeter, resultInCentimeter, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert kilometer to 1000 meter", func(t *testing.T) {
		var kilometer = Kilometer
		expectedMeter := 1000.

		resultInCentimeter := kilometer.Meter()

		assert.Equal(expectedMeter, resultInCentimeter, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert kilometer to 5000 meter", func(t *testing.T) {
		var kilometer = 5 * Kilometer
		expectedMeter := 5000.

		resultInCentimeter := kilometer.Meter()

		assert.Equal(expectedMeter, resultInCentimeter, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert 1000 meter to kilometer", func(t *testing.T) {
		var thousandMeter = 1000 * Meter
		expectedKilometer := 1.

		resultInKilometer := thousandMeter.Kilometer()

		assert.Equal(expectedKilometer, resultInKilometer, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert 1000 centimeter to kilometer", func(t *testing.T) {
		var centimeter = 1000 * Centimeter
		expectedKilometer := 5000.

		resultInKilometer := centimeter.Kilometer()

		assert.Equal(expectedKilometer, resultInKilometer, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert 1000 centimeter to inch", func(t *testing.T) {
		var thousandCentimeter = 1000 * Centimeter
		expectedInch := 393.7

		resultInInch := thousandCentimeter.Inch()

		assert.InDeltaf(expectedInch, resultInInch, 0.01, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert feet to inch", func(t *testing.T) {
		var feet = Feet
		expectedInch := 12.

		resultInInch := feet.Inch()

		assert.InDeltaf(expectedInch, resultInInch, 0.01, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert 60 centimeter to feet", func(t *testing.T) {
		var sixtyCentimeter = 60 * Centimeter
		expectedFeet := 1.96

		resultInFeet := sixtyCentimeter.Feet()

		assert.InDeltaf(expectedFeet, resultInFeet, 0.01, "100 centimeter should be equal to 1 meter")
	})

	t.Run("Convert 100 Meter to feet", func(t *testing.T) {
		var hundredMeter = 100 * Meter
		expectedFeet := 328.084

		resultInFeet := hundredMeter.Feet()

		assert.InDeltaf(expectedFeet, resultInFeet, 0.01, "100 hundredMeter should be equal to 1 meter")
	})
}
