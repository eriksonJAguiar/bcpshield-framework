package diffpriv

import (
	"fmt"
	"testing"
)

func TestNumericDifpriv(t *testing.T) {

	originalData := []float64{0.5, 0.9, 1.3, 0.1, 0.3, 0.6}
	//First step is convert any data type to Matrix
	convertedData := TransforFloatData(originalData)
	//define amount data recovered
	amountQuery := 4

	//Send a query in dataset
	DataFromQuery := Query(convertedData, amountQuery)
	//Apply diiferential privacy function
	epsilon := 1.0
	noiseData := DiffPriv(DataFromQuery, amountQuery, convertedData, epsilon)

	fmt.Println("Dataset noise ", noiseData)
}

func TestSymbolicDifpriv(t *testing.T) {
	originalData := []string{"So close no matter how far", "It couldn't be much more from the heart", "Forever trusting who we are", "And nothing else matters"}
	//First step is convert any data type to Matrix
	convertedData := TransforSymbolicData(originalData)
	//define amount data recovered
	amountQuery := 0
	//Send a query in dataset
	DataFromQuery := Query(convertedData, amountQuery)
	//Apply diiferential privacy function
	epsilon := 1.0
	noiseData := DiffPriv(DataFromQuery, amountQuery, convertedData, epsilon)
	fmt.Println(convertedData)
	fmt.Println("---------------------------")
	fmt.Println("Dataset noise ", noiseData)
}
