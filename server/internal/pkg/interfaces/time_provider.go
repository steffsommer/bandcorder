package interfaces

import "time"

type TimeProvider interface {
	Now() time.Time
}
