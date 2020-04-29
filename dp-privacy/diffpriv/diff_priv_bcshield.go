package diffpriv

import (
	"encoding/json"
	"fmt"
	"math"

	distcos "github.com/gaspiman/cosine_similarity"
	tfidf "github.com/numbleroot/go-tfidf"
	"golang.org/x/exp/rand"
)

// Represents a type from differential privacy attributes
type diffPrivVal struct {
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
type laplace struct {
	Mu    float64
	Scale float64
	Src   rand.Source
}

/*
* -----------------------------------------------------------
* -----------------------------------------------------------
 */

// Represents a type for dabaset used
type Matrix struct {
	Data []float64
	Type string
}

/*
* -----------------------------------------------------------
* Imported from github.com/gonum/stat/distuv - Gonum
* Due to troubles in library
* -----------------------------------------------------------
 */

//Imported from github.com/gonum/stat/distuv - Gonum
func (l laplace) LogProb(x float64) float64 {
	return -math.Ln2 - math.Log(l.Scale) - math.Abs(x-l.Mu)/l.Scale
}

// Imported from github.com/gonum/stat/distuv - Gonum
func (l laplace) Prob(x float64) float64 {
	return math.Exp(l.LogProb(x))
}

// Imported from github.com/gonum/stat/distuv - Gonum
func (l laplace) Rand() float64 {
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

// Remove a element by index of the array float64 that represents Numeric elements
func removeNumericElement(index int, array []float64) []float64 {
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

// Remove a element by index of the array float64 that represents symbolic values
func removeElement(index int, array []Matrix) []Matrix {

	if array[0].Type == "numeric" {
		numericaData := removeNumericElement(index, array[0].Data)
		return []Matrix{{Data: numericaData}}
	}

	var part1 []Matrix
	var part2 []Matrix

	part1 = array[:index]
	part2 = array[index+1:]

	sizeNew := len(part1) + len(part2)
	newArray := make([]Matrix, 0, sizeNew)

	for _, elm := range part1 {
		newArray = append(newArray, elm)
	}

	for _, elm := range part2 {
		newArray = append(newArray, elm)
	}

	return newArray
}

// Removo elements from dataset
// func removeIndex(dataset []Matrix, index int) []Matrix {
// 	if dataset[0].Type == "numeric" {
// 		datasetValues := dataset[0].Data
// 		var datasetCropped []Matrix
// 		var matNewElement Matrix
// 		matNewElement.Data = append(datasetValues[:index], datasetValues[index+1:]...)
// 		matNewElement.Type = "numeric"

// 		datasetCropped = append(datasetCropped, matNewElement)

// 		return datasetCropped

// 	}

// 	return append(dataset[:index], dataset[index+1:]...)
// }

// Query amount from id record
func Query(data []Matrix, amount int) []Matrix {

	var matrixQuery []Matrix

	if data[0].Type == "string" {
		if amount >= len(data) {
			return data
		}

		matrixQuery = data[:amount]
		return matrixQuery
	}

	if amount >= len(data[0].Data) {
		return data
	}

	var auxMatrix Matrix

	auxMatrix.Data = data[0].Data[:amount]
	auxMatrix.Type = "numeric"

	return []Matrix{auxMatrix}
}

// Calculate the maximum difference for numeric values
func maxDifferenceForSymbolic(q1 []Matrix, q2 []Matrix) float64 {

	var max float64

	for j := 0; j < len(q1); j++ {
		distanceDatabases, _ := distcos.Cosine(q1[j].Data, q2[j].Data)
		max = math.Max(max, distanceDatabases)
	}

	return max
}

// Calculate the maximum difference for numeric values
func maxDifferenceForNumeric(q1 []Matrix, q2 []Matrix) float64 {
	var max float64

	for i := 0; i < len(q1[0].Data); i++ {
		val := math.Abs((q1[0].Data[i]) - (q2[0].Data[i]))
		max = math.Max(max, val)
	}

	return max
}

// Calculate the sensitivity for database
func sensitivity(database []Matrix, amountQuery int) float64 {

	var val float64
	var max float64

	for i := 0; i < len(database); i += 2 {
		d1 := removeElement(i, database)
		d2 := removeElement(i+1, database)
		q1 := Query(d1, amountQuery)
		q2 := Query(d2, amountQuery)

		if q1[0].Type == "string" {
			val = maxDifferenceForSymbolic(q1, q2)
		} else {
			val = maxDifferenceForNumeric(q1, q2)
		}

		max = math.Max(max, val)

	}

	return max
}

// Calulate b value in the differential privacy
func blaplace(sens float64, epsilon float64) float64 {
	return (sens / epsilon)
}

// Calculate a Laplacian noise and generate a random distribuion
func dflaplace(database []Matrix, amountQuery int, epsilon float64) *diffPrivVal {

	//var sample []float64

	s := sensitivity(database, amountQuery)
	b := blaplace(s, epsilon)

	lap := laplace{Mu: 0, Scale: b}
	df := diffPrivVal{MinProb: 0, indexMin: -1}
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
}

// Add Noise in numeric data
func addNoiseForNumericData(query []float64, noise float64) []float64 {
	var qnoise []float64

	for _, q := range query {
		qnoise = append(qnoise, q+noise)
	}

	return qnoise
}

// Add Noise in Symbolic data
func addNoiseForSymbolicData(query []Matrix, noise float64) []Matrix {
	var qnoise []Matrix

	for _, q := range query {
		var auxQnoise Matrix
		auxQnoise.Data = addNoiseForNumericData(q.Data, noise)
		auxQnoise.Type = "string"
		qnoise = append(qnoise, auxQnoise)
	}

	return qnoise
}

// Add noise on query selected through laplace distribution
func addNoise(query []Matrix, df diffPrivVal) []Matrix {

	var qnoise []Matrix

	if query[0].Type == "numeric" {
		dataValue := addNoiseForNumericData(query[0].Data, df.Noise[df.indexMin])
		qnoise = []Matrix{{Data: dataValue, Type: "numeric"}}
		return qnoise
	}

	qnoise = addNoiseForSymbolicData(query, df.Noise[df.indexMin])

	return qnoise
}

// Traform data from symbolic to Matrix
func TransforSymbolicData(dataset []string) []Matrix {
	var dtokens [][]string

	for _, dt := range dataset {
		dtokens = append(dtokens, tfidf.TokenizeDocument(dt))
	}

	idf := tfidf.InverseDocumentFrequencies(dtokens, tfidf.TermWeightingRaw)

	var datasetTfid []Matrix

	for _, terms := range dtokens {
		var docRes Matrix
		for _, t := range terms {
			freq := tfidf.TermFrequency(t, false, terms, tfidf.TermWeightingRaw)
			res := freq * idf[t]
			docRes.Data = append(docRes.Data, res)
			docRes.Type = "string"
		}
		datasetTfid = append(datasetTfid, docRes)
	}

	return datasetTfid
}

// Transform data from Int to Matrix
func TransforIntData(data []int) []Matrix {
	newData := make([]float64, 0, len(data))
	for _, d := range data {
		newData = append(newData, float64(d))
	}
	return []Matrix{{Data: newData, Type: "numeric"}}

}

// Transform data from Float to Matrix
func TransforFloatData(data []float64) []Matrix {
	return []Matrix{{Data: data, Type: "numeric"}}
}

// Main function to calculate the differential privacy
func DiffPriv(query []Matrix, amountQuery int, dataset []Matrix, epsilon float64) string {

	var dfNoise diffPrivVal

	dfNoise = (*dflaplace(dataset, amountQuery, epsilon))
	noiseQuery := addNoise(query, dfNoise)

	jsonQuey, err := json.Marshal(noiseQuery)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(jsonQuey)

}
