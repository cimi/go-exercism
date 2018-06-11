package space

import (
	"time"
)

type Planet string

const earthOrbitSeconds int = 31557600

func OrbitDuration(earthYears float64) time.Duration {
	seconds := int(float64(earthOrbitSeconds) * earthYears)
	return time.Duration(seconds) * time.Second
}

var orbits = map[Planet]time.Duration{
	"Earth":   OrbitDuration(1.0),
	"Mercury": OrbitDuration(0.2408467),
	"Venus":   OrbitDuration(0.61519726),
	"Mars":    OrbitDuration(1.8808158),
	"Jupiter": OrbitDuration(11.862615),
	"Saturn":  OrbitDuration(29.447498),
	"Uranus":  OrbitDuration(84.016846),
	"Neptune": OrbitDuration(164.79132)}

func Age(seconds float64, planet Planet) float64 {
	return seconds / orbits[planet].Seconds()
}
