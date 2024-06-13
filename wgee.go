package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var output string
	flag.StringVar(&output, "O", "", "Specify output file name")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: ./wgee [-O output_file] <URL>")
		os.Exit(1)
	}

	url := flag.Args()[0]
	fullURL := "https://newtrojan.lizhiyin.us.kg/" + url

	var cmd *exec.Cmd

	if output != "" {
		// Use wget with output option
		cmd = exec.Command("wget", "-O", output, fullURL)
	} else {
		// Use wget without output option
		cmd = exec.Command("wget", fullURL)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		// If wget fails, try using curl
		fmt.Println("wget failed, trying with curl...")

		cmd := exec.Command("curl", "-o", output, fullURL)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error running curl: %v\n", err)
			os.Exit(1)
		}
	}
}

