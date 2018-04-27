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

package interval_test

import (
	"testing"
	"time"

	"fmt"

	interval "github.com/retailify/go-interval"
	"github.com/stretchr/testify/assert"
)

const (
	timeFormat      = "2006-01-02 15:04 MST"
	TwentyFourHours = 24 * time.Hour
)

var i1, i2, i3, i4, i5, i6, i7 *interval.TimeInterval
var t1s, t1e time.Time

func init() {
	i1, _ = interval.MakeTimeIntervalFromStrings("2014-05-03 00:00 UTC", "2014-05-04 00:00 UTC", timeFormat)
	i2, _ = interval.MakeTimeIntervalFromStrings("2014-05-05 00:00 UTC", "2014-05-06 00:00 UTC", timeFormat)
	i3, _ = interval.MakeTimeIntervalFromStrings("2014-05-01 00:00 UTC", "2014-05-18 00:00 UTC", timeFormat)
	i4, _ = interval.MakeTimeIntervalFromStrings("2014-05-14 00:00 UTC", "2014-05-30 00:00 UTC", timeFormat)
	i5, _ = interval.MakeTimeIntervalFromStrings("2014-05-15 00:00 UTC", "2014-05-30 00:00 UTC", timeFormat)
	i6, _ = interval.MakeTimeIntervalFromStrings("2014-05-15 00:00 UTC", "2014-05-25 00:00 UTC", timeFormat)
	i7, _ = interval.MakeTimeIntervalFromStrings("2014-05-06 00:00 UTC", "2014-05-10 00:00 UTC", timeFormat)
}

func TestMakeTimeIntervalWithEmptyStartTime(t *testing.T) {
	_, err := interval.MakeTimeInterval(nil, nil)
	assert.EqualError(t, err, interval.TimeIntervalEmptyStartTimeError)
}

func TestMakeTimeIntervalWithEmptyEndTime(t *testing.T) {
	startTime := time.Now()
	_, err := interval.MakeTimeInterval(&startTime, nil)
	assert.EqualError(t, err, interval.TimeIntervalEmptyEndTimeError)
}

func TestTimeInterval_String(t *testing.T) {
	assert.Equal(t, "[2014-05-03T00:00:00Z, 2014-05-04T00:00:00Z]", i1.String(time.RFC3339))
}

func ExampleTimeInterval_String() {
	v1s := "2014-05-03 00:00 UTC"
	v1e := "2014-05-04 00:00 UTC"
	startTime, _ := time.Parse(timeFormat, v1s)
	endTime, _ := time.Parse(timeFormat, v1e)
	interval, _ := interval.MakeTimeInterval(&startTime, &endTime)
	fmt.Println(interval.String(time.RFC3339))
	// prints [2014-05-03T00:00:00Z .. 2014-05-04T00:00:00Z]
}

func TestMakeInterval(t *testing.T) {
	startTime, endTime := time.Now(), time.Now()
	interval, err := interval.MakeTimeInterval(&startTime, &endTime)
	assert.NoError(t, err)
	assert.Equal(t, &startTime, interval.Start())
	assert.Equal(t, &endTime, interval.End())
}

func ExampleMakeTimeInterval() {
	startTime, endTime := time.Now(), time.Now()
	interval.MakeTimeInterval(&startTime, &endTime)
}

func TestTimeInterval_Equals(t *testing.T) {
	assert.False(t, i1.Equals(i2))
	assert.True(t, i1.Equals(i1))
	assert.False(t, i1.Equals(nil))
	state, _ := i1.Relation(nil, time.Duration(0))
	assert.Equal(t, interval.Unknown, state)
	state, _ = i1.Relation(i1, time.Duration(0))
	assert.Equal(t, interval.Equals, state)
}

func TestTimeInterval_Meets(t *testing.T) {
	constraint := time.Duration(TwentyFourHours)
	assert.True(t, i1.Meets(i2, constraint))
	assert.False(t, i1.Meets(i1, constraint))
	assert.False(t, i1.Meets(nil, TwentyFourHours))

	state, _ := i1.Relation(i2, constraint)
	assert.Equal(t, interval.Meets, state)
}

func TestTimeInterval_MetBy(t *testing.T) {
	constraint := time.Duration(TwentyFourHours)
	assert.True(t, i2.MetBy(i1, constraint))
	assert.False(t, i1.MetBy(i1, constraint))
	assert.False(t, i1.MetBy(nil, TwentyFourHours))

	state, _ := i2.Relation(i1, constraint)
	assert.Equal(t, interval.MetBy, state)
}

func TestTimeInterval_Precedes(t *testing.T) {

	constraint := time.Duration(TwentyFourHours)

	assert.True(t, i1.Precedes(i4, constraint))
	assert.True(t, i1.Precedes(i5, constraint))
	assert.True(t, i1.Precedes(i6, constraint))
	assert.True(t, i1.Precedes(i7, constraint))
	assert.False(t, i1.Precedes(i2, constraint))
	assert.False(t, i1.Precedes(nil, TwentyFourHours))

	state, _ := i1.Relation(i7, constraint)
	assert.Equal(t, interval.Precedes, state)
}

