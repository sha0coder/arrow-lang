package main

// subtokenizing
func Tokenize(tok string, l int, t int) []string {
	var subtokens []string

	// remove envelope
	last := len(tok) - 1
	if (tok[0] == '"' && tok[last] == '"') ||
		(tok[0] == '\'' && tok[last] == '\'') ||
		(tok[0] == '[' && tok[last] == ']') {
		tok = tok[1:last]
	}

	accum := ""
	onStr := false
	onVar := false
	onNum := false

	for i := 0; i < len(tok); i++ {
		c := tok[i]

		if !onStr && !onVar && !onNum {

			// detect start of tokens
			accum = ""

			// spaces between tokens
			if c == ' ' {
				continue
			}

			// init string
			if c == '"' || c == '\'' {
				onStr = true
				accum = string(c)
				continue
			}

			// init var
			if c == '$' {
				onVar = true
				accum = string(c)
				continue
			}

			// init num
			switch c {
			case '(', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				{
					onNum = true
					accum = string(c)
					continue
				}
			}

			// comparator
			if IsComparator(tok[i:]) {
				if tok[i] == '<' || tok[i] == '>' {
					subtokens = append(subtokens, "C"+string(tok[i]))
				} else {
					if i+2 > len(tok)-1 {
						Error(l, t, "Incorrect comparator byte "+tok)
					}
					subtokens = append(subtokens, "C"+string(tok[i:i+2]))
				}
				continue
			}

			continue
			//Error(l, t, "Incorrect byte "+string(c))

		} else {
			// subtoken already started

			// end-subtoken conditions

			if onStr {
				if c == '"' || c == '\'' {
					onStr = false
					accum += string(c)
					subtokens = append(subtokens, "S"+ResolveStr(accum, l, t))
				}
			}

			if onVar {
				if c == ' ' {
					onVar = false
					subtokens = append(subtokens, "V"+ResolveNum(accum, l, t))
				}
			}

			if onNum {
				if c == ' ' {
					onNum = false
					subtokens = append(subtokens, "N"+ResolveNum(accum, l, t))
				}
			}

			// a subtoken open
			accum += string(c)

		}
	}

	// push last element

	if onStr {
		subtokens = append(subtokens, "S"+ResolveStr(accum, l, t))
	}

	if onVar {
		subtokens = append(subtokens, "V"+ResolveNum(accum, l, t))
	}

	if onNum {
		subtokens = append(subtokens, "N"+ResolveNum(accum, l, t))
	}

	return subtokens
}
