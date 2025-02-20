package utils

import "fmt"

func ShowHelp() {
	fmt.Printf("Usage: webfetch [OPTIONS]\n\n")
	fmt.Println("Options:")
	fmt.Println(" -art=<FILE>	Path to the ASCII art file")
	fmt.Println(" -port=<PORT>	Specify the port number")
	fmt.Println(" -help		Show help message and exit")
}
