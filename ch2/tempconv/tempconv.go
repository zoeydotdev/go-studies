// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Kelvin float64
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroK	Kelvin = 0
	FreezingK			Kelvin = 273.15
	BoilingK			Kelvin = 373.15
)

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC			Celsius = 0
	BoilingC			Celsius = 100
)

func (k Kelvin) String() string 		{ return fmt.Sprintf("%.5g°K", k) }
func (c Celsius) String() string		{ return fmt.Sprintf("%.5g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.5g°F", f) }