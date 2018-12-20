package main

import (
	"regexp"
	"strconv"
)

func EvalCondition(sign string, a string, b string, l int, t int, tok string) bool {

	fa, err_a := strconv.ParseFloat(a, 32)
	fb, err_b := strconv.ParseFloat(b, 32)

	stringComparison := false
	if err_a != nil || err_b != nil {
		stringComparison = true
	}

	eval := false
	switch sign {
	case "==":
		{
			if stringComparison {
				if a == b {
					eval = true
				}
			} else {
				if fa == fb {
					eval = true
				}
			}
		}

	case "<":
		{
			if stringComparison {
				if a[0] < b[0] {
					eval = true
				}

			} else {
				if fa < fb {
					eval = true
				}
			}
		}

	case "<=":
		{
			if stringComparison {
				if a[0] <= b[0] {
					eval = true
				}
			} else {
				if fa <= fb {
					eval = true
				}
			}
		}

	case ">":
		{
			if stringComparison {
				if a[0] > b[0] {
					eval = true
				}
			} else {
				if fa > fb {
					eval = true
				}
			}
		}

	case ">=":
		{
			if stringComparison {
				if a[0] >= b[0] {
					eval = true
				}
			} else {
				if fa >= fb {
					eval = true
				}
			}
		}

	case "!=":
		{
			if stringComparison {
				if a != b {
					eval = true
				}
			} else {
				if fa != fb {
					eval = true
				}
			}
		}

	case "=~":
		{
			matched, err := regexp.MatchString(b, a)
			if err != nil {
				Error(l, t, "Bad regular expression on the condition: "+tok)
			}
			if matched {
				eval = true
			}

		}

	case "!~":
		{
			matched, err := regexp.MatchString(b, a)
			if err != nil {
				Error(l, t, "Bad regular expression on the condition: "+tok)
			}
			if !matched {
				eval = true
			}
		}

	} // switch

	return eval
}
