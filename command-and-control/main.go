//Group: GOROUTINE

// Group Members username:
// Agene Okoh
// Janai Egeonu
// Chibueze Maxwell
// Emmanuel Unogwu
// Edwin Ejembi
// Blessing Anebi
// Faith Ejembi
// Ruth Agi
// Ummulkusum Musa

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var lastResult float64
var hasLast bool
var history []string
var calcHistory []float64

func main() {
	fmt.Println("════════════════════════════════════════════════")
	fmt.Println("  SENTINEL — COMMAND & CONTROL CONSOLE")
	fmt.Println("     All systems nominal. Type 'help' to begin.")
	fmt.Println("════════════════════════════════════════════════")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("C&C> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		runInput(input)
	}
}

func runInput(input string) {
	if strings.Contains(input, "|") {
		runPipe(input)
		return
	}
	formatted, plain := runCommand(input)
	addHistory(input + " → " + plain)
	fmt.Println(formatted)
}

func runPipe(input string) {
	parts := strings.SplitN(input, "|", 2)
	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])

	leftFormatted, leftPlain := runCommand(left)
	addHistory(left + " → " + leftPlain)
	fmt.Println(leftFormatted)

	rightFields := strings.Fields(right)
	if hasLast && len(rightFields) == 2 &&
		(rightFields[0] == "base" || rightFields[0] == "calc") {
		right = right + " " + fmt.Sprintf("%.6g", lastResult)
	}

	rightFormatted, rightPlain := runCommand(right)
	addHistory(right + " → " + rightPlain)
	fmt.Println(rightFormatted)
}

func runCommand(input string) (string, string) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return "", ""
	}

	prefix := strings.ToLower(fields[0])

	switch prefix {
	case "help":
		return helpText(), "help"
	case "history":
		return showHistory(), "history"
	case "clear":
		return doClear(), "clear"
	case "exit":
		fmt.Println("  SENTINEL offline. Goodbye.")
		os.Exit(0)
	case "calc":
		return runCalc(fields[1:])
	case "base":
		return runBase(fields[1:])
	case "str":
		return runStr(fields[1:])
	default:
		return fmt.Sprintf("   Unknown command: %q — type 'help' to see commands.", prefix), "error"
	}

	return "", ""
}

func runCalc(args []string) (string, string) {
	if len(args) == 0 {
		return "   Usage: calc <operation> <a> <b>  (e.g. calc add 5 3)", "error"
	}

	op := strings.ToLower(args[0])

	switch op {
	case "last":
		if !hasLast {
			return "   No previous result in this session.", "no result"
		}
		out := fmt.Sprintf("%.6g", lastResult)
		return fmt.Sprintf("   Last result: %s", out), out

	case "history":
		if len(calcHistory) == 0 {
			return "   No calculation history yet.", "empty"
		}
		var sb strings.Builder
		for i, v := range calcHistory {
			sb.WriteString(fmt.Sprintf("  %d.  %s\n", i+1, fmt.Sprintf("%.6g", v)))
		}
		return strings.TrimRight(sb.String(), "\n"), "calc history"
	}

	if len(args) < 3 {
		return fmt.Sprintf("   'calc %s' needs two numbers. Usage: calc %s <a> <b>", op, op), "error"
	}

	a, err := parseNumber(args[1])
	if err != nil {
		return fmt.Sprintf("   Invalid first argument %q — expected a number.", args[1]), "error"
	}

	b, err := parseNumber(args[2])
	if err != nil {
		return fmt.Sprintf("   Invalid second argument %q — expected a number.", args[2]), "error"
	}

	var result float64

	switch op {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			return "   Error: cannot divide by zero.", "error"
		}
		result = a / b
	case "mod":
		if b == 0 {
			return "   Error: cannot mod by zero.", "error"
		}
		result = math.Mod(a, b)
	case "pow":
		result = math.Pow(a, b)
	default:
		return fmt.Sprintf("   Unknown calc operation %q. Try: add sub mul div mod pow", op), "error"
	}

	out := fmt.Sprintf("%.6g", result)
	lastResult = result
	hasLast = true
	calcHistory = append(calcHistory, result)
	if len(calcHistory) > 5 {
		calcHistory = calcHistory[1:]
	}

	return fmt.Sprintf("   Result: %s", out), out
}

func parseNumber(s string) (float64, error) {
	if strings.ToLower(s) == "last" {
		if !hasLast {
			return 0, fmt.Errorf("no previous result")
		}
		return lastResult, nil
	}
	return strconv.ParseFloat(s, 64)
}

