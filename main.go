package main

func main() {

}

// Convert temperature from Celsius to Fahrenheit.
// Parameters:
// - temperatureCelsius: temperature in Celsius.
// Returns:
// - The temperature converted to Fahrenheit.
func CelsiusToFahrenheit(temperatureCelsius float64) float64 {
	temperatureFahrenheit := (9.0 / 5.0 * temperatureCelsius) + 32.0
	return temperatureFahrenheit
}

// Convert temperature from Celsius to Kelvin.
// Parameters:
// - temperatureCelsius: temperature in Celsius.
// Returns:
// - the temperature converted to Kelvin.
func CelsiusToKelvin(temperatureCelsius float64) float64 {
	tKelvin := temperatureCelsius + 273.15
	return tKelvin
}
