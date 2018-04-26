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
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04 MST"
	TwentyFourHours = 24 * time.Hour
)

var i1, i2 *TimeInterval

func init() {
	// interval 2
	v1s := "2014-05-03 00:00 UTC"
	v1e := "2014-05-04 00:00 UTC"
	t1s, _ := time.Parse(timeFormat, v1s)
	t1e, _ := time.Parse(timeFormat, v1e)
	i1, _ = MakeTimeInterval(&t1s, &t1e)

	// interval 2
	v2s := "2014-05-05 00:00 UTC"
	v2e := "2014-05-06 00:00 UTC"
	t2s, _ := time.Parse(timeFormat, v2s)
	t2e, _ := time.Parse(timeFormat, v2e)
	i2, _ = MakeTimeInterval(&t2s, &t2e)
}

func TestMakeTimeIntervalWithEmptyStartTime(t *testing.T) {
	_, err := MakeTimeInterval(nil, nil)
	assert.EqualError(t, err, EmptyStartTimeError)
}

func TestMakeTimeIntervalWithEmptyEndTime(t *testing.T) {
	startTime := time.Now()
	_, err := MakeTimeInterval(&startTime, nil)
	assert.EqualError(t, err, EmptyEndTimeError)
}

func TestTimeInterval_ToString(t *testing.T) {
	assert.Equal(t, "2014-05-03T00:00:00Z - 2014-05-04T00:00:00Z", i1.ToString(time.RFC3339))
}

func TestMakeInterval(t *testing.T) {
	startTime, endTime := time.Now(), time.Now()
	interval, err := MakeTimeInterval(&startTime, &endTime)
	assert.NoError(t, err)
	assert.Equal(t, &startTime, interval.Start)
	assert.Equal(t, &endTime, interval.End)
}

func TestTimeInterval_EqualWithNilParameter(t *testing.T) {
	assert.False(t, i1.Equal(nil))
}

func TestTimeInterval_Equal(t *testing.T) {
	assert.False(t, i1.Equal(i2))
	assert.True(t, i1.Equal(i1))
}

func TestTimeInterval_MeetsWithNilParameter(t *testing.T) {
	assert.False(t, i1.Meets(nil, TwentyFourHours))
}

func TestTimeInterval_Meets(t *testing.T) {
	constraint := time.Duration(TwentyFourHours)
	assert.True(t, i2.Meets(i1, constraint))
	assert.False(t, i1.Meets(i1, constraint))
}

func TestTimeInterval_MetByWithNilParameter(t *testing.T) {
	assert.False(t, i1.MetBy(nil, TwentyFourHours))
}

func TestTimeInterval_MetBy(t *testing.T) {
	constraint := time.Duration(TwentyFourHours)
	assert.True(t, i1.MetBy(i2, constraint))
	assert.False(t, i1.MetBy(i1, constraint))
}

func TestTimeInterval_PrecedesWithNilParameter(t *testing.T) {
	assert.False(t, i1.Precedes(nil, TwentyFourHours))
}

func TestTimeInterval_Precedes(t *testing.T) {
	// interval 2
	vPrecedeS := "2014-05-06 00:00 UTC"
	vPrecedeE := "2014-05-10 00:00 UTC"
	tPrecedeS, _ := time.Parse(timeFormat, vPrecedeS)
	tPrecedeE, _ := time.Parse(timeFormat, vPrecedeE)
	iPrecede, _ := MakeTimeInterval(&tPrecedeS, &tPrecedeE)
	constraint := time.Duration(TwentyFourHours)
	assert.True(t, i1.Precedes(iPrecede, constraint))
	assert.False(t, i1.Precedes(i2, constraint))
}