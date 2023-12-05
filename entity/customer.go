package entity

import "time"

type Customer struct {
	Id            string
	Name          string
	Phone         string
	Active_member bool
	Join_date     time.Time
	Gender        string
}
