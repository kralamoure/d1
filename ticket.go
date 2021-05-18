package retro

import "time"

type Ticket struct {
	Id           string
	AccountId    string
	GameServerId int
	Created      time.Time
}
