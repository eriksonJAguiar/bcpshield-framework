package diffpriv

import (
	"testing"
)

func TestNumericDifpriv(t *testing.T) {

	originalData := []float64{0.5, 0.9, 1.3, 0.1, 0.3, 0.6}
	//First step is convert any data type to Matrix
	convertedData := transforFloatData(originalData)

	//Send a query in dataset
	DataFromQuery := query(convertedData, 4)

	//Apply diiferential privacy function
	epsilon := 1.0

	diffPriv(DataFromQuery, convertedData, epsilon)
}

func TestSymbolicDifpriv(t *testing.T) {
	originalData := []string{"So close no matter how far", "It couldn't be much more from the heart", "Forever trusting who we are", "And nothing else matters"}
	//First step is convert any data type to Matrix
	convertedData := transforSymbolicData(originalData)

	//Send a query in dataset
	DataFromQuery := query(convertedData, 4)

	//Apply diiferential privacy function
	epsilon := 1.0

	diffPriv(DataFromQuery, convertedData, epsilon)
}
