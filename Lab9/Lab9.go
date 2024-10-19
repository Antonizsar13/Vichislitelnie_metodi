package main

import (
	"fmt"
	"math"
	"math/rand"
	// "time"
)

type Interval struct {
	start float64
	end   float64
}

func f(x float64, y float64, z float64) float64 {
	return x
	// return x * math.Pow(y,2)
	// return math.Sin(y) * x
	// return x * math.Pow(z, 3)
	// return math.Pow(x,3)
	// return x* math.Exp(y)
}

func main() {

	Ns := []float64{math.Pow(10, 4),
		math.Pow(10, 6),
		math.Pow(10, 8)}
	answerds := make([]float64, len(Ns))
	intervals := []Interval{{1, 2}, {2, 4}, {1, 5}}
	multiplication := 1.0
	for _, interval := range intervals {
		multiplication *= (interval.end - interval.start)
	}

	for i, N := range Ns {
		answerds[i] = multiplication / N
		sum := 0.0
		for j := 0; j < int(N); j++ {
			// t := rand.Float64()
			xyz := make([]float64, len(intervals))
			for m, interval := range intervals {
				t := rand.Float64()
				xyz[len(xyz)-1-m] = interval.start + (interval.end-interval.start)*t
			}
			sum += f(xyz[0], xyz[1], xyz[2])
		}
		answerds[i] *= sum
		if i == 0 {
			fmt.Printf("Приближенное значение интеграла при n=%.1f: %f\n", N, answerds[i])
		} else {
			fmt.Printf("Приближенное значение интеграла при n=%.1f: %f. Разница=%f\n", N, answerds[i], math.Abs(answerds[i]-answerds[i-1]))
		}

	}

}
