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

var i1, i2, i3, i4, i5 *interval.TimeInterval
var t1s, t1e time.Time

func init() {
	// interval 2
	v1s := "2014-05-03 00:00 UTC"
	v1e := "2014-05-04 00:00 UTC"
	t1s, _ = time.Parse(timeFormat, v1s)
	t1e, _ = time.Parse(timeFormat, v1e)
	i1, _ = interval.MakeTimeInterval(&t1s, &t1e)

	// interval 2
	v2s := "2014-05-05 00:00 UTC"
	v2e := "2014-05-06 00:00 UTC"
	t2s, _ := time.Parse(timeFormat, v2s)
	t2e, _ := time.Parse(timeFormat, v2e)
	i2, _ = interval.MakeTimeInterval(&t2s, &t2e)

	// interval 3
	v3s := "2014-05-01 00:00 UTC"
	v3e := "2014-05-18 00:00 UTC"
	t3s, _ := time.Parse(timeFormat, v3s)
	t3e, _ := time.Parse(timeFormat, v3e)
	i3, _ = interval.MakeTimeInterval(&t3s, &t3e)

	// interval 4
	v4s := "2014-05-14 00:00 UTC"
	v4e := "2014-05-30 00:00 UTC"
	t4s, _ := time.Parse(timeFormat, v4s)
	t4e, _ := time.Parse(timeFormat, v4e)
	i4, _ = interval.MakeTimeInterval(&t4s, &t4e)

	// interval 5
	v5s := "2014-05-15 00:00 UTC"
	v5e := "2014-05-30 00:00 UTC"
	t5s, _ := time.Parse(timeFormat, v5s)
	t5e, _ := time.Parse(timeFormat, v5e)
	i5, _ = interval.MakeTimeInterval(&t5s, &t5e)
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

func TestTimeInterval_EqualsWithNilParameter(t *testing.T) {
	assert.False(t, i1.Equals(nil))
}

func TestTimeInterval_Equals(t *testing.T) {
	assert.False(t, i1.Equals(i2))
	assert.True(t, i1.Equals(i1))
}

func TestTimeInterval_Meets(t *testing.T) {
	constraint := time.Duration(TwentyFourHours)
	assert.True(t, i2.Meets(i1, constraint))
	assert.False(t, i1.Meets(i1, constraint))
	assert.False(t, i1.Meets(nil, TwentyFourHours))
}

func TestTimeInterval_MetBy(t *testing.T) {
	constraint := time.Duration(TwentyFourHours)
	assert.True(t, i1.MetBy(i2, constraint))
	assert.False(t, i1.MetBy(i1, constraint))
	assert.False(t, i1.MetBy(nil, TwentyFourHours))
}

func TestTimeInterval_Precedes(t *testing.T) {
	// interval 2
	vPrecedeS := "2014-05-06 00:00 UTC"
	vPrecedeE := "2014-05-10 00:00 UTC"
	tPrecedeS, _ := time.Parse(timeFormat, vPrecedeS)
	tPrecedeE, _ := time.Parse(timeFormat, vPrecedeE)
	iPrecede, _ := interval.MakeTimeInterval(&tPrecedeS, &tPrecedeE)
	constraint := time.Duration(TwentyFourHours)
	assert.True(t, i1.Precedes(iPrecede, constraint))
	assert.False(t, i1.Precedes(i2, constraint))
	assert.False(t, i1.Precedes(nil, TwentyFourHours))
}

func TestTimeInterval_Duration(t *testing.T) {
	assert.Equal(t, TwentyFourHours, i1.Duration())
}

func TestTimeInterval_Start(t *testing.T) {
	assert.Equal(t, t1s, *i1.Start())
}

func TestTimeInterval_End(t *testing.T) {
	assert.Equal(t, t1e, *i1.End())
}

func TestTimeInterval_Overlaps(t *testing.T) {
	assert.False(t, i1.Overlaps(nil))
	assert.True(t, i3.Overlaps(i4))
}

func TestTimeInterval_FinishedBy(t *testing.T) {
	assert.True(t, i4.FinishedBy(i5))
	assert.False(t, i1.FinishedBy(nil))
}

func TestTimeInterval_Contains(t *testing.T) {
	assert.True(t, i3.Contains(i2))
	assert.False(t, i1.Contains(nil))
}

