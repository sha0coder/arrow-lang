package main

import "strings"

const (
	SignSum  = 1
	SignSub  = 2
	SignMult = 3
	SignDiv  = 4
	SignNum  = 5
	SignErr  = 6
)

func NumType(b byte) int {
	switch b {
	case '+':
		return SignSum
	case '-':
		return SignSub
	case '*':
		return SignMult
	case '/':
		return SignDiv
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		return SignNum
	}
	return SignErr
}

func ResolveNum(tok string, l int, t int) string {
	return strings.Replace(tok, "$", "", -1)
} 

/*
	num := ResolveStr(tok, l, t)

	var faccum float32 = 0
	var elm []float32
	var opers []int
	accum := ""

	for i := 0; i < len(num); i++ {
		ty := NumType(num[i])

		if ty == SignErr {
			Error(l, t, "Bad numeric or operation")
		}

		if ty == SignNum {
			accum += string(num[i])
		} else {
			if accum == "" {
				Error(l, t, "Sign without previous number")
			}

			f, err := strconv.ParseFloat(accum, 32)
			if err != nil {
				Error(l, t, "bad float number: ("+accum+")")
			}

			elm = append(elm, float32(f))
			accum = ""

			opers = append(opers, ty)
		}
	}

	if accum == "" {
		Error(l, t, "last subtoken of the operation is not a number")
	}

	f, err := strconv.ParseFloat(accum, 32)
	if err != nil {
		Error(l, t, "bad float number: ("+accum+")")
	}
	elm = append(elm, float32(f))

	faccum = elm[0]
	ptr := 1

	for _, op := range opers {
		switch op {
		case SignSum:
			faccum += elm[ptr]
		case SignSub:
			faccum -= elm[ptr]
		case SignMult:
			faccum *= elm[ptr]
		case SignDiv:
			faccum /= elm[ptr]
		}
		ptr++
	}

	accum = fmt.Sprintf("%f", faccum)
	Trace(l, t, "resolving numbers "+tok+" to "+accum)
	return accum
}*/
