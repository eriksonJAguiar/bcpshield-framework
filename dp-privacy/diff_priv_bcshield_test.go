package test

import (
	"testing"
)

func TestNumericDifpriv(t *testing.T) {
}

func TestSymbolicDifpriv(t *testing.T) {
	originalData := []string{"So close no matter how far", "It couldn't be much more from the heart", "Forever trusting who we are", "And nothing else matters"}
	//First step is convert any data type to Matrix
	convertedData := dp.transforSymbolicData(originalData)

	//Send a query in dataset
	DataFromQuery := dp.query(convertedData, 4)

	//Apply diiferential privacy function
	epsilon := 1.0

	dp.diffPriv(DataFromQuery, convertedData, epsilon)
}
