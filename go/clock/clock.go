package clock

import "fmt"

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	return Clock{hour: 0, minute: 0}.Add(hour*60 + minute)
}

func (c Clock) Add(minutes int) Clock {
	t := minutes + c.minute
	carry := t / 60
	if t < 0 && t%60 != 0 {
		carry -= 1
	}
	return Clock{hour: (c.hour + carry%24 + 24) % 24, minute: (t%60 + 60) % 60}
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
