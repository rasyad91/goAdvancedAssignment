package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DisplayHeader(s string) {
	fmt.Printf("\n%s\n", s)
	fmt.Println(strings.Repeat("=", len(s)))
}

func InputParserStringToInt(s *string) (int, error) {
	err := InputNotEmpty(s)
	if err != nil {
		return -1, err
	}
	i, err := strconv.Atoi(*s)
	if err != nil {
		return -1, fmt.Errorf("[Invalid input] Please use intergers only.")
	}
	return i, nil
}

func InputParserStringToFloat64(s *string) (float64, error) {
	err := InputNotEmpty(s)
	if err != nil {
		return -1, err
	}
	f, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return -1, fmt.Errorf("[Invalid input] Please use float only.")
	}
	return f, nil
}

func InputNotEmpty(s *string) error {
	if len(*s) == 0 {
		return errors.New("Input required")
	}
	return nil
}

func StringTitle(s *string) error {
	*s = strings.TrimSpace(strings.Title(strings.ToLower(*s)))
	return nil
}

func SingleLineInput(s *string) error {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*s = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	*s = strings.TrimSpace(*s)
	return nil
}

func MultiLinesInput(s *string) error {
	var slice []string
	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Entering multi-line mode. To exit type: !q")
	for scanner.Scan() {
		userInput = scanner.Text()
		if userInput == "!q" {
			break
		}
		slice = append(slice, userInput)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	*s = strings.Join(slice, " ")
	return nil
}

func Contains(s []bool, b bool) bool {
	for _, v := range s {
		if v == b {
			return true
		}
	}
	return false
}
