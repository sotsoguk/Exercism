package clock

import "fmt"

type Clock struct {
	h int
	m int
}

func New(h int, m int) Clock {
	newClock := Clock{h: h, m: m}
	newClock.correct()
	return newClock
}

func (c Clock) Add(m int) Clock {
	res := Clock{c.h, c.m + m}
	res.correct()
	return res
}

func (c Clock) Subtract(m int) Clock {
	res := Clock{c.h, c.m - m}
	res.correct()
	return res
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
func (c *Clock) correct() {
	for c.m >= 60 {
		c.m -= 60
		c.h++
	}
	for c.m < 0 {
		c.m += 60
		c.h--
	}
	for c.h < 0 {
		c.h += 24
	}
	c.h %= 24
}
