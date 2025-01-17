package main

import "fmt"

func main() {
	fmt.Println("Enter desired temperature conversion")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Celsius to Kelvin")
	fmt.Println("3. Fahrenheit to Celsius")
	fmt.Println("4. Fahrenheit to Kelvin")
	fmt.Println("5. Kelvin to Celsius")
	fmt.Println("6. Kelvin to Fahrenheit")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scanf("%d", &choice)
	for choice < 1 || choice > 6 {
		fmt.Print("Temperature conversion not available. Enter your choice: ")
		fmt.Scanf("%d", &choice)
	}

	var temperature float64
	fmt.Print("\nEnter temperature: ")
	fmt.Scanf("%f", temperature)

	var finalTemperature float64
	switch choice {
	case 1:
		finalTemperature = CelsiusToFahrenheit(temperature)
	case 2:
		finalTemperature = CelsiusToKelvin(temperature)
	case 3:
		finalTemperature = FahrenheitToCelsius(temperature)
	case 4:
		finalTemperature = FahrenheitToKelvin(temperature)
	case 5:
		finalTemperature = KelvinToCelsius(temperature)
	case 6:
		finalTemperature = KelvinToFahrenheit(temperature)
	}

	fmt.Printf("\nFinal temperature after conversion is: %.2f\n", finalTemperature)

	fmt.Println("\nThanks for your time :)")
}

// Convert temperature from Celsius to Fahrenheit.
//
// Parameters:
//
// - temperatureCelsius: temperature in Celsius.
//
// Returns:
//
// - The temperature converted to Fahrenheit.
func CelsiusToFahrenheit(temperatureCelsius float64) float64 {
	temperatureFahrenheit := ((9.0 / 5.0) * temperatureCelsius) + 32.0
	return temperatureFahrenheit
}

// Convert temperature from Celsius to Kelvin.
//
// Parameters:
//
// - temperatureCelsius: temperature in Celsius.
//
// Returns:
//
// - the temperature converted to Kelvin.
func CelsiusToKelvin(temperatureCelsius float64) float64 {
	tKelvin := temperatureCelsius + 273.15
	return tKelvin
}

// Convert temperature from Fahrenheit to Celsius.
//
// Parameters:
//
// - temperatureFahrenheit: temperature in Fahrenheit.
//
// Returns:
//
// - the temperature converted to Celsius.
func FahrenheitToCelsius(temperatureFahrenheit float64) float64 {
	temperatureCelsius := (temperatureFahrenheit - 32.0) * (5.0 / 9.0)
	return temperatureCelsius
}

// Convert temperature from Fahrenheit to Kelvin.
//
// Parameters:
//
// - temperatureFahrenheit: temperature in Fahrenheit.
//
// Returns:
//
// - the temperature converted to Kelvin.
func FahrenheitToKelvin(temperatureFahrenheit float64) float64 {
	temperatureKelvin := (temperatureFahrenheit + 459.67) * (5.0 / 9.0)
	return temperatureKelvin
}

// Convert temperature from Kelvin to Celsius.
//
// Parameters:
//
// - temperatureKelvin: temperature in Kelvin.
//
// Returns:
//
// - the temperature converted to Celsius.
func KelvinToCelsius(temperatureKelvin float64) float64 {
	temperatureCelsius := temperatureKelvin - 273.15
	return temperatureCelsius
}

// Convert temperature from Kelvin to Fahrenheit.
//
// Parameters:
//
// - temperatureKelvin: temperature in Kelvin.
//
// Returns:
//
// - the temperature converted to Fahrenheit.
func KelvinToFahrenheit(temperatureKelvin float64) float64 {
	temperatureFahrenheit := (temperatureKelvin * (9.0 / 5.0)) - 459.67
	return temperatureFahrenheit
}
