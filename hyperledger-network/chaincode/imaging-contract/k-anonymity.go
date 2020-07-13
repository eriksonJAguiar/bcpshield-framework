package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type limValues struct {
	limInf float64
	limSup float64
}

func KAnonymitySupression(column []string) []string {
	var newColumn []string
	for i := 0; i < len(column); i++ {
		newColumn = append(newColumn, "*")
	}

	return newColumn
}

func KAnonymitygeneralizationSymbolic(column []string) []string {
	var newColumn []string
	for _, col := range column {
		var token string

		re, err := regexp.Compile(`\d{4}-\d{2}-\d{2}`)
		if err != nil {
			panic("Error: " + err.Error())
			return nil
		}
		if re.MatchString(col) {
			tm := strings.Split(col, "-")
			token = fmt.Sprint(tm[0])
			newColumn = append(newColumn, token)
			continue
		}

		re, err = regexp.Compile(`\d+`)
		if err != nil {
			panic("Error: " + err.Error())
			return nil
		}
		if re.MatchString(col) {
			var re = regexp.MustCompile(`(.{3})\s*$`)
			token = re.ReplaceAllString(col, "***")
			newColumn = append(newColumn, token)
			continue
		}

		auxToken := strings.Split(col, " ")
		token = ""

		for _, aux := range auxToken {
			if len(aux) > 3 {
				var re = regexp.MustCompile(`(.{3})\s*$`)
				token += re.ReplaceAllString(aux, "***") + " "

			} else if len(aux) == 3 {
				var re = regexp.MustCompile(`(.{2})\s*$`)
				token += re.ReplaceAllString(aux, "**") + " "
			} else if len(aux) == 2 {
				var re = regexp.MustCompile(`(.{1})\s*$`)
				token += re.ReplaceAllString(aux, "*") + " "
			} else {
				token += aux + " "
			}

		}

		newColumn = append(newColumn, token)

	}

	return newColumn
}

func KAnonymityGeneralizationNumeric(column []float64) []string {
	k := 1 + 3.332*math.Log10(float64(len(column)))
	k = math.Round(k)
	limSup, limInf := kAnonymityFindMaxMin(column)
	a := math.Round((limSup - limInf) / k)
	var val limValues
	val.limInf = limInf
	val.limSup = (limInf + a) - 1
	var class []limValues
	class = append(class, val)
	for i := 1; i < int(k); i++ {
		val.limInf = val.limSup + 1
		val.limSup = (val.limInf + a) + 1
		class = append(class, val)
	}

	var newColumn []string

	for _, c := range column {
		for _, lim := range class {
			if c >= lim.limInf && c < lim.limSup+1 {
				li := fmt.Sprint(lim.limInf)
				ls := fmt.Sprint(lim.limSup)
				newColumn = append(newColumn, li+"-"+ls)
			}
		}
	}

	return newColumn
}

func kAnonymityFindMaxMin(values []float64) (float64, float64) {
	max := values[0]
	min := values[0]
	for _, v := range values {
		max = math.Max(v, max)
		min = math.Min(v, min)
	}

	return max, min
}
