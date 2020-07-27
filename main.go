package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
func between(value string, a string, b string) (string, int, int) {
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return "", 0, 0
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return "", 0, 0
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return "", 0, 0
	}
	return value[posFirstAdjusted:posLast], posFirstAdjusted, posLast
}
func preventNested(Text string) (OutText string, FunctionArray []string) {
	OutText = Text
	for {
		Function, start, end := between(OutText, "{", "}")
		if Function == "" {
			break
		} else {
			FunctionArray = append(FunctionArray, Function)
			OutText = OutText[0:start-1] + OutText[end+1:]
		}
	}
	return
}

// RunCommand is the function for proccessing a command
func RunCommand(RawText string) error {
	OutText, FunctionArray := preventNested(RawText)
	Text := strings.Split(OutText, "-")
	Command := Text[0]
	if strings.Contains(Command, "notOk") {
		openbrowser("https://www.awesomeinventions.com/funny-pictures-brighten-your-day/")
	} else if strings.Contains(Command, "ok") {
		fmt.Println("OK!")
	} else if strings.Contains(Command, "promptOk") {
		Argument := Text[1]
		GoodFunction := FunctionArray[0]
		BadFunction := FunctionArray[1]
		var Response string
		fmt.Print(Argument + " Are you ok? [y/n] ")
		fmt.Scan(&Response)
		if Response == "y" {
			fmt.Println("Good!")
			RunCommand(GoodFunction)
		} else {
			fmt.Println(":(")
			RunCommand(BadFunction)
		}
	} else if strings.Contains(Command, "print") {
		Argument := Text[1]
		fmt.Print(Argument)
	} else if strings.Contains(Command, "printOK") {
		Argument := Text[1]
		fmt.Println(Argument)
	} else {
		return errors.New("Invalid Command")
	}
	return nil
}
func main() {
	file, err := os.Open("./main.okng")
	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner.Split(bufio.ScanLines)
	var LineNumer uint32
	for scanner.Scan() {
		LineNumer++
		err := RunCommand(scanner.Text())
		if err != nil {
			log.Fatal(err.Error() + " on line " + fmt.Sprint(LineNumer))
		}
	}
}
