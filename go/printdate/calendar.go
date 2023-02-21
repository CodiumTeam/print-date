package printdate

import "time"

type Calendar interface {
	today() time.Time
}

func NewCalendar() Calendar {
	return calendar{}
}

type calendar struct {
}

func (calendar calendar) today() time.Time {
	return time.Now()
}
