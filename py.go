package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

type PYTHON struct {
	Code     []string
	indent   int
	filename string
}

func (py *PYTHON) Init() {
	py.Push("#!/usr/bin/env python\n")
	py.Push("import re")
	py.Push("import os")
	py.Push("import sys")
	py.Push("import socket")
	py.Push("import random")
	py.Push("import requests")
	py.Push("import subprocess")
	py.Push("from threading import Thread")
	py.Push("")
	PY.Push("_ua='Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36'")
	PY.Push("args=sys.argv")
	PY.Push("_=''")
}

func (py *PYTHON) Push(line string) {
	for i := 0; i < py.indent; i++ {
		line = " " + line
	}
	py.Code = append(py.Code, line)
}

func (py *PYTHON) Write(filename string) {
	py.filename = filename
	ioutil.WriteFile(filename, []byte(strings.Join(py.Code, "\n")), 0755)
}

func (py PYTHON) Print() {
	fmt.Println(strings.Join(py.Code, "\n"))
}

func (py *PYTHON) GetIndent(line string) string {
	py.indent = 0
	for i := 0; i < len(line); i++ {
		if line[i] != ' ' && line[i] != '\t' {
			break
		}
		py.indent++
	}

	line = line[py.indent:]
	return line
}

func (py PYTHON) Exec() {
	cmd := exec.Command("python", py.filename)
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
