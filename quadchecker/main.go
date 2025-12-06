package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func quadA(w, h int) string {
	out := ""
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if (x == 0 && y == 0) || (x == w-1 && y == 0) ||
				(x == 0 && y == h-1) || (x == w-1 && y == h-1) {
				out += "o"
			} else if y == 0 || y == h-1 {
				out += "-"
			} else if x == 0 || x == w-1 {
				out += "|"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

func quadB(w, h int) string {
	out := ""
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x == 0 && y == 0 {
				out += "/"
			} else if x == w-1 && y == 0 {
				out += "\\"
			} else if x == 0 && y == h-1 {
				out += "\\"
			} else if x == w-1 && y == h-1 {
				out += "/"
			} else if y == 0 || y == h-1 {
				out += "*"
			} else if x == 0 || x == w-1 {
				out += "*"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

func quadC(w, h int) string {
	out := ""
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x == 0 && y == 0 {
				out += "A"
			} else if x == w-1 && y == 0 {
				out += "A"
			} else if x == 0 && y == h-1 {
				out += "C"
			} else if x == w-1 && y == h-1 {
				out += "C"
			} else if y == 0 || y == h-1 || x == 0 || x == w-1 {
				out += "B"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

func quadD(w, h int) string {
	out := ""
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x == 0 && y == 0 {
				out += "A"
			} else if x == w-1 && y == 0 {
				out += "C"
			} else if x == 0 && y == h-1 {
				out += "A"
			} else if x == w-1 && y == h-1 {
				out += "C"
			} else if y == 0 || y == h-1 || x == 0 || x == w-1 {
				out += "B"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

func quadE(w, h int) string {
	out := ""
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x == 0 && y == 0 {
				out += "A"
			} else if x == w-1 && y == 0 {
				out += "C"
			} else if x == 0 && y == h-1 {
				out += "C"
			} else if x == w-1 && y == h-1 {
				out += "A"
			} else if y == 0 || y == h-1 || x == 0 || x == w-1 {
				out += "B"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	h := len(lines)
	w := len(lines[0])

	for _, line := range lines {
		if len(line) != w {
			fmt.Println("Not a quad function")
			return
		}
	}

	input := ""
	for _, l := range lines {
		input += l + "\n"
	}

	names := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	createstheQuad := []func(int, int) string{
		quadA, quadB, quadC, quadD, quadE,
	}

	var matches []string

	for i := range createstheQuad {
		if createstheQuad[i](w, h) == input {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", names[i], w, h))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	sort.Strings(matches)
	fmt.Println(strings.Join(matches, " || "))
}
