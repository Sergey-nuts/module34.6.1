package repkg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`([0-9])+([+-\\*]){1}([0-9])+=`)
)

// Calc вычисляет математические выражения из файла in используя метод Scan
// и буферизировано записывает в out
func CalcScan(in *os.File, out *os.File) {

	writer := bytes.Buffer{} // writer для временного считывания

	scanner := bufio.NewScanner(in) // сканнер файла in
	for scanner.Scan() {
		writer.WriteString(scanner.Text())
	}

	// fmt.Println(wr.String())
	calcRegex := re.FindAllStringSubmatch(writer.String(), -1)
	// fmt.Println(calcRegex)

	outBuf := bufio.NewWriter(out) // буфер для записи в файл
	for _, s := range calcRegex {
		var res int
		a, _ := strconv.Atoi(s[1])
		b, _ := strconv.Atoi(s[3])
		switch s[2] {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		}
		fmt.Fprintf(outBuf, "%v%v\n", s[0], res)
	}
	err := outBuf.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

// Calc2 вычисляет математические выражения из файла in используя метод ReadLine
// и буферизировано записывает в out
func CalcReadLine(in *os.File, out *os.File) {
	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		// fmt.Println("line:", string(line))
		calcRegex := re.FindAllStringSubmatch(string(line), -1)
		// fmt.Println("regex:", calcRegex)

		for _, s := range calcRegex {
			var res int
			a, _ := strconv.Atoi(s[1])
			b, _ := strconv.Atoi(s[3])
			switch s[2] {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			// fmt.Println("res:", res)
			fmt.Fprintf(writer, "%v%v\n", s[0], res)
		}
	}
	err := writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
