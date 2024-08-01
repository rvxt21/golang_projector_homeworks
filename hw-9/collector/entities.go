package collector

import "time"

type Student struct {
	ID        int
	Name      string
	Surname   string
	Age       uint8
	Group     int
	Marks     []int
	UpdatedAt time.Time
}
