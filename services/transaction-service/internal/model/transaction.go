package model

import "time"

type Transaction struct {
	ID          int64
	AccountID   int64
	Amount      int64 //cents
	Description string
	Category    int32
	Type        int32
	Timestamp   time.Time
}
