package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/montanaflynn/stats"
)

func findMaxMin(vals []float64) (float64, float64) {

	var max float64
	var min float64

	for _, val := range vals {
		max = math.Max(max, val)
		min = math.Min(min, val)
	}

	return min, max
}

func sensitivity(query float64, database []float64) float64 {

	var vals []float64

	for _, q := range database {
		vals = append(vals, (math.Abs(query - q)))
	}
	_, max := findMaxMin(vals)

	return max
}

func blaplace(sens float64, epsilon float64) float64 {
	return (sens / epsilon)
}

func dflaplace(query float64, database []float64, epsilon float64) []float64 {

	var dist []float64

	s := sensitivity(query, database)
	b := blaplace(s, epsilon)
	mi, _ := stats.Mean(database)

	for _, d := range database {
		l := ((1 / (2 * b)) * (math.Exp(-((d - mi) / b))))
		dist = append(dist, l)
	}

	return dist
}

func randomGenerate(min float64, max float64, amount int) []float64 {
	rand.Seed(time.Now().UnixNano())
	var rdnum []float64

	for i := 0; i < amount; i++ {
		rdnum = append(rdnum, (min + (rand.Float64() * (max - min))))
	}

	return rdnum
}

func addNoise(query float64, noise []float64) []float64 {
	var qnoise []float64

	for _, n := range noise {
		qnoise = append(qnoise, (query + n))
	}

	return qnoise
}

func diffPriv(query float64, data []float64, epsilon float64, distribution string) []float64 {
	switch distribution {
	case "laplacian":
		noise := dflaplace(query, data, epsilon)
		noiseQuery := addNoise(query, noise)
		return noiseQuery
	case "standard":
		break
	case "exponential":
		break
	default:
		log.Println("[ERROR] Distribution not found!")
	}

	return nil
}

func main() {
	rdnum := randomGenerate(0, 1, 100)
	pdfNoise := diffPriv(0.6, rdnum, 0.1, "laplacian")
	fmt.Println("A funcao de ruido eh: ")
	fmt.Println(pdfNoise)
}
