package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func isValidInputString(s string) bool {

	if len(s) == 0 || s[0] == ' ' || s[len(s)-1] == ' ' {
		return false
	}

	for _, char := range s {
		if !(unicode.IsDigit(char) || char == '-' || char == ' ') {
			return false
		}
	}

	substrings := []string{"--", "  ", "- "}
	for _, substring := range substrings {
		if strings.Contains(s, substring) {
			return false
		}
	}
	return true
}

func isValidInputField(s []string) bool {
	for _, str := range s {
		if len(str) > 0 && str[0] == '0' {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	tStr, err := in.ReadString('\n')
	if err != nil {
		fmt.Fprintln(out, "no")
		return
	}
	t, err := strconv.Atoi(strings.TrimSpace(tStr))
	if err != nil || t <= 0 {
		fmt.Fprintln(out, "no")
		return
	}

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(in, &n)

		line1, _ := in.ReadString('\n')
		line2, _ := in.ReadString('\n')

		if len(line1) != len(line2) || isValidInputString(line2) {
			fmt.Fprintln(out, "no")
			continue
		}

		elements1 := strings.Fields(line1)
		elements2 := strings.Fields(line2)

		if len(elements1) != n || len(elements2) != n || !isValidInputField(elements2) {
			fmt.Fprintln(out, "no")
			continue
		}

		numbers1 := make([]int, n)
		numbers2 := make([]int, n)
		match := true

		for j := 0; j < n; j++ {
			num1, err1 := strconv.Atoi(elements1[j])
			num2, err2 := strconv.Atoi(elements2[j])
			if err1 != nil || err2 != nil {
				match = false
				break
			}
			numbers1[j] = num1
			numbers2[j] = num2
		}

		if !match {
			fmt.Fprintln(out, "no")
			continue
		}
		sort.Ints(numbers1)
		match = true
		for idx := 0; idx < n; idx++ {
			if numbers1[idx] != numbers2[idx] {
				match = false
				fmt.Fprintln(out, "no")
				break
			}
		}
		if match {
			fmt.Fprintln(out, "yes")
		}
	}
}
