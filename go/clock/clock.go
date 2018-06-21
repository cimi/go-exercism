package clock

import "fmt"

type Clock struct {
	hour, minute int
}

func cap(num, max int) int {
	return (num%max + max) % max
}

func hours(minute int) int {
	carry := minute / 60
	if minute < 0 && minute%60 != 0 {
		carry -= 1
	}
	return carry
}
func New(hour, minute int) Clock {
	hour += hours(minute)
	return Clock{hour: cap(hour, 24), minute: cap(minute, 60)}
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