func TestTimeInterval_PrecededBy(t *testing.T) {
	constraint := time.Duration(TwentyFourHours)

	assert.True(t, i7.PrecededBy(i1, constraint))
	state, _ := i7.Relation(i1, constraint)
	assert.Equal(t, interval.PrecededBy, state)

	assert.False(t, i2.PrecededBy(i1, constraint))
	assert.False(t, i1.PrecededBy(nil, TwentyFourHours))

}

func TestTimeInterval_Duration(t *testing.T) {
	assert.Equal(t, TwentyFourHours, i1.Duration())
}

func TestTimeInterval_Start(t *testing.T) {
	t1s, _ := time.Parse(timeFormat, "2014-05-03 00:00 UTC")
	assert.Equal(t, t1s, *i1.Start())
}

func TestTimeInterval_End(t *testing.T) {
	t1e, _ := time.Parse(timeFormat, "2014-05-04 00:00 UTC")
	assert.Equal(t, t1e, *i1.End())
}

func TestTimeInterval_Overlaps(t *testing.T) {
	assert.False(t, i1.Overlaps(nil))
	assert.True(t, i3.Overlaps(i4))
	assert.True(t, i3.Overlaps(i5))
	assert.True(t, i3.Overlaps(i6))
	assert.False(t, i4.Overlaps(i3))
	state, _ := i3.Relation(i4, time.Duration(0))
	assert.Equal(t, interval.Overlaps, state)
}

func TestTimeInterval_OverlappedBy(t *testing.T) {
	assert.False(t, i1.OverlappedBy(nil))
	assert.False(t, i3.OverlappedBy(i4))
	assert.True(t, i4.OverlappedBy(i3))
	state, _ := i4.Relation(i3, time.Duration(0))
	assert.Equal(t, interval.OverlappedBy, state)
}

func TestTimeInterval_Finishes(t *testing.T) {
	assert.True(t, i5.Finishes(i4))
	assert.False(t, i4.Finishes(i5))
	assert.False(t, i1.Finishes(nil))
	state, _ := i5.Relation(i4, time.Duration(0))
	assert.Equal(t, interval.Finishes, state)
}

func TestTimeInterval_FinishedBy(t *testing.T) {
	assert.True(t, i4.FinishedBy(i5))
	state, _ := i4.Relation(i5, time.Duration(0))
	assert.Equal(t, interval.FinishedBy, state)

	assert.False(t, i5.FinishedBy(i4))
	assert.False(t, i1.FinishedBy(nil))
}

func TestTimeInterval_Contains(t *testing.T) {
	assert.True(t, i3.Contains(i2))
	state, _ := i3.Relation(i2, time.Duration(0))
	assert.Equal(t, interval.Contains, state)

	assert.False(t, i2.Contains(i3))
	assert.False(t, i1.Contains(nil))
}

func TestTimeInterval_During(t *testing.T) {
	assert.True(t, i2.During(i3))
	assert.False(t, i3.During(i2))
	assert.False(t, i1.During(nil))
	state, _ := i2.Relation(i3, time.Duration(0))
	assert.Equal(t, interval.During, state)
}

func TestTimeInterval_Starts(t *testing.T) {
	assert.False(t, i1.Starts(nil))
	assert.True(t, i6.Starts(i5))
	assert.False(t, i5.Starts(i6))
	state, _ := i6.Relation(i5, time.Duration(0))
	assert.Equal(t, interval.Starts, state)
}

func TestTimeInterval_StartedBy(t *testing.T) {
	assert.False(t, i1.StartedBy(nil))
	assert.True(t, i5.StartedBy(i6))
	assert.False(t, i6.StartedBy(i5))
	state, _ := i5.Relation(i6, time.Duration(0))
	assert.Equal(t, interval.StartedBy, state)
}

func TestMakeTimeIntervalFromStrings(t *testing.T) {
	i, err := interval.MakeTimeIntervalFromStrings("2014-05-03 00:00 UTC", "2014-05-04 00:00 UTC", timeFormat)
	assert.NoError(t, err)
	assert.Equal(t, i1.Start(), i.Start())
	_, err = interval.MakeTimeIntervalFromStrings("123", "123", timeFormat)
	assert.Error(t, err)
	_, err = interval.MakeTimeIntervalFromStrings("2014-05-03 00:00 UTC", "123", timeFormat)
	assert.Error(t, err)

}

func ExampleTimeInterval_Relation() {
	var iv1, iv2 *interval.TimeInterval
	iv1, _ = interval.MakeTimeIntervalFromStrings(
		"2014-05-03 00:00 UTC", "2014-05-04 00:00 UTC", "2006-01-02 15:04 MST")
	iv2, _ = interval.MakeTimeIntervalFromStrings(
		"2014-05-15 00:00 UTC", "2014-05-30 00:00 UTC", "2006-01-02 15:04 MST")

	if state, err := iv1.Relation(iv2, time.Duration(0)); err == nil && state == interval.Equals {
		fmt.Println("intervals are idendtical")
	}
}
