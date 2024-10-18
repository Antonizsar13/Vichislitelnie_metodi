package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func adamsBashforth3(f func(float64, float64) float64, u0, u1, u2 float64, h float64, xStart float64, xEnd float64) []float64 {
	uValues := []float64{u0, u1, u2}
	x := xStart + 2*h

	for x < xEnd {
		uNext := uValues[len(uValues)-1] + (h/12)*(23*f(x-h, uValues[len(uValues)-1])-16*f(x-2*h, uValues[len(uValues)-2])+5*f(x-3*h, uValues[len(uValues)-3]))
		uValues = append(uValues, uNext)
		x += h
	}

	return uValues
}

func adamsMoulton3(f func(float64, float64) float64, u0, u1, u2 float64, h float64, xStart float64, xEnd float64) []float64 {
	uValues := []float64{u0, u1, u2}
	x := xStart + 2*h

	for x < xEnd {
		uPrev := uValues[len(uValues)-1]
		uNext := uPrev + (h/24)*(9*f(x, uPrev)+19*f(x-h, uValues[len(uValues)-2])-5*f(x-2*h, uValues[len(uValues)-3]) )
		uValues = append(uValues, uNext)
		x += h
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
	return u + math.Pow(x, 2)
}
func f2(x float64, u float64) float64 {
	return x * u
}
func f3(x float64, u float64) float64 {
	return math.Exp(x) + u
}

func exactSolution1(x float64) float64 {
	return -math.Pow(x,2) -2*x+ 3* math.Exp(x) -2
}

func exactSolution2(x float64) float64 {
	return 2 * math.Exp((math.Pow(x, 2) / 2))
}

func exactSolution3(x float64) float64 {
	return (math.Exp(x) + x*math.Exp(x))
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

	p.Legend.Add("Явный трёхшаговый", line1)
	p.Legend.Add("Неявный трёхшаговый", line2)
	p.Legend.Add("Точный", line3)

	p.Legend.Top = true
	p.Legend.Left = true

	if err := p.Save(8*vg.Inch, 8*vg.Inch, name+".png"); err != nil {
		panic(err)
	}
}

func printTable(xStart float64, xEnd float64, h float64, u1 []float64, u2 []float64, u3 []float64) {
	fmt.Printf("%-10s %-15s %-15s %-15s %-15s %-15s\n", "", "", "", "", "", "")
	fmt.Printf("%-10s %-15s %-15s %-15s %-15s %-15s\n", "h", "Явный", "Неявный", "Точный", "Разница Явного", "Разница Неявного")
	fmt.Printf("%-10s %-15s %-15s %-15s %-15s %-15s\n", "", "", "", "", "", "")
	for i := 0; i < len(u3); i++ {
		eulerValue := u1[i]
		symmetricValue := u2[i]
		accurateValue := u3[i]

		eulerDiff := accurateValue - eulerValue
		symmetricDiff := accurateValue - symmetricValue

		fmt.Printf("%-10.2f %-15.2f %-15.2f %-15.2f %-15.2f %-15.2f\n", xStart, eulerValue, symmetricValue, accurateValue, eulerDiff, symmetricDiff)
		xStart += h
	}
}

func main() {
	hValues := []float64{0.1, 0.01}

	fmt.Println("Уравнение 1:")
	for _, h := range hValues {
		uValues1 := accurate(exactSolution1, h, 0, 1)
		uValuesB := adamsBashforth3(f1, uValues1[0], uValues1[1], uValues1[2], h, 0, 1)
		uValuesA := adamsMoulton3(f1, uValues1[0], uValues1[1], uValues1[2], h, 0, 1)

		printTable(0, 2, h, uValuesB, uValuesA, uValues1)
		creategrafig("График 1", 0, 1, h, uValuesB, uValuesA, uValues1)
	}

	fmt.Println("Уравнение 2:")
	for _, h := range hValues {
		uValues1 := accurate(exactSolution2, h, 0, 1)
		uValuesB := adamsBashforth3(f2, uValues1[0], uValues1[1], uValues1[2], h, 0, 2)
		uValuesA := adamsMoulton3(f2, uValues1[0], uValues1[1], uValues1[2], h, 0, 2)

		printTable(0, 1, h, uValuesB, uValuesA, uValues1)
		creategrafig("График 2", 0, 2, h, uValuesB, uValuesA, uValues1)
	}

	fmt.Println("Уравнение 3:")
	for _, h := range hValues {
		uValues1 := accurate(exactSolution3, h, 0, 1)
		uValuesB := adamsBashforth3(f3, uValues1[0], uValues1[1], uValues1[2], h, 0, 1)
		uValuesA := adamsMoulton3(f3, uValues1[0], uValues1[1], uValues1[2], h, 0, 1)

		printTable(0, 1, h, uValuesB, uValuesA, uValues1)
		creategrafig("График 3", 0, 1, h, uValuesB, uValuesA, uValues1)
	}
}
