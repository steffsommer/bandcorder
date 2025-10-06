package testutils

import "time"

type FakeTimeProvider struct {
	Time time.Time
}

func (f *FakeTimeProvider) Now() time.Time {
	return f.Time
}

