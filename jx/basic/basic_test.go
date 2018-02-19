package jxbasic

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
    "os/exec"
	"strings"
	"errors"
	"bytes"
	"fmt"
	"net/url"
)

type Jx struct {
	StdOut string
}

type Command struct{
	Name string
	Args []string
}

func (jx *Jx) reset(interface{}) {
	jx.StdOut = ""
}


func (jx *Jx) executing(cmd string) (err error) {

	    c := getCommand(cmd)

		var out bytes.Buffer
		var exe = &exec.Cmd{}
		if len(c.Args) == 0 {
			exe = exec.Command(c.Name,"")
		}else{
			exe = exec.Command(c.Name,c.Args...)
		}
		exe.Stdout = &out
		exe.Stdin = bytes.NewBufferString("\n")
		err = exe.Run()
		output := string(out.Bytes()[:])
		jx.StdOut = output



	return err

}

func (jx *Jx) stdoutShouldContain(body *gherkin.DocString) (err error) {
	if !strings.Contains(jx.StdOut,body.Content) {
		err = errors.New("No " + body.Content + " found in " + jx.StdOut)
	}
	return
}

func executingMyWebBrowserShouldOpenAPageToTheJenkinsConsoleOfJenkinsRunningInMyCluster(cmd string) error {

	setNamespace("jx")
	c := getCommand(cmd)

	//First ensure we are ru
	out,err := exec.Command(c.Name, c.Args...).Output()


	if err != nil {
		fmt.Printf("Error: Command failed  %s %s\n", c.Name, strings.Join(c.Args, " "))
	}else {
		fmt.Printf("casptured: %s",out)
		if len(out) > 0 {
			output := string(out[:])
			Url := strings.Split(output,"http")
			if Url != nil {
				UrlString := "http" + Url[1]
				UrlString = strings.TrimSpace(UrlString)
				u, error := url.ParseRequestURI(UrlString)
				if error != nil{
					err = error
				}else if u == nil {
					err = errors.New("Invalid url:" + UrlString)
				}
			}
			fmt.Println("URL = " + Url[1])
		}else {
			err = errors.New("no out put from " + cmd)
		}
	}
	return err
}

func  getCommand(cmd string) *Command {
	array := strings.Split(cmd," ")
	var name string
	args := make([]string,0)

		name = array[0]
		if len(array) > 1 {
			for i := 1; i < len(array); i++ {
				args = append(args,array[i])
			}
		}

		return &Command{
			name,
			 args,
		}

}



func setNamespace(name string){
	jx := &Jx{}
	jx.executing("jx namespace "+ name)
}

func FeatureContext(s *godog.Suite) {
	jx := &Jx{}

	s.BeforeScenario(jx.reset)
	s.Step(`^executing "([^"]*)"$`, jx.executing)
	s.Step(`^stdout should contain$`, jx.stdoutShouldContain)
	s.Step(`^executing "([^"]*)" my web browser should open a page to the Jenkins console of Jenkins running in my cluster$`, executingMyWebBrowserShouldOpenAPageToTheJenkinsConsoleOfJenkinsRunningInMyCluster)
}