package main

import (
	"fmt"

	dp "github.com/icmc-wines/dp-privacy/diffpriv"
)

func main() {

	originalData := []float64{0.5, 0.9, 1.3, 0.1, 0.3, 0.6}
	//First step is convert any data type to Matrix
	convertedData := dp.TransforFloatData(originalData)
	//Send a query in dataset
	DataFromQuery := dp.Query(convertedData, 4)
	//Apply diiferential privacy function
	epsilon := 1.0
	noiseData := dp.DiffPriv(DataFromQuery, convertedData, epsilon)

	fmt.Println("Dataset noise ", noiseData)
}
