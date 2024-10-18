package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
    "gonum.org/v1/plot/plotutil"
)

func eulerMethod(f func(float64, float64) float64, u0 float64, h float64, xStart float64, xEnd float64) []float64 {
	x := xStart
	u := u0
	uValues := []float64{u}

	for x < xEnd {
		u += h * f(x, u)
		x += h
		uValues = append(uValues, u)
	}

	return uValues
}

func simmetrMethod(f func(float64, float64) float64, u0 float64, h float64, xStart float64, xEnd float64) []float64 {
	x := xStart
	u := u0
	uValues := []float64{u}

	for x < xEnd {
		uPrev := uValues[len(uValues)-1]
		uNow := uPrev + h*f(x-h, uPrev)
		u := uPrev + h/2*(f(x-h, uPrev)+f(x+h, uNow))
		x += h
		uValues = append(uValues, u)
	}

	return uValues
}

func accurate(exactFunc func(float64) float64, h float64, xStart float64, xEnd float64) []float64 {
	x := xStart
	accurateValues := []float64{}

	for x <= xEnd {
		accurateValues = append(accurateValues, exactFunc(x))
		x += h
	}

	return accurateValues
}

func f1(x float64, u float64) float64 {
	return x*x + 3*u
}
func f2(x float64, u float64) float64 {
	return 2*u + 4*x
}
func f3(x float64, u float64) float64 {
	return 2*u + math.Exp(x)
}

func exactSolution1(x float64) float64 {
	return (56.0/27.0)*math.Exp(3*x) - (1.0/3.0)*x*x - (2.0/9.0)*x - (2.0 / 27.0)
}

func exactSolution2(x float64) float64 {
	return -2*x + 2*math.Exp(2*x) - 1
}

func exactSolution3(x float64) float64 {
	return (math.Exp(x) * (math.Exp(x) - 1))
}

func creategrafig(name string, xStart float64, xEnd float64, h float64, u1 []float64, u2 []float64, u3 []float64) {

	p := plot.New()

	p.Title.Text = name
	p.X.Label.Text = "X"
	p.Y.Label.Text = "U"

	var x []float64
	for i := xStart; i <= xEnd+h; i += h {
		x = append(x, i)
	}

	pts1 := make(plotter.XYs, len(u1))

	for i := range pts1 {
		pts1[i].X = x[i]
		pts1[i].Y = u1[i]
	}

	pts2 := make(plotter.XYs, len(u2))

	for i := range pts2 {
		pts2[i].X = x[i]
		pts2[i].Y = u2[i]
	}

	pts3 := make(plotter.XYs, len(u3))

	for i := range pts3 {
		pts3[i].X = x[i]
		pts3[i].Y = u3[i]
	}

	line1, _ := plotter.NewLine(pts1)
	line1.Color = plotutil.Color(0) 
	line1.LineStyle.Width = 2

	line2, _ := plotter.NewLine(pts2)
	line2.Color = plotutil.Color(2) 
	line2.LineStyle.Width = 2

	line3, _ := plotter.NewLine(pts3)
	line3.Color = plotutil.Color(5)
    line3.LineStyle.Width = 2
    p.Add(line1, line2, line3)

	p.Legend.Add("Эйлер", line1)
	p.Legend.Add("Симметричный", line2)
	p.Legend.Add("Точный", line3)

	p.Legend.Top = true
	p.Legend.Left = true

    
	if err := p.Save(8*vg.Inch, 8*vg.Inch, name+".png"); err != nil {
		panic(err)
	}

}

func main() {
	hValues := []float64{0.1, 0.01}

	fmt.Println("Уравнение 1:")
	for _, h := range hValues {

		uValues := eulerMethod(f1, 2.0, h, 0, 2)
		uValues2 := simmetrMethod(f1, 2.0, h, 0, 2)
		uValues3 := accurate(exactSolution1, h, 0, 2)

		fmt.Printf("Эйлер h = %.2f: %v\n\n", h, uValues)
		fmt.Printf("Симметричный h = %.2f: %v\n\n", h, uValues2)
		fmt.Printf("Точный h = %.2f: %v\n\n", h, uValues3)

		creategrafig("График 1", 0, 2, h, uValues, uValues2, uValues3)

	}

	fmt.Println("Уравнение 2:")
	for _, h := range hValues {
		uValues := eulerMethod(f2, 1.0, h, 0, 1)
		uValues2 := simmetrMethod(f2, 1.0, h, 0, 1)
		uValues3 := accurate(exactSolution2, h, 0, 1)

		fmt.Printf("Эйлер h = %.2f: %v\n\n", h, uValues)
		fmt.Printf("Симметричный h = %.2f: %v\n\n", h, uValues2)
		fmt.Printf("Точный h = %.2f: %v\n\n", h, uValues3)

		creategrafig("График 2", 0, 1, h, uValues, uValues2, uValues3)
	}

	fmt.Println("Уравнение 3:")
	for _, h := range hValues {
		uValues := eulerMethod(f3, 0.0, h, 0, 1)
		uValues2 := simmetrMethod(f3, 0.0, h, 0, 1)
		uValues3 := accurate(exactSolution3, h, 0, 1)

		fmt.Printf("Эйлер h = %.2f: %v\n\n", h, uValues)
		fmt.Printf("Симметричный h = %.2f: %v\n\n", h, uValues2)
		fmt.Printf("Точный h = %.2f: %v\n\n", h, uValues3)

		creategrafig("График 3", 0, 1, h, uValues, uValues2, uValues3)
	}

}
