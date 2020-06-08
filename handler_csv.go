package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

func handlerCsv(in io.Reader, out io.Writer, fields []string) error {
	csvReader := csv.NewReader(in)
	csvWriter := csv.NewWriter(out)

	header, err := csvReader.Read()
	if err != nil {
		return fmt.Errorf("could not parse csv header: %w", err)
	}

	skip := findCsvSkip(header, fields)

	if err := writeCsvRow(csvWriter, header, skip); err != nil {
		return fmt.Errorf("could not write csv header: %w", err)
	}

	for line := 1; ; line++ {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not decode record on line %d: %w", line, err)
		}

		if err := writeCsvRow(csvWriter, row, skip); err != nil {
			return fmt.Errorf("could not write a csv row on line %d: %w", line, err)
		}
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("could not flush output file: %w", err)
	}

	return nil
}

func findCsvSkip(header, fields []string) []int {
	skip := make([]int, 0, len(fields))
	for _, f := range fields {
		for i, h := range header {
			if f == h {
				skip = append(skip, i)
			}
		}
	}
	sort.Slice(skip, func(i, j int) bool { return skip[i] > skip[j] })

	return skip
}

func writeCsvRow(w *csv.Writer, row []string, skip []int) error {
	for _, i := range skip {
		row = append(row[0:i], row[i+1:]...)
	}

	return w.Write(row)
}
