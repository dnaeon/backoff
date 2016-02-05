package backoff

import (
	"math"
	"math/rand"
	"time"
)

// Backoff is a time.Duration counter. It starts at Min.
// After every call to Duration() it is  multiplied by Factor.
// It is capped at Max. It returns to Min on every call to Reset().
// Used in conjunction with the time package.
type Backoff struct {
	// Factor is the multiplying factor for each increment step
	Factor float64

	// Jitter eases contention by randomizing backoff steps
	Jitter bool

	// Min and Max are the minimum and maximum values of the counter
	Min, Max time.Duration

	// Number of attempts already made
	attempts int
}

// Returns the current value of the counter and then
// multiplies it by Factor
func (b *Backoff) Duration() time.Duration {
	// Calculate the duration
	d := float64(b.Min) * math.Pow(float64(b.Factor), float64(b.attempts))
	if b.Jitter == true {
		d = rand.Float64()*(d-float64(b.Min)) + float64(b.Min)
	}

	// Capped!
	if d > float64(b.Max) {
		return b.Max
	}

	b.attempts++

	return time.Duration(dur)
}

//Resets the current value of the counter back to Min
func (b *Backoff) Reset() {
	b.attempts = 0
}
