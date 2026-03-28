// Package weather provides tools to forcast the weather.
package weather

var (
	// CurrentCondition represents the current condition in the city.
	CurrentCondition string
    // CurrentLocation represents the current location.
	CurrentLocation  string
)

// Forecast will forecase the weather condition, givent he city and the condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
