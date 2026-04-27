// Package weather provides tools for weather forecasting.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current city.
var CurrentLocation  string

/* Forecast accepts the current city and the current weather condition
	and then returns a string value equal to the current weather condition
    in the city.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
