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

// Package interval contains methods to analyse time intervals and relations between two intervals for go.
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
	//Precedes relation
	Precedes State = 0

	// Meets relation
	Meets State = 1

	// Overlaps relation
	Overlaps State = 2

	// FinishedBy relation
	FinishedBy State = 3

	// Contains relation
	Contains State = 4

	// Starts relation
	Starts State = 5

	// Equals relation
	Equals State = 6

	// StartedBy relation
	StartedBy State = 7

	// During relation
	During State = 8

	// Finishes relation
	Finishes State = 9

	// OverlappedBy relation
	OverlappedBy State = 10

	// MetBy relation
	MetBy State = 11

	// PrecededBy relation
	PrecededBy State = 12
)

const (
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

// Equal checks two time intervals are equal or not
func (i *TimeInterval) Equal(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.startTime.Equal(*interval.startTime) && i.endTime.Equal(*interval.endTime) {
		return true
	}
	return false
}

// Meets returns true if interval A meets B
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
func (i *TimeInterval) MetBy(interval *TimeInterval, constraint time.Duration) bool {
	if interval == nil {
		return false
	}
	if result := interval.startTime.Sub(*i.endTime); result != constraint {
		return false
	}
	return true
}

// Precedes returns true if interval A preceds B
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

/* TODO

// Overlaps
// converse relation of OverlappedBy
func (i *TimeInterval) Overlaps(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}

	fmt.Println(i.endTime.Sub(*interval.startTime))

	return false;
}

// FinishedBy
func (i *TimeInterval) FinishedBy(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	return false;
}

// Contains
func (i *TimeInterval) Contains(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	return false;
}

// Starts
func (i *TimeInterval) Starts(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	return false;
}

// StartedBy

// During

// Finishes

// OverlappedBy
// converse relation of OverlappedBy

// MetBy

// PrecededBy
// converse relation of Precedes
*/
