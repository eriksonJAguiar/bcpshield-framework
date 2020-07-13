package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestVariance(t *testing.T) {
	_, err := stats.Variance([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestPopulationVariance(t *testing.T) {
	e, err := stats.PopulationVariance([]float64{})
	if !math.IsNaN(e) {
		t.Errorf("%.1f != %.1f", e, math.NaN())
	}
	if err != stats.EmptyInputErr {
		t.Errorf("%v != %v", err, stats.EmptyInputErr)
	}

	pv, _ := stats.PopulationVariance([]float64{1, 2, 3})
	a, err := stats.Round(pv, 1)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if a != 0.7 {
		t.Errorf("%.1f != %.1f", a, 0.7)
	}
}

func TestSampleVariance(t *testing.T) {
	m, err := stats.SampleVariance([]float64{})
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
	if err != stats.EmptyInputErr {
		t.Errorf("%v != %v", err, stats.EmptyInputErr)
	}
	m, _ = stats.SampleVariance([]float64{1, 2, 3})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}
}

func TestCovariance(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 2, 3, 5, 6}
	s4 := []float64{}

	_, err := stats.Covariance(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := stats.Covariance(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 3.2499999999999996 {
		t.Errorf("Covariance %v != %v", a, 3.2499999999999996)
	}

	_, err = stats.Covariance(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestCovariancePopulation(t *testing.T) {
	s1 := []float64{1, 2, 3.5, 3.7, 8, 12}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{0.5, 1, 2.1, 3.4, 3.4, 4}
	s4 := []float64{}

	_, err := stats.CovariancePopulation(s1, s2)
	if err == nil {
		t.Errorf("Mismatched slice lengths should have returned an error")
	}

	a, err := stats.CovariancePopulation(s1, s3)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	if a != 4.191666666666666 {
		t.Errorf("CovariancePopulation %v != %v", a, 4.191666666666666)
	}

	_, err = stats.CovariancePopulation(s1, s4)
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
