package test

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func RunTest(name string, t *testing.T) {
	cmdBuild := exec.Command("go", "build", name+".go")
	err := cmdBuild.Run()
	if err != nil {
		t.Errorf("build failed")
	}

	file, err := os.Open(name + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	start := time.Now()
	defer os.Remove(name)
	defer log.Printf("%s: %.05f s", name, float64(time.Since(start).Nanoseconds())/1000000000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "__IN__" {
			break
		}
	}
	i := 0
	for scanner.Scan() {
		var buf bytes.Buffer
		buf.WriteString(scanner.Text())
		buf.WriteString("\n")
		for scanner.Scan() {
			if scanner.Text() == "__OUT__" {
				break
			}
			buf.WriteString(scanner.Text())
			buf.WriteString("\n")
		}
		input := buf.String()
		buf.Reset()
		for scanner.Scan() {
			line := scanner.Text()
			if line == "__IN__" {
				break
			}
			buf.WriteString(scanner.Text())
			buf.WriteString("\n")
		}
		output := strings.TrimRight(buf.String(), "\n")
		cmd := exec.Command("./" + name)
		cmd.Stdin = strings.NewReader(input)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			t.Errorf("test number %v failed %q ", i+1, err)
		}
		erres := strings.TrimRight(stderr.String(), "\n")
		res := strings.TrimRight(out.String(), "\n")
		if res != output {
			if len(erres) > 0 {
				fmt.Println("os.Stderr>", erres)
			}
			t.Fatalf("TEST %d\nOUTPUT\n%s\nINPUT\n%s"+
				"EXPECTED\n%s\n\n", i+1, res, input, output)
		} else {
			fmt.Printf("Test %d PASS\n", i+1)
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
