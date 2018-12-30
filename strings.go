package main

func ResolveStr(tok string, l int, t int) string {
	str := ""
	varMode := false
	varName := ""
	var varNames []string

	if tok[0] == '"' && tok[len(tok)-1] != '"' {
		Error(l, t, "Unclosed literal "+tok)
	}
	if tok[0] == '\'' && tok[len(tok)-1] != '\'' {
		Error(l, t, "Unclosed literal "+tok)
	}

	for i := 0; i < len(tok); i++ {

		// enable var mode on $ except \$ and ending $
		if !varMode && i < len(tok)-2 && tok[i] == '$' {
			//fmt.Printf("entra  %s   %d %d \n", tok, i, len(tok))
			if i == 0 || (i > 0 && tok[i-1] != '\\') {
				varMode = true
				varName = ""
			}
		}

		// disable var mode
		// it must:
		//  - varMode enabled
		//  - not at first digit of var $
		//  - at a terminator (space or $)

		if varMode && varName != "" && (tok[i] == ' ' || tok[i] == '$' || i == len(tok)-1) {

			// add last letter of the variable to varName except space and $
			//if i == len(tok)-1 && tok[i] != '$' && tok[i] != ' ' {
			//	varName += string(tok[i])
			//}

			// remove the $
			varName = varName[1:]

			str += "%s"
			varNames = append(varNames, varName)

			varName = ""
			varMode = false
			if tok[i] == '$' {
				continue
			}

			if i == len(tok)-1 && tok[i] != '$' && tok[i] != ' ' {
				str += string(tok[i])
				continue
			}
		}

		if varMode {
			varName += string(tok[i])
		} else {
			str += string(tok[i])
		}
	}

	Trace(l, t, "Resolving str from "+tok+" to '"+str+"'")

	// if it isn't an string literal
	if str[0] != '"' && str[0] != '\'' {
		return str
	}

	// if it's a string

	str = "'" + str[1:len(str)-1] + "'"

	if len(varNames) > 0 {
		str += " % ("

		for i, v := range varNames {
			if i > 0 {
				str += ","
			}
			str += "str(" + v + ")"
		}

		str += ")"
	}

	//TODO resolve vars
	return str
}
