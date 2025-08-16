package math_test

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

type Point struct {
	X float64
	Y float64
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func CalcTime(tm time.Time) []Point {
	seconds := CalculHand(tm, "seconds")
	minutes := CalculHand(tm, "minutes")
	hours := CalculHand(tm, "hours")
	return []Point{
		{hours.X, hours.Y},
		{minutes.X, minutes.Y},
		{seconds.X, seconds.Y}}
}

func CalculHand(tm time.Time, hand string) Point {
	var h, l int

	switch hand {
	case "seconds":
		{
			h = tm.Second()
			l = 90
		}
	case "minutes":
		{
			h = tm.Minute()
			l = 80
		}
	case "hours":
		{
			h = tm.Hour() * 5
			l = 50
		}
	}
	x := 150 + (math.Cos((math.Pi/180)*calculAngle(float64(h))) * float64(l))
	y := 150 - (math.Sin((math.Pi/180)*calculAngle(float64(h))) * float64(l))
	return Point{math.Round(x), math.Round(y)}
}

func calculAngle(val float64) float64 {
	angleMidiVal := (val * 360) / 60
	angle090Val := angleMidiVal - 90
	return 360 - angle090Val
}

func SVGWriter(w io.Writer, tm time.Time) {
	points := CalcTime(tm)
	sb := strings.Builder{}
	sb.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">

  <!-- bezel -->
  <circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>
	`)
	sb.WriteString("<!-- hour hand -->\n")
	hourHand := fmt.Sprintf(`  <line x1="150" y1="150" x2="%.f" y2="%.f"
        style="fill:none;stroke:#000;stroke-width:5px;"/>`, points[0].X, points[0].Y)
	sb.WriteString(hourHand + "\n")

	sb.WriteString("  <!-- minute hand -->")
	minuteHand := fmt.Sprintf(`<line x1="150" y1="150" x2="%.f" y2="%.f"
        style="fill:none;stroke:#000;stroke-width:5px;"/>`, points[1].X, points[1].Y)
	sb.WriteString(minuteHand + "\n")

	sb.WriteString("  <!-- second hand -->")
	secondHand := fmt.Sprintf(`<line x1="150" y1="150" x2="%.f" y2="%.f"
        style="fill:none;stroke:#f00;stroke-width:3px;"/>`, points[2].X, points[2].Y)
	sb.WriteString(secondHand + "\n")
	sb.WriteString("</svg>")
	w.Write([]byte(sb.String()))
	fmt.Printf("seconds : %d \n Point {x : %f, y : %f} \n", tm.Second(), points[2].X, points[2].Y)
}
