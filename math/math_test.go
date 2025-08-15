package math_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSecondHandAt25Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 25, 0, time.UTC)

	want := Point{X: 150 + 45, Y: 150 + 78}
	got := CalculHand(tm, "seconds")

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestCalcTime(t *testing.T) {
	tests := []struct {
		tm     time.Time
		Points []Point
	}{
		{time.Date(1337, time.January, 1, 15, 30, 25, 0, time.UTC), []Point{
			{X: 200, Y: 150},
			{X: 150, Y: 150 + 80},
			{X: 150 + 45, Y: 150 + 78},
		}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.tm), func(t *testing.T) {
			want := test.Points
			got := CalcTime(test.tm)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %f, wanted %f", got, want)
			}
		})
	}
}

func TestCalculAngle(t *testing.T) {
	tests := []struct {
		Aiguille float64
		Angle    float64
	}{
		{25, 300},
		{30, 90 * 3},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%.0f", test.Aiguille), func(t *testing.T) {
			want := test.Angle
			got := calculAngle(test.Aiguille)
			if got != want {
				t.Errorf("Got %f, wanted %f", got, want)
			}
		})
	}
}

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	b := bytes.Buffer{}

	SVGWriter(&b, tm)

	svg := SVG{}

	xml.Unmarshal(b.Bytes(), &svg)

	want := Line{150, 150, 150, 60}

	for _, line := range svg.Line {
		if line == want {
			return
		}
	}

	t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			fmt.Println(b.String())
			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}
