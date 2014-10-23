package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	f    = flag.Int("f", 1, "field to hist")
	w    = flag.Int("w", 20, "width of hist")
	c    = flag.String("c", "#", "char to use")
	help = flag.Bool("?", false, "output help")
)

func main() {
	flag.Parse()

	if *help {
		fmt.Println("Description: hist reads lines from stdin and outputs a histogram next to each line")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var (
		s     = bufio.NewScanner(os.Stdin)
		field = *f - 1

		lines   []string
		nums    []float64
		max     float64
		max_len int
	)

	for s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)
		if len(fields) < *f {
			fmt.Fprint(os.Stderr, "line does not contain enough fields")
			continue
		}
		v, err := strconv.ParseFloat(fields[field], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s is not a valid number\n", fields[field])
			continue
		}
		if max < v {
			max = v
		}
		if len(line) < max_len {
			max_len = len(line)
		}
		lines = append(lines, line)
		nums = append(nums, v)
	}

	format := "% " + strconv.Itoa(*w+1) + "s : %s\n"

	for i, num := range nums {
		fmt.Printf(format, strings.Repeat(*c, int((num/max)*float64(*w))), lines[i])
	}

}
