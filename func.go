package main

import (
	"strconv"
)

var FuncNonReturn = []string{"print"}

func IsFuncRet(function string) bool {
	for _, f := range FuncNonReturn {
		if f == function {
			return false
		}
	}
	return true
}

func CallFunc(p []string, l int, t int) bool {
	funcName := p[0]

	// resolve vars and operation on params
	/* ya se estan resolviendo
	for i := 1; i < len(p); i++ {
		p[i] = ResolveNum(ResolveStr(p[i], l, t), l, t)
	}*/

	switch funcName {
	case "out":
		{
			PY.Push("sys.stdout.write(_)")
			Trace(l, t, "calling out")
			return true
		}
	case "in":
		{
			PY.Push("_ = raw_input().strip()")
			return true
		}
	case "sz":
		{
			PY.Push("__=_")
			PY.Push("_ = len(__)")
			Trace(l, t, "calling sz")
			return true
		}
	case "end":
		{
			PY.Push("sys.exit(1)") //TODO: que no sea necesario pasarle un numero
			Trace(l, t, "calling end")
			return true
		}
	case "pass":
		{
			PY.Push("pass")
			return true
		}
	case "get":
		{
			if len(p) != 2 {
				Error(l, t, "get requires one arg with an regexp")
			}
			if p[1][0] != '\'' && p[1][0] != '"' {
				Error(l, t, "invalid regex on get")
			}
			PY.Push("_ = re.findall(" + p[1] + ", _)")
			return true
		}

	case "grep": // array -> grep reg
		{
			if len(p) != 2 {
				Error(l, t, "grep requires one arg with an regexp")
			}
			if p[1][0] != '\'' && p[1][0] != '"' {
				Error(l, t, "invalid regex on grep")
			}
			PY.Push("_ = filter(lambda __:re.findall(" + p[1] + ",__),_)")
			return true
		}

	case "post":
		{
			if len(p) != 2 {
				Error(l, t, "post requires one string argument")
			}
			if p[1][0] != '\'' && p[1][0] != '"' {
				Error(l, t, "invalid string on post")
			}
			PY.Push("_res = requests.post(" + p[1] + ",data=_,headers={'UserAgent':_ua})")
			PY.Push("_code = res.status_code; _sz = len(_res.text); _ = _res.text")
			return true
		}

	case "load":
		{
			PY.Push("__=_;fd=open(__);_=fd.read();fd.close()")
		}
	case "save":
		{
			PY.Push("fd=open(" + ResolveStr(p[1], l, t) + ");_=fd.read();fd.close()")
		}

	case "rand":
		{
			if len(p) != 3 {
				Error(l, t, "RAND function need two numeric arguments or numeric vars")
			}

			f1, err1 := strconv.ParseFloat(p[1], 32)
			if err1 != nil {
				Error(l, t, "on RAND function, first parm is not a number")
			}
			f2, err2 := strconv.ParseFloat(p[2], 32)
			if err2 != nil {
				Error(l, t, "on RAND function, second parm is not a number")
			}
			if f2 < f1 {
				Error(l, t, "RAND second arg is bigger than first arg")
			}

			PY.Push("_ = random.randint(" + p[1] + "," + p[2] + ")")
			return true
		}
	case "return":
		{
			PY.Push("return _")
			return true
		}

	case "break":
		{
			PY.Push("break")
			return true
		}
	case "cont":
		{
			PY.Push("continue")
			return true
		}
	case "continue":
		{
			PY.Push("continue")
			return true
		}

	case "sort":
		{
			PY.Push("_.sort()")
			return true
		}
	case "uniq":
		{
			PY.Push("_ = list(set(_))")
			return true
		}
	case "range":
		{
			if len(p) <= 1 {
				Error(l, t, "not enought params for a range")
			}
			cmd := "_=range("
			for i, pp := range p {
				if i > 0 {
					cmd += ","
				}
				cmd += pp
			}
			cmd += ")"
			PY.Push(cmd)
			return true
		}
	case "replace":
		{
			if len(p) != 3 {
				Error(l, t, "REPLACE function need two string arguments, regexp and string, and none argumets where given")
			}
			regex := p[1]
			subs := p[2]
			PY.Push("_ = re.sub(" + regex + "," + subs + ",_)")
			return true

		}
	case "list":
		{
			// list 1,2,3,4 -> $a

			if len(p) != 2 {
				PY.Push("_=[]")
			} else {
				PY.Push("_=[" + p[1] + "]")
			}
			return true
		}

	case "append":
		{
			if len(p) != 2 {
				Error(l, t, "filename missing on APPEND function")
			}
			PY.Push("fd=open(" + p[1] + ",'a+'); fd.write(_); fd.close()")
			return true
		}
	case "join":
		{
			if len(p) != 2 {
				Error(l, t, "JOIN function need one string argument, and none argumets where given")
			}
			//PY.Push("__=_")
			PY.Push("_=" + p[1] + ".join(_)")
			return true
		}
	case "split":
		{
			if len(p) != 2 {
				Error(l, t, "SPLIT function need one string argument, and none argumets where given")
			}
			if p[1][0] != '\'' && p[1][0] != '"' {
				Error(l, t, "SPLIT provided arg is not string")
			}
			PY.Push("__=_")
			PY.Push("_ = __.split(" + p[1] + ")")
			return true
		}

	case "push":
		{
			if len(p) != 2 {
				Error(l, t, "SPLIT function need one string argument, and none argumets where given")
			}
			arr := p[1]
			PY.Push(arr + ".append(_); _=" + arr)
			return true
		}
	case "pop":
		{
			if len(p) != 2 {
				Error(l, t, "SPLIT function need one string argument, and none argumets where given")
			}
			if p[1][0] != '$' {
				Error(l, t, "POP needs an array parameter")
			}
			arr := p[1][1:]
			PY.Push("_ = " + arr + ".pop()")
			return true
		}
	case "connect":
		{
			PY.Push("(_host,_port) = _.split(':')")
			PY.Push("_= socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)")
			PY.Push("_.connect((_host, int(_port)))")
			return true
		}
	case "recv":
		{
			PY.Push("__=_")
			if len(p) == 1 {
				PY.Push("_ = __.recv()")
			} else if len(p) > 1 {
				PY.Push("_ = __.recv(" + p[1] + ")")
			}
			return true
		}
	case "send":
		{
			PY.Push("_ = " + p[1] + ".send(_)")
			return true
		}

	case "close":
		{
			PY.Push("_.close()")
			return true
		}

	}

	return false
}
