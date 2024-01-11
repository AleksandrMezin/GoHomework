package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func calculateExpression(expression string) (int, error) {
	re := regexp.MustCompile(`(\d+)\s*([\+\-*/])\s*(\d+)`)
	matches := re.FindStringSubmatch(expression)

	if len(matches) != 4 {
		return 0, fmt.Errorf("Invalid expression: %s", expression)
	}

	num1, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	num2, err := strconv.Atoi(matches[3])
	if err != nil {
		return 0, err
	}

	operator := matches[2]

	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Division by zero: %s", expression)
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Unsupported operator: %s", operator)
	}
}

func processFile(inputFileName string, outputFileName string) error {
	inputData, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		return err
	}

	lines := regexp.MustCompile(`\r?\n`).Split(string(inputData), -1)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	for _, line := range lines {
		result, err := calculateExpression(line)
		if err == nil {
			output := fmt.Sprintf("%s=%d\n", line, result)
			writer.WriteString(output)
		}
	}

	writer.Flush()

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: program input.txt output.txt")
		os.Exit(1)
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	if _, err := os.Stat(outputFileName); os.IsNotExist(err) {
		_, _ = os.Create(outputFileName)
	}

	err := ioutil.WriteFile(outputFileName, []byte(""), 0644)
	if err != nil {
		fmt.Printf("Error clearing output file: %v\n", err)
		os.Exit(1)
	}

	err = processFile(inputFileName, outputFileName)
	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Results written to", outputFileName)
}
