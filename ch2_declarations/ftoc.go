package ch2declarations

// 1. func: declare a function, with a return type of float64, and a float64 param
// convert the Fahrenheit to Celsius
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}