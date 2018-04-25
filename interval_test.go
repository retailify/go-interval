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

func TestMakeTimeIntervalWithEmptyStartTime(t *testing.T) {
	_, err := MakeTimeInterval(nil, nil)
	assert.EqualError(t, err, EmptyStartTimeError)
}

func TestMakeTimeIntervalWithEmptyEndTime(t *testing.T) {
	startTime := time.Now()
	_, err := MakeTimeInterval(&startTime, nil)
	assert.EqualError(t, err, EmptyEndTimeError)
}

func TestMakeInterval(t *testing.T) {
	startTime, endTime := time.Now(), time.Now()
	interval, err := MakeTimeInterval(&startTime, &endTime)
	assert.NoError(t, err)
	assert.Equal(t, &startTime, interval.Start)
	assert.Equal(t, &endTime, interval.End)
}

func TestTimeInterval_Equal(t *testing.T) {
	startTime, endTime := time.Now(), time.Now()
	i1, _ := MakeTimeInterval(&startTime, &endTime)
	i2, _ := MakeTimeInterval(&startTime, &endTime)
	assert.True(t, true, i1.Equal(i2))
}