func runBase(args []string) (string, string) {
	if len(args) == 0 {
		return "   Usage: base <dec|hex|bin> <number>", "error"
	}

	op := strings.ToLower(args[0])

	if len(args) < 2 {
		return fmt.Sprintf("   'base %s' needs a number. Usage: base %s <number>", op, op), "error"
	}

	numStr := args[1]

	switch op {
	case "dec":
		n, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			return fmt.Sprintf("   %q is not a valid decimal number.", numStr), "error"
		}
		if n < 0 {
			return "   Negative numbers are not supported for base dec.", "error"
		}
		bin := strconv.FormatInt(n, 2)
		hex := strings.ToUpper(strconv.FormatInt(n, 16))
		lastResult = float64(n)
		hasLast = true
		out := fmt.Sprintf("Binary: %s / Hex: %s", bin, hex)
		return fmt.Sprintf("   Binary : %s\n   Hex    : %s", bin, hex), out

	case "hex":
		upper := strings.ToUpper(numStr)
		for _, c := range upper {
			if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'F')) {
				return fmt.Sprintf("   %q is not valid hex.", numStr), "error"
			}
		}
		n, err := strconv.ParseInt(numStr, 16, 64)
		if err != nil {
			return fmt.Sprintf("   %q is not valid hex.", numStr), "error"
		}
		lastResult = float64(n)
		hasLast = true
		out := fmt.Sprintf("%d", n)
		return fmt.Sprintf("   Decimal: %d", n), out

	case "bin":
		for _, c := range numStr {
			if c != '0' && c != '1' {
				return fmt.Sprintf("   %q is not valid binary.", numStr), "error"
			}
		}
		n, err := strconv.ParseInt(numStr, 2, 64)
		if err != nil {
			return fmt.Sprintf("   %q is not valid binary.", numStr), "error"
		}
		lastResult = float64(n)
		hasLast = true
		out := fmt.Sprintf("%d", n)
		return fmt.Sprintf("   Decimal: %d", n), out

	default:
		return fmt.Sprintf("   Unknown base operation %q. Try: dec hex bin", op), "error"
	}
}

var smallWords = map[string]bool{
	"a": true, "an": true, "and": true, "as": true, "at": true,
	"but": true, "by": true, "for": true, "in": true, "nor": true,
	"of": true, "on": true, "or": true, "so": true, "the": true,
	"to": true, "up": true, "yet": true,
}

func runStr(args []string) (string, string) {
	if len(args) == 0 {
		return "   Usage: str <upper|lower|cap|title|snake|reverse> <text>", "error"
	}

	op := strings.ToLower(args[0])

	if len(args) < 2 {
		return "   No text provided.", "error"
	}

	text := strings.Join(strings.Fields(strings.Join(args[1:], " ")), " ")
	if text == "" {
		return "  No text provided.", "error"
	}

	var result string

	switch op {
	case "upper":
		result = strings.ToUpper(text)
	case "lower":
		result = strings.ToLower(text)
	case "cap":
		wordList := strings.Fields(text)
		for i, w := range wordList {
			if len(w) > 0 {
				wordList[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
			}
		}
		result = strings.Join(wordList, " ")
	case "title":
		wordList := strings.Fields(text)
		for i, w := range wordList {
			lower := strings.ToLower(w)
			if i == 0 || !smallWords[lower] {
				wordList[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
			} else {
				wordList[i] = lower
			}
		}
		result = strings.Join(wordList, " ")
	case "snake":
		wordList := strings.Fields(strings.ToLower(text))
		var cleaned []string
		for _, w := range wordList {
			var sb strings.Builder
			for _, c := range w {
				if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
					sb.WriteRune(c)
				}
			}
			if sb.Len() > 0 {
				cleaned = append(cleaned, sb.String())
			}
		}
		result = strings.Join(cleaned, "_")
	case "reverse":
		wordList := strings.Fields(text)
		for i, w := range wordList {
			runes := []rune(w)
			for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
				runes[l], runes[r] = runes[r], runes[l]
			}
			wordList[i] = string(runes)
		}
		result = strings.Join(wordList, " ")
	default:
		return fmt.Sprintf("  Unknown str operation %q. Try: upper lower cap title snake reverse", op), "error"
	}

	return fmt.Sprintf("  ✦ %s", result), result
}

func addHistory(entry string) {
	history = append(history, entry)
	if len(history) > 10 {
		history = history[1:]
	}
}

func showHistory() string {
	if len(history) == 0 {
		return "  ✗ No history yet."
	}
	var sb strings.Builder
	for i, h := range history {
		sb.WriteString(fmt.Sprintf("  %d.  %s\n", i+1, h))
	}
	return strings.TrimRight(sb.String(), "\n")
}

func doClear() string {
	fmt.Print("  Type CONFIRM to clear the session: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	confirm := strings.TrimSpace(scanner.Text())
	if confirm == "CONFIRM" {
		history = nil
		calcHistory = nil
		lastResult = 0
		hasLast = false
		return "  ✦ Session cleared."
	}
	return "  ✦ Clear cancelled."
}

func helpText() string {
	return `  calc  add|sub|mul|div|mod|pow <a> <b>
  calc  last | history
  base  dec|hex|bin <number>
  str   upper|lower|cap|title|snake|reverse <text>
  history | clear | exit`
}
