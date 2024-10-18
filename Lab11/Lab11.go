package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func runge(f func(float64, float64, float64) (float64, float64), u0 float64, u1 float64, h float64, xStart float64, xEnd float64) ([]float64, []float64) {
	x := xStart
	u := u0
	uValues := []float64{u}
	u1Values := []float64{u1} 

	for x < xEnd {
		k1u, k1u1 := f(x, u, u1)
		k2u, k2u1 := f(x+h/2, u+k1u*h/2, u1+k1u1*h/2)
		k3u, k3u1 := f(x+h, u+k2u*h, u1+k2u1*h)

		u += (h / 6) * (k1u + 4*k2u + k3u)
		u1 += (h / 6) * (k1u1 + 4*k2u1 + k3u1)

		x += h
		uValues = append(uValues, u)
		u1Values = append(u1Values, u1)
	}

	return uValues, u1Values
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

func f1(x float64, u float64, u1 float64) (float64, float64) {
	return u1, x * x
}


func f2(x float64, u float64, u1 float64) (float64, float64) {
	return u1, 4 * u
}

func f3(x float64, u float64, u1 float64) (float64, float64) {
	return u1, 2*u1 - u
}

func exactSolution1(x float64) float64 {
	return ((math.Pow(x, 4) / 12) + 1)
}

func exactSolution2(x float64) float64 {
	return math.Exp(x * 2)
}

func exactSolution3(x float64) float64 {
	return (math.Exp(x) - x*math.Exp(x))
}

func creategrafig(name string, xStart float64, xEnd float64, h float64, u1 []float64, u2 []float64) {

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

	line1, _ := plotter.NewLine(pts1)
	line1.Color = plotutil.Color(0)
	line1.LineStyle.Width = 2

	line2, _ := plotter.NewLine(pts2)
	line2.Color = plotutil.Color(2)
	line2.LineStyle.Width = 2

	p.Legend.Add("Рунге Кутты", line1)
	p.Legend.Add("Точный", line2)

	p.Legend.Top = true
	p.Legend.Left = true

	p.Add(line1, line2)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, name+".png"); err != nil {
		panic(err)
	}

}

func printTable(xStart float64, xEnd float64, h float64, u1 []float64, u2 []float64) {
	fmt.Printf("%-10s %-15s %-15s %-15s\n", "", "", "", "",)
	fmt.Printf("%-10s %-15s %-15s %-15s\n", "h", "Рунге Кутт", "Точный", "Разница")
	fmt.Printf("%-10s %-15s %-15s %-15s\n", "", "", "", "")
	for i := 0; i < len(u2); i++ {
		rungeValue := u1[i]
		accurateValue := u2[i]

		diff := math.Abs(accurateValue - rungeValue)

		fmt.Printf("%-10.2f %-15.6f %-15.6f %-15.6f \n", xStart, rungeValue, accurateValue, diff)
		xStart += h
	}
}

func main() {
	hValues := []float64{0.1, 0.01}

	fmt.Println("Уравнение 1:")
	for _, h := range hValues {

		xStart := 0.0
		xEnd := 1.0
		runge, _ := runge(f1, 1.0, 0.0, h, xStart, xEnd)
		real := accurate(exactSolution1, h, xStart, xEnd)

		printTable(xStart, xEnd, h, runge, real)
		creategrafig("График 1", xStart, xEnd, h, runge, real)

	}

	fmt.Println("Уравнение 2:")
	for _, h := range hValues {

		xStart := 0.0
		xEnd := 1.0
		runge, _ := runge(f2, 1.0, 2.0, h, xStart, xEnd)
		real := accurate(exactSolution2, h, xStart, xEnd)

		printTable(xStart, xEnd, h, runge, real)
		creategrafig("График 2", xStart, xEnd, h, runge, real)

	}

	fmt.Println("Уравнение 3:")

	for _, h := range hValues {

		xStart := 0.0
		xEnd := 1.0
		runge, _ := runge(f3, 1.0, 0.0, h, xStart, xEnd)
		real := accurate(exactSolution3, h, xStart, xEnd)

		printTable(xStart, xEnd, h, runge, real)
		creategrafig("График 3", xStart, xEnd, h, runge, real)

	}

}
