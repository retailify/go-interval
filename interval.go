/*
Copyright © 2018 Thomas Meitz <info@retailify.de>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions
of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

// Package interval contains methods to analyze time intervals and relations between two intervals for go.
//
// The package is inspired by Allen's interval algebra
// and implements the described 13 basic relations between time intervals.
//
// See also
//
// http://www.ics.uci.edu/~alspaugh/cls/shr/allen.html
package interval

import (
	"errors"
	"fmt"
	"time"
)

// TimeInterval is a struct that contains the start and end time for a interval
type TimeInterval struct {
	startTime *time.Time
	endTime   *time.Time
}

// State type represents the relation type of two TimeInterval's
type State int

const (
	//Unknown relation
	Unknown State = 0

	//Precedes relation
	Precedes State = 1

	// Meets relation
	Meets State = 2

	// Overlaps relation
	Overlaps State = 3

	// FinishedBy relation
	FinishedBy State = 4

	// Contains relation
	Contains State = 5

	// Starts relation
	Starts State = 6

	// Equals relation
	Equals State = 7

	// StartedBy relation
	StartedBy State = 8

	// During relation
	During State = 9

	// Finishes relation
	Finishes State = 10

	// OverlappedBy relation
	OverlappedBy State = 11

	// MetBy relation
	MetBy State = 12

	// PrecededBy relation
	PrecededBy State = 13
)

const (
	// TimeIntervalError error message
	TimeIntervalError = "interval must not nil or unknown error"

	// TimeIntervalEmptyStartTimeError error message
	TimeIntervalEmptyStartTimeError = "start time must not empty"

	// TimeIntervalEmptyEndTimeError error message
	TimeIntervalEmptyEndTimeError = "end time must not empty"
)

// MakeTimeInterval makes a new TimeInterval
func MakeTimeInterval(start, end *time.Time) (*TimeInterval, error) {
	if start == nil {
		return nil, errors.New(TimeIntervalEmptyStartTimeError)
	}
	if end == nil {
		return nil, errors.New(TimeIntervalEmptyEndTimeError)
	}

	return &TimeInterval{startTime: start, endTime: end}, nil
}

// String returns the formatted interval dates
func (i *TimeInterval) String(format string) string {
	return fmt.Sprintf("[%v, %v]", i.startTime.Format(format), i.endTime.Format(format))
}

// Start returns the start time of the interval
func (i *TimeInterval) Start() *time.Time {
	return i.startTime
}

// End returns the end time of the interval
func (i *TimeInterval) End() *time.Time {
	return i.endTime
}

// Duration returns the time.Duration of the interval
func (i *TimeInterval) Duration() time.Duration {
	return i.endTime.Sub(*i.startTime)
}

// Relation return the state of two intervals
func (i *TimeInterval) Relation(interval *TimeInterval, constraint time.Duration) (state State, err error) {

	if interval == nil {
		state, err = Unknown, errors.New(TimeIntervalError)
	} else if i.Precedes(interval, constraint) {
		state, err = Precedes, nil
	} else if i.Meets(interval, constraint) {
		state, err = Meets, nil
	} else if i.Overlaps(interval) {
		state, err = Overlaps, nil
	} else if i.FinishedBy(interval) {
		state, err = FinishedBy, nil
	} else if i.Contains(interval) {
		state, err = Contains, nil
	} else if i.Starts(interval) {
		state, err = Starts, nil
	} else if i.Equals(interval) {
		state, err = Equals, nil
	} else if i.StartedBy(interval) {
		state, err = StartedBy, nil
	} else if i.During(interval) {
		state, err = During, nil
	} else if i.Finishes(interval) {
		state, err = Finishes, nil
	} else if i.OverlappedBy(interval) {
		state, err = OverlappedBy, nil
	} else if i.MetBy(interval, constraint) {
		state, err = MetBy, nil
	} else if i.PrecededBy(interval, constraint) {
		state, err = PrecededBy, nil
	}

	return
}

// Equals checks two time intervals are equal or not
//
// converse relation of Equals is Equals
func (i *TimeInterval) Equals(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.startTime.Equal(*interval.startTime) && i.endTime.Equal(*interval.endTime) {
		return true
	}
	return false
}

// Meets returns true if interval A meets B
//
// converse relation of MetBy
func (i *TimeInterval) Meets(interval *TimeInterval, constraint time.Duration) bool {
	if interval == nil {
		return false
	}
	if result := i.startTime.Sub(*interval.endTime); result != constraint {
		return false
	}
	return true
}

// MetBy returns true if interval A is met by B
//
// converse relation of Met
//
// TODO
func (i *TimeInterval) MetBy(interval *TimeInterval, constraint time.Duration) bool {
	if interval == nil {
		return false
	}
	if result := interval.startTime.Sub(*i.endTime); result != constraint {
		return false
	}
	return true
}

// Precedes returns true if interval A precedes B by duration
//
// converse relation of PrecedesBy
func (i *TimeInterval) Precedes(interval *TimeInterval, constraint time.Duration) bool {
	if interval == nil {
		return false
	}
	if result := interval.startTime.Sub(*i.endTime); result <= constraint {
		return false
	}
	return true
}

// Overlaps returns true if interval A overlaps B
//
// converse relation of OverlappedBy
func (i *TimeInterval) Overlaps(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}

	if i.startTime.Before(*interval.startTime) && interval.startTime.Before(*i.endTime) {
		return true
	}

	return false
}

// FinishedBy returns true if interval A is finished by B
//
// converse relation of OverlappedBy
func (i *TimeInterval) FinishedBy(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.startTime.Before(*interval.startTime) && i.endTime.Equal(*interval.endTime) {
		return true
	}
	return false
}

// Contains returns true if interval A contains B
//
// converse relation to During
func (i *TimeInterval) Contains(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.startTime.Before(*interval.startTime) && i.endTime.After(*interval.endTime) {
		return true
	}
	return false
}

// Starts returns true if interval A's startTime is identical with interval B's startTime and interval A's duration
// is greater than interval B's duration
//
// converse relation of StartedBy
func (i *TimeInterval) Starts(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.startTime.Equal(*interval.startTime) && i.endTime.Before(*interval.endTime) {
		return true
	}
	return false
}

// StartedBy returns true if interval A startTime is identical with interval B startTime and interval A's duration
// is shorter than interval B's duration
//
// converse relation of Starts
func (i *TimeInterval) StartedBy(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.startTime.Equal(*interval.startTime) && i.endTime.After(*interval.endTime) {
		return true
	}
	return false
}

// During returns true if interval A is contained by interval B
//
// converse relation of Contains
func (i *TimeInterval) During(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.startTime.After(*interval.startTime) && i.endTime.Before(*interval.endTime) {
		return true
	}
	return false
}

// Finishes returns true if interval A's endTime is identical with B's endTime and B's startTime is greater than
// A's startTime
//
// converse relation of FinishedBy
func (i *TimeInterval) Finishes(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.startTime.After(*interval.startTime) && i.endTime.Equal(*interval.endTime) {
		return true
	}
	return false
}

// OverlappedBy returns true if interval A's startTime is contained in interval B and greater than B's startTime
// and A's endTime is greater than B's endTime
//
// converse relation to Overlaps
func (i *TimeInterval) OverlappedBy(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}

	if i.startTime.After(*interval.startTime) && i.endTime.After(*interval.endTime) {
		return true
	}

	return false
}

// PrecededBy returns true if interval A is preceded by B. That's the case if interval B's endTime + constraint is
// greater than A's startTime
//
// converse relation of Precedes
func (i *TimeInterval) PrecededBy(interval *TimeInterval, constraint time.Duration) bool {
	if interval == nil {
		return false
	}
	if result := i.startTime.Sub(*interval.endTime); result <= constraint {
		return false
	}
	return true
}
