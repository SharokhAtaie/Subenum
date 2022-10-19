package commands

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Beauty(str string) string {
	tr := Tr(str)
	sed := Sed(tr)
	return sed
}

// tr
func Tr(str string) string {
	tr := exec.Command("tr", "-d", "[\"]")
	tr.Stdin = strings.NewReader(str)

	var trout bytes.Buffer
	tr.Stdout = &trout
	err := tr.Run()
	if err != nil {
		fmt.Println("tr not working")
	}
	return trout.String()
}

// sed
func Sed(str string) string {
	sed := exec.Command("sed", `s/\,/\n/g`)
	sed.Stdin = strings.NewReader(str)

	var sedout bytes.Buffer
	sed.Stdout = &sedout
	err := sed.Run()
	if err != nil {
		fmt.Println("sed not working")
	}
	return sedout.String()
}

func Sort(str string) string {
	sort := exec.Command("sort", "-u")
	sort.Stdin = strings.NewReader(str)

	var sortout bytes.Buffer
	sort.Stdout = &sortout

	err := sort.Run()
	if err != nil {
		fmt.Println("sort not working")
	}
	return sortout.String()
}

func Sed2(str string, command string) string {
	sed2 := exec.Command("sed", "-e", `s/$/.`+command+`/`)
	sed2.Stdin = strings.NewReader(str)

	var sed2out bytes.Buffer
	sed2.Stdout = &sed2out

	err := sed2.Run()
	if err != nil {
		fmt.Println("sed2 not working")
	}
	return sed2out.String()
}
