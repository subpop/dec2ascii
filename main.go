package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %v: convert integers read from stdin to ASCII\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			log.Fatalf("cannot parse value %v: %v", s.Text(), err)
		}
		fmt.Printf("%c", i)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
