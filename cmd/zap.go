package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	args "github.com/ashtyn3/zap/arg-parser"
)

var path string

type valid struct {
	Name  string
	Value string
}

func build() {
	cmd := exec.Command("go", "build")
	cmd.Env = append(os.Environ(), "GOPATH="+os.Getenv("GOPATH"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin
	cmd.Run()
	rm := exec.Command("rm", path)
	rm.Stdout = os.Stdout
	rm.Stderr = os.Stdin
	rm.Run()
	mv := exec.Command("mv", path+".copy", path)
	mv.Stdout = os.Stdout
	mv.Stderr = os.Stderr
	mv.Run()
}

func main() {
	flags := args.ArgParser(os.Args[1:])
	var aval []valid
	if len(flags) == 0 {
		fmt.Println("zap [-env <variables,...>][-in <filename>]")
	}
	for _, flag := range flags {
		if flag.Flag == "-env" {
			if flag.Param != "" {
				evs := strings.Split(flag.Param, ",")
				for _, v := range evs {
					tv := strings.Split(v, "=")
					aval = append(aval, valid{Name: tv[0], Value: tv[1]})
				}
			}
		}
		if flag.Flag == "-in" {
			// fmt.Println(flag.Param)
			// f, _ := os.OpenFile(flag.Param, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			contentB, err := ioutil.ReadFile(flag.Param)
			if err != nil {
				fmt.Println(err)
			}
			er := ioutil.WriteFile(flag.Param+".copy", []byte(contentB), 0644)
			if er != nil {
				fmt.Println(er)
			}
			path = flag.Param
			content := string(contentB)
			sN := strings.Split(content, "\n")
			for i, line := range sN {
				if m, _ := regexp.MatchString(`//\s*@zap\s+var:`, line); m == true {
					if m, _ := regexp.MatchString(`var\s+\S+\s+\w+\s*=`, sN[i+1]); m == true {
						params := strings.Fields(sN[i+1])
						name := params[1]
						tp := params[2]
						r := regexp.MustCompile(`//\s*@zap\s+var:`)
						s := r.Split(line, -1)

						for _, a := range aval {
							if a.Name == strings.Replace(strings.Join(s[1:], ""), " ", "", -1) {
								newVar := fmt.Sprintln("var", name, tp, "= \""+a.Value+"\"")
								sN[i+1] = newVar
							}
						}
					}
				}
			}

			file := strings.Join(sN, "\n")
			ioutil.WriteFile(flag.Param, []byte(file), 0644)
			build()
		}
	}
}
