package main

import (
	"fmt"
	"github.com/inancgumus/prettyslice"
)

const pageSize = 4

func main() {
	realMadridPlayer := []string{
		"Camavinga", "Tchouameni", "Bellingham", "Valverde",
		"Vini Jr", "Rodrygo", "Modric", "Kross",
		"Carvajal",
	}

	prettyslice.MaxPerLine = pageSize

	l := len(realMadridPlayer)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize

		if to > l {
			to = l
		}

		page := fmt.Sprintf("Page #%d", from/pageSize+1)
		prettyslice.Show(page, realMadridPlayer[from:to])
	}
}
