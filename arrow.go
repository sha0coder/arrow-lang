package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var PY *PYTHON
var TraceEnabled = false
var Vars map[string]string

func Check(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func Trace(line int, tok int, msg string) {
	if TraceEnabled {
		fmt.Printf("trace %d:%d %s\n", line+1, tok+1, msg)
	}
}

func loadLines(filename string) []string {
	raw, err := ioutil.ReadFile(filename)
	Check(err, "not a source file or permissions problem file:"+filename)
	return strings.Split(string(raw), "\n")
}

func Error(line int, token int, msg string) {
	fmt.Printf("%d:%d Error: %s", (line + 1), (token + 1), msg)
	os.Exit(-1)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("USAGE:\n  %s [source code file]  <trace>", os.Args[0])
		os.Exit(-2)
	}

	qsourcefile := os.Args[1]

	//qsourcefile := "code.arr"

	PY = new(PYTHON)
	PY.Init()

	if len(os.Args) >= 3 && os.Args[2] == "trace" {
		TraceEnabled = true
	}

	if qsourcefile == "" {
		fmt.Println("try --help")
		os.Exit(1)
	}

	Vars = make(map[string]string)

	code := loadLines(qsourcefile)
	async_pending := ""
	for i, l := range code {

		l = PY.GetIndent(l)
		if len(l) == 0 || l[0] == ';' {
			continue
		}

		if strings.HasPrefix(l, "-> ") {
			l = "$_ " + l
		}

		tokens := strings.Split(l, " -> ")

		for j, tok := range tokens {
			tok = strings.TrimRight(tok, " ")
			ty := GetTokType(tok)

			switch ty {

			case TokVar:
				{
					if j == 0 {
						PY.Push(fmt.Sprintf("_=%s", tok[1:]))
					} else {
						PY.Push(fmt.Sprintf("%s=_", tok[1:]))
					}

				}

			case TokStr:
				{
					str := ResolveStr(tok, i, j)
					PY.Push("_=" + str)
					PY.Push("_sz=len(" + str + ")")

					Trace(i, j, "str "+tok+" resolved to "+str)
				}

			case TokNum, TokOper:
				{
					num := ResolveNum(tok, i, j)
					PY.Push("_=" + num)
					Trace(i, j, "number "+tok+" resolved to "+num)
				}

			case TokIf:
				{
					if tok == "[]" {
						PY.Push("else:")
						continue
					}

					last := len(tok) - 1

					// validaitons
					if tok[last] != ']' || len(tok) < 5 {
						Error(i, j, "Unstructured condition: "+tok)
					}

					// err last line an if

					if i >= len(code)-1 {
						Error(i, j, "program cant end in a condition")
					}
					if len(code[i+1]) <= 2 {
						Error(i, j, "invalid expression after the condition")
					}

					if code[i+1][0] != ' ' && code[i+1][0] != '\t' {
						Error(i, j, "Next line to a condition must be indented "+tok)
					}

					res := ""
					stokens := Tokenize(tok, i, j)

					for k := 0; k < len(stokens); k++ {
						stok := stokens[k]
						typ := stok[0]
						stok = stok[1:len(stok)] //TODO: Error()

						// next token will be a regex?
						if k < len(stokens)-1 {
							if stokens[k+1] == "C=~" {
								if typ == 'C' {
									Error(i, j, "duplicated comparator")
								}

								res += "re.findall(" + stokens[k+2][1:] + "," + stok + ")"
								k += 2
								continue
							}
						}

						switch typ {
						case 'S':
							{

								res += stok
							}
						case 'V':
							{

								res += stok
							}
						case 'N':
							{

								res += stok
							}
						case 'C':
							{
								if stok == "=~" {
									Error(i, j, "WTF")

								} else {
									if stok == "&&" {
										stok = "and"
									} else if stok == "||" {
										stok = "or"
									}

									res += " " + stok + " "
								}
							}
						}
					}

					PY.Push("if " + res + ":")
					//TODO: inline ifs
				}

			case TokArr:
				{
					Error(i, j, "arrays unimplemented for now")
				}

			case TokIter:
				{
					/*
						=>    	endless loop
						(3) =>		3 iterations with 3 in _
						($sz) =>	num iterations
						@sz =>	iterate array
					*/

					if tok == "=>" {
						// endless loop
						PY.Push("while True:")
					} else {
						p := strings.Split(tok, " ")
						if len(p) == 1 {
							Error(i, j, "iterator without spaces "+tok)
						}

						if tok[0] == '(' {
							// num of iterations
							PY.Push("for _ in xrange(int" + ResolveNum(p[0], i, j) + "):")
						}
						if tok[0] == '$' {
							// array iterator
							PY.Push("for _ in " + ResolveNum(p[0], i, j) + ":")
						}
					}
				}

			case TokEndAsyncIter:
				{
					if async_pending == "" {
						Error(i, j, "incorrect asyncronous iteration closing token")
					}
					PY.Push(async_pending)
					PY.Push("[_t.start() for _t in _th]")
					PY.Push("[_t.join() for _t in _th]")
				}

			case TokAsyncIter:
				{
					/*
						(3) :=>		3 iterations with 3 in _
						($sz) :=>	num iterations
						$sz :=>	iterate array
					*/

					p := strings.Split(tok, " ")
					if len(p) == 1 {
						Error(i, j, "iterator without spaces "+tok)
					}

					if tok[0] == '(' {
						// num of iterations

						n := strconv.Itoa(rand.Intn(90000))
						async_func_name := "async_func_" + n
						PY.Push("_th = []")
						PY.Push("def " + async_func_name + "(_):")
						async_pending = "[_th.append(Thread(target=" + async_func_name + ", args=(x,))) for x in xrange(int" + ResolveNum(p[0], i, j) + ")]"

					}
					if tok[0] == '$' {
						// array iterator
						n := strconv.Itoa(rand.Intn(90000))
						async_func_name := "async_func_" + n
						PY.Push("_th = []")
						PY.Push("def " + async_func_name + "(_):")
						async_pending = "[_th.append(Thread(target=" + async_func_name + ", args=(x,))) for x in " + ResolveNum(p[0], i, j) + "]"
					}

				}

			case TokFile:
				{
					if j == 0 {
						// read mode
						PY.Push("fd = open('" + tok + "','rb'); _=fd.read(); fd.close(); _sz=len(_)")
					} else {
						// write mode
						PY.Push("fd = open('" + tok + "','wb'); fd.write(_.encode('utf-8')); fd.close()")
					}
				}

			case TokURL:
				{
					tok = ResolveStr("'"+tok+"'", i, j)
					PY.Push("r = requests.get(" + tok + ",headers={'UserAgent':_ua})")
					PY.Push("_code = r.status_code")
					PY.Push("_ = r.text")
					PY.Push("_sz = len(_)")
				}

			case TokExec:
				{
					//TODO: stdin
					if len(tok) < 2 {
						Error(i, j, "command execution with ! but not enought bytes")
					}
					tok = ResolveStr("'"+tok[1:]+"'", i, j)
					PY.Push("__=_")
					PY.Push("_p = subprocess.Popen(" + tok + ",shell=True,stdout=subprocess.PIPE,stdin=subprocess.PIPE)")
					PY.Push("_p.stdin.write(_); _p.stdin.close()")
					PY.Push("_ = _p.stdout.read()")
					PY.Push("_p.wait()")
				}

			case TokFuncDef:
				{
					tok = strings.Replace(tok[2:], "$", "", -1)
					PY.Push("def " + tok + ":")
				}

			case TokFunc:
				{
					//tok = ResolveStr(tok, i, j)
					p := strings.Split(tok, " ")

					// resolver
					for k, subtok := range p {
						typ := GetTokType(subtok)
						if typ == TokStr {
							p[k] = ResolveStr(subtok, i, j)
						} else {
							p[k] = ResolveNum(subtok, i, j)
						}
					}

					call := ""

					if CallFunc(p, i, j) {
						continue
					}

					/*
						p[0] -> program name
						_ -> param 1
						p[1] -> param 2
						p[2] -> param 3
						...
					*/

					hasRet := IsFuncRet(p[0])

					if hasRet {
						call = "_ = " + p[0] + "(__"
					} else {
						call = p[0] + "(_"
					}

					for k := 1; k < len(p); k++ {
						call += ",str(" + p[k] + ")"
					}
					call += ")"
					if hasRet {
						PY.Push("__=_")
					}
					PY.Push(call)
					//LastElem = CallFunc(LastElem, p, i, j)
				}
			}
		}
	}

	fname := strings.Split(qsourcefile, ".")[0] + ".py"
	//PY.Print()
	PY.Write(fname)
	fmt.Println("./" + fname)
	//PY.Exec()
}
