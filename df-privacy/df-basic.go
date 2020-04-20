package diffPriv

import (
	"math"
	"time"

	"github.com/montanaflynn/stats"
	"golang.org/x/exp/rand"
)

// Represents a type from differential privacy attributes
type DiffPrivVal struct {
	MinProb  float64
	indexMin int
	Prob     []float64
	Noise    []float64
}

/*
* -----------------------------------------------------------
* Imported from github.com/gonum/stat/distuv - Gonum
* Due to troubles in library
* -----------------------------------------------------------
 */
//Represents a type for Laplacian distribution attributes
type Laplace struct {
	Mu    float64
	Scale float64
	Src   rand.Source
}

/*
* -----------------------------------------------------------
* -----------------------------------------------------------
 */

/*
* -----------------------------------------------------------
* Imported from github.com/gonum/stat/distuv - Gonum
* Due to troubles in library
* -----------------------------------------------------------
 */

//Imported from github.com/gonum/stat/distuv - Gonum
func (l Laplace) LogProb(x float64) float64 {
	return -math.Ln2 - math.Log(l.Scale) - math.Abs(x-l.Mu)/l.Scale
}

// Imported from github.com/gonum/stat/distuv - Gonum
func (l Laplace) Prob(x float64) float64 {
	return math.Exp(l.LogProb(x))
}

// Imported from github.com/gonum/stat/distuv - Gonum
func (l Laplace) Rand() float64 {
	var rnd float64
	if l.Src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(l.Src).Float64()
	}
	u := rnd - 0.5
	if u < 0 {
		return l.Mu + l.Scale*math.Log(1+2*u)
	}
	return l.Mu - l.Scale*math.Log(1-2*u)
}

/*
* -----------------------------------------------------------
* -----------------------------------------------------------
 */

// Findo the min  and max values to array float64
func findMaxMin(vals []float64) (float64, float64) {

	var max float64
	var min float64

	for _, val := range vals {
		max = math.Max(max, val)
		min = math.Min(min, val)
	}

	return min, max
}

// Remove a element by index of the array float64
func removeElement(index int, array []float64) []float64 {

	var part1 []float64
	var part2 []float64

	part1 = array[:index]
	part2 = array[index+1:]

	sizeNew := len(part1) + len(part2)
	newArray := make([]float64, 0, sizeNew)

	for _, elm := range part1 {
		newArray = append(newArray, elm)
	}

	for _, elm := range part2 {
		newArray = append(newArray, elm)
	}

	return newArray
}

// Example for average query from database used
func query(data []float64) float64 {
	mu, _ := stats.Mean(data)

	return mu
}

// Calculte the sensitiy of the query
func sensitivity(database []float64) float64 {

	var vals []float64

	for i := 0; i < len(database)-1; i += 2 {
		d1 := removeElement(i, database)
		d2 := removeElement(i+1, database)
		q1 := query(d1)
		q2 := query(d2)
		vals = append(vals, (math.Abs(q1 - q2)))
	}
	_, max := findMaxMin(vals)

	return max
}

// Calulate a b value in the differential privacy
func blaplace(sens float64, epsilon float64) float64 {
	return (sens / epsilon)
}

// Calcule a Laplacian noise and generate a random distribuion
func dflaplace(database []float64, epsilon float64) *DiffPrivVal {

	//var sample []float64

	s := sensitivity(database)
	b := blaplace(s, epsilon)

	lap := Laplace{Mu: 0, Scale: b}
	df := DiffPrivVal{MinProb: 0, indexMin: -1}
	df.Prob = make([]float64, 0, len(database))
	df.Noise = make([]float64, 0, len(database))

	df.Noise = append(df.Noise, lap.Rand())
	df.Prob = append(df.Prob, lap.Prob(df.Noise[0]))
	df.MinProb = df.Prob[0]
	df.indexMin = 0

	for i := 1; i < len(database); i++ {
		df.Noise = append(df.Noise, lap.Rand())
		df.Prob = append(df.Prob, lap.Prob(df.Noise[i]))
		if df.Prob[i] < df.MinProb {
			df.MinProb = df.Prob[i]
			df.indexMin = i
		}
	}

	return &df

	// for _, d := range database {
	// 	l := ((1 / (2 * b)) * (math.Exp(-((d) / b))))
	// 	dist = append(dist, l)
	// }
}

// Add noise on original query
func addNoise(query float64, df DiffPrivVal) float64 {

	var qnoise float64

	qnoise = query + df.Noise[df.indexMin]

	return qnoise
}

// Principal function for calculus of the differential pri
func diffPriv(query float64, data []float64, epsilon float64) float64 {
	var dfNoise DiffPrivVal
	dfNoise = (*dflaplace(data, epsilon))
	noiseQuery := addNoise(query, dfNoise)

	return noiseQuery

}

func randomGenerate(min float64, max float64, amount int) []float64 {
	rand.Seed(uint64(time.Now().UnixNano()))
	var rdnum []float64

	for i := 0; i < amount; i++ {
		rdnum = append(rdnum, (min + (rand.Float64() * (max - min))))
	}

	return rdnum
}
