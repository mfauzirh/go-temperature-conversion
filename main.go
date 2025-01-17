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

// Convert temperature from Fahrenheit to Celsius.
// Parameters:
// - temperatureFahrenheit: temperature in Fahrenheit.
// Returns:
// - the temperature converted to Celsius.
func FahrenheitToCelsius(temperatureFahrenheit float64) float64 {
	temperatureCelsius := (temperatureFahrenheit - 32.0) * (5.0 / 9.0)
	return temperatureCelsius
}

// Convert temperature from Fahrenheit to Kelvin.
// Parameters:
// - temperatureFahrenheit: temperature in Fahrenheit.
// Returns:
// - the temperature converted to Kelvin.
func FahrenheitToKelvin(temperatureFahrenheit float64) float64 {
	temperatureKelvin := (temperatureFahrenheit + 459.67) * (5.0 / 9.0)
	return temperatureKelvin
}

// Convert temperature from Kelvin to Celsius.
// Parameters:
// - temperatureKelvin: temperature in Kelvin.
// Returns:
// - the temperature converted to Celsius.
func KelvinToCelsius(temperatureKelvin float64) float64 {
	temperatureCelsius := temperatureKelvin - 273.15
	return temperatureCelsius
}

// Convert temperature from Kelvin to Fahrenheit.
// Parameters:
// - temperatureKelvin: temperature in Kelvin.
// Returns:
// - the temperature converted to Fahrenheit.
func KelvinToFahrenheit(temperatureKelvin float64) float64 {
	temperatureFahrenheit := (temperatureKelvin * (9.0 / 5.0)) - 459.67
	return temperatureFahrenheit
}
