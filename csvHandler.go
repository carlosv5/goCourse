package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
)

var (
	reSep     = regexp.MustCompile(",")
	reInteger = regexp.MustCompile("\\d+")
	reFloat   = regexp.MustCompile("[+-]?([0-9]*[.])?[0-9]+")
)

// openCSV opens a CSV file
func openCSV(csvname string) (io.Reader, error) {
	file, err := os.Open(csvname)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// scan scans the lines given in a csvreader
func scan(csvreader io.Reader) func() ([]interface{}, error) {
	scanner := bufio.NewScanner(csvreader)
	return func() ([]interface{}, error) {
		for scanner.Scan() {
			iline := scanner.Text()
			dataline := make([]interface{}, 0)
			fields := reSep.Split(iline, -1)
			for _, field := range fields {
				switch {
				case reFloat.MatchString(field):
					value, _ := strconv.ParseFloat(field, 64)
					dataline = append(dataline, value)
				case reInteger.MatchString(field):
					value, _ := strconv.Atoi(field)
					dataline = append(dataline, value)
				default:
					dataline = append(dataline, field)
				}
			}
			return dataline, nil
		}
		return nil, errors.New("EOF")
	}
}
