package main

import "os"

func main() {
	defer func() {
		if recover() != nil {
			os.Exit(1)
		}
	}()

	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("Usage : ./verhoeff <input>\n"))
		return
	}

	result, err := Generate(os.Args[1])

	if err != nil {
		os.Stdout.Write([]byte("Invalid input"))
		os.Exit(1)
		return
	}

	os.Stdout.Write([]byte(result))

	os.Exit(0)
}
