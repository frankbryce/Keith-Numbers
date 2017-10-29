package main

import (
	"flag"
	"fmt"
	"keith"
	coll "keith/collection"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	// main idea
	// create "collection" package to provide primitives of datum to collect
	//   about keith sequences
	// generate keith sequences and collect data
	// create dump of files for runs
	// copy into google sheets, for analysis
	// command line parsing
	start := flag.Int("start", 10, "start number for keith analysis")
	end := flag.Int("end", 10000, "end number for keith analysis")
	inc := flag.Int("inc", 1, "amount to increment between numbers")
	flag.Parse()

	// start and increment
	n := big.NewInt(int64(*start))
	one := big.NewInt(int64(*inc))

	// single point to change separater in csv
	filename := fmt.Sprintf("keiths_data_%v_%v_%v.csv", *start, *end, *inc)

	// if file exists, don't do this again, just return immediately
	if _, err := os.Stat(filename); err == nil {
		return
	}

	// If the file doesn't exist, create it
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	sep := "\t"

	// write the header to the file
	header := "num"
	for _, c := range coll.Collections {
		header += sep + strings.TrimPrefix(fmt.Sprintf("%T", c), "*collection.")
	}
	if _, err := f.Write([]byte(header + "\n")); err != nil {
		log.Fatal(err)
	}

	// analysis for all specified values
	for i := *start; i <= *end; i += *inc {
		keith.IsKeithCollect(n)

		row := n.String()
		for _, c := range coll.Collections {
			row += sep + fmt.Sprintf("%v", c.Get())
		}
		if _, err := f.Write([]byte(row + "\n")); err != nil {
			log.Fatal(err)
		}

		n.Add(n, one)
	}
}

func naive() {
	n := big.NewInt(int64(10000000000000))
	n.Mul(n, n)
	n.Mul(n, big.NewInt(10))

	zero := big.NewInt(0)
	one := big.NewInt(1)
	nineteen := big.NewInt(19)
	for mod := big.NewInt(0).Mod(n, nineteen); mod.Cmp(zero) != 0; mod = big.NewInt(0).Mod(n, nineteen) {
		n.Add(n, one)
	}

	c := int64(0)
	for ; !keith.IsKeith(n); n.Add(n, nineteen) {
		if c%100000 == 0 {
			fmt.Println("%v : %v digits", n, len(n.String()))
		}
		c++
	}

	fmt.Println("%v : %v digits", n, len(n.String()))
}
