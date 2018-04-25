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

package interval

import (
	"time"
	"errors"
)

type TimeInterval struct {
	Start *time.Time
	End *time.Time
}

const EmptyStartTimeError = "start time must not empty"
const EmptyEndTimeError = "end time must not empty"

// MakeTimeInterval makes a new TimeInterval
func MakeTimeInterval(start, end *time.Time) (*TimeInterval, error){
	if start == nil {
		return nil, errors.New(EmptyStartTimeError)
	}
	if end == nil {
		return nil, errors.New(EmptyEndTimeError)
	}

	return &TimeInterval{Start:start, End: end}, nil
}

// Equal checks two time intervals are equal or not
func (i *TimeInterval) Equal(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	if i.Start.Equal(*interval.Start) && i.End.Equal(*interval.End) {
		return true
	}
	return false;
}

/* TODO
// Precedes
// converse relation of PrecedesBy
func (i *TimeInterval) Precedes(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	return false;
}

// Meets
// converse relation of MeetsBy
func (i *TimeInterval) Meets(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
	return false;
}

// Overlaps
// converse relation of OverlappedBy
func (i *TimeInterval) Overlaps(interval *TimeInterval) bool {
	if interval == nil {
		return false
	}
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