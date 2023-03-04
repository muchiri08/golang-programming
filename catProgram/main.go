package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var tlc, twc, tcc int
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		//if _, err := io.Copy(os.Stdout, file); err != nil {
		//	fmt.Fprint(os.Stderr, err)
		//}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}
		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
		tlc += lc
		twc += wc
		tcc += cc

		file.Close()
	}

	fmt.Printf("%7d %7d %7d totals\n", tlc, twc, tcc)
}
