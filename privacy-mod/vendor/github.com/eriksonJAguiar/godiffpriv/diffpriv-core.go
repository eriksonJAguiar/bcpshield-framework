package godiffpriv

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/montanaflynn/stats"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// Internal representation for numeric datasets
type quantitative struct {
	data []float64
}

// Internal representation for symbolic datasets
type qualitative struct {
	data []string
}

// Interface to represents private data intenal, it was implemented for query and sesitivity methods
type private interface {
	query()
	sensitivity()
}

//Privatevalue is a interface to represents a private values to implements the object factory
type Privatevalue interface {
	ApplyPrivacy(float64) ([]byte, error)
}

// Object will be build to represents for symbolic datasets
type privatequali struct {
	data []string
}

// Object will be build to represents for numeric datasets
type privatequant struct {
	data []float64
}

// LapMechanism define noise mechanism to used on differential privacy, to calculate random noise
// Params: mi (float64): distribution mean; scale (float64): it is a standard deviation;
// times (int: amount random values will be gererate
// Returns: float65 array with values of the map
func lapMechanism(mi float64, scale float64, times int) ([]float64, error) {
	var dLap distuv.Laplace
	dLap.Mu = 0
	dLap.Scale = scale
	dLap.Src = rand.NewSource(uint64(time.Now().UTC().UnixNano()))
	var probs []float64
	for i := 0; i < times; i++ {
		probs = append(probs, dLap.Rand())
	}

	return probs, nil
}

func (q *quantitative) query() (float64, error) {
	return stats.Mean(q.data)
}

func (q *qualitative) query() (map[string]int, error) {
	hist := make(map[string]int)
	for _, item := range q.data {
		hist[item]++
	}

	return hist, nil
}

// mapToSliceInt is a fucntion to convert map values to slice
// Params: data (maps[string]int): dataset will be convert
// Returns: float65 array with values of the map
func mapToSliceInt(data map[string]int) []float64 {
	var values []float64
	for _, value := range data {
		values = append(values, float64(value))
	}

	return values
}

// sensitivity is a method to calculate sensitivity on dataset with numeric data
// Params: None
// Returns: float64 to represents sensitivity value and the error if there exists
func (q *quantitative) sensitivity() (float64, error) {

	var val float64
	var max float64

	for i := 0; i < len(q.data); i++ {
		d1Slice := make([]float64, len(q.data))
		d2Slice := make([]float64, len(q.data))
		var d1 quantitative
		var d2 quantitative
		copy(d1Slice, q.data)
		copy(d2Slice, q.data)
		d1.data = append(d1Slice[:i], d1Slice[i+1:]...)
		if (i + 1) > len(d2Slice)-1 {
			d2.data = d2Slice[1:]
		} else {
			d2.data = append(d2Slice[:i+1], d2Slice[(i+1)+1:]...)
		}
		q1, err := d1.query()
		q2, err := d2.query()

		if err != nil {
			return 0.0, err
		}

		val = math.Abs(q1 - q2)

		max = math.Max(max, val)

	}

	return max, nil
}

// sensitivity is a method to calculate sensitivity on dataset with symbolic data
// Params: None
// Returns: float64 to represents sensitivity value and the error if there exists
func (q *qualitative) sensitivity() (float64, error) {

	var val float64
	var max float64

	for i := 0; i < len(q.data); i++ {
		d1Slice := make([]string, len(q.data))
		d2Slice := make([]string, len(q.data))
		var d1 qualitative
		var d2 qualitative
		copy(d1Slice, q.data)
		copy(d2Slice, q.data)
		d1.data = append(d1Slice[:i], d1Slice[i+1:]...)
		if (i + 1) > len(d2Slice)-1 {
			d2.data = d2Slice[1:]
		} else {
			d2.data = append(d2Slice[:i+1], d2Slice[(i+1)+1:]...)
		}
		q1, _ := d1.query()
		q2, _ := d2.query()

		arrayQ1 := mapToSliceInt(q1)
		arrayQ2 := mapToSliceInt(q2)

		size := 0

		if len(arrayQ1) > len(arrayQ2) {
			size = len(arrayQ2)
		} else if len(arrayQ2) > len(arrayQ1) {
			size = len(arrayQ1)
		}

		for j := 0; j < size; j++ {
			result := math.Abs(arrayQ1[j] - arrayQ2[j])
			val = math.Max(val, result)
		}

		max = math.Max(max, val)

	}

	return max, nil
}

// ApplyPrivacy is a method to apply privacy on numeric data
// Params: epsilon (float64): noise level
// Returns: a byte array that would be converted to map[string]float64
func (priv *privatequant) ApplyPrivacy(epsilon float64) ([]byte, error) {

	q := new(quantitative)
	q.data = priv.data

	s, _ := q.sensitivity()
	b := s / epsilon
	noise, _ := lapMechanism(0, b, 1)
	data, _ := q.query()

	privData := data + noise[0]

	privBytes, err := json.Marshal(map[string]float64{"data": privData})

	if err != nil {
		return nil, err
	}

	return privBytes, nil
}

// ApplyPrivacy is a method to apply privacy on symbolic data
// Params: epsilon (float64): noise level
// Returns: a byte array that would be converted to map[string]float64
func (priv *privatequali) ApplyPrivacy(epsilon float64) ([]byte, error) {
	q := new(qualitative)
	q.data = priv.data

	s, _ := q.sensitivity()
	b := s / epsilon
	noise, err := lapMechanism(0, b, 1)
	if err != nil {
		return nil, err
	}
	data, err := q.query()
	if err != nil {
		return nil, err
	}

	privData := make(map[string]float64)

	i := 1
	for _, val := range data {
		key := strconv.Itoa(i)
		privData[key] = float64(val) + noise[0]
		i++
	}

	privBytes, err := json.Marshal(privData)

	if err != nil {
		return nil, err
	}

	return privBytes, nil
}

// PrivateDataFactory is a factory method to generate private objects
// Params: dataset: a string to describes object type
// Returns: object to numeric or symbolic data
func PrivateDataFactory(dataset interface{}) Privatevalue {
	val := reflect.ValueOf(dataset)
	t := val.Index(0)
	switch t.Kind() {
	case reflect.Float64:
		//dt := make([]float64, val.Len())
		var dt []float64
		for i := 0; i < val.Len(); i++ {
			dt = append(dt, float64(val.Index(i).Float()))
		}
		quant := new(privatequant)
		quant.data = dt
		return quant
	case reflect.String:
		var dt []string
		for i := 0; i < val.Len(); i++ {
			dt = append(dt, string(val.Index(i).String()))
		}
		quali := new(privatequali)
		quali.data = dt
		return quali
	default:
		return nil
	}
}
