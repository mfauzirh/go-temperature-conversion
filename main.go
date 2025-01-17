package main

import "fmt"

type Celsius struct {
	Temperature float64
}

func (c *Celsius) ToCelsius() float64 {
	return c.Temperature
}

func (c *Celsius) ToFahrenheit() float64 {
	return ((9.0 / 5.0) * c.Temperature) + 32.0
}

func (c *Celsius) ToKelvin() float64 {
	return c.Temperature + 273.15
}

type Fahrenheit struct {
	Temperature float64
}

func (f *Fahrenheit) ToCelsius() float64 {
	return (f.Temperature - 32.0) * (5.0 / 9.0)
}

func (f *Fahrenheit) ToFahrenheit() float64 {
	return f.Temperature
}

func (f *Fahrenheit) ToKelvin() float64 {
	return (f.Temperature + 459.67) * (5.0 / 9.0)
}

type Kelvin struct {
	Temperature float64
}

func (k *Kelvin) ToCelsius() float64 {
	return k.Temperature - 273.15
}

func (k *Kelvin) ToFahrenheit() float64 {
	return (k.Temperature * (9.0 / 5.0)) - 459.67
}

func (k *Kelvin) ToKelvin() float64 {
	return k.Temperature
}

type TemperatureConverter interface {
	ToCelsius() float64
	ToFahrenheit() float64
	ToKelvin() float64
}

func main() {
	fmt.Println("Enter your current temperature")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Print("Enter your choice: ")

	var fromTemperature int
	fmt.Scanf("%d", &fromTemperature)
	for fromTemperature < 1 || fromTemperature > 6 {
		fmt.Print("Choice is not available. Enter your choice: ")
		fmt.Scanf("%d", &fromTemperature)
	}

	fmt.Println("Enter your desired temperature conversion to")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Print("Enter your choice: ")

	var toTemperature int
	fmt.Scanf("%d", &toTemperature)
	for toTemperature < 1 || toTemperature > 6 {
		fmt.Print("Choice is not available. Enter your choice: ")
		fmt.Scanf("%d", &toTemperature)
	}

	var temperature float64
	fmt.Print("\nEnter temperature: ")
	fmt.Scanf("%f", temperature)

	var temperatureConverter TemperatureConverter
	switch fromTemperature {
	case 1:
		temperatureConverter = &Celsius{temperature}
	case 2:
		temperatureConverter = &Fahrenheit{temperature}
	case 3:
		temperatureConverter = &Kelvin{temperature}
	}

	var finalTemperature float64
	switch toTemperature {
	case 1:
		finalTemperature = temperatureConverter.ToCelsius()
	case 2:
		finalTemperature = temperatureConverter.ToFahrenheit()
	case 3:
		finalTemperature = temperatureConverter.ToKelvin()
	}

	fmt.Printf("\nFinal temperature after conversion is: %.2f\n", finalTemperature)

	fmt.Println("\nThanks for your time :)")
}
