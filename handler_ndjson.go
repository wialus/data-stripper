package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/spyzhov/ajson"
)

type jsonRecord = map[string]interface{}

func handlerNdjson(in io.Reader, out io.Writer, fields []string) error {
	inBuffered := bufio.NewReader(in)
	jsonDecoder := json.NewDecoder(inBuffered)

	outBuffered := bufio.NewWriter(out)
	jsonEncoder := json.NewEncoder(outBuffered)

	paths := make([][]string, len(fields))
	for i, f := range fields {
		p, err := ajson.ParseJSONPath(f)
		if err != nil {
			return fmt.Errorf("could not parse jsonpath (\"%s\"): %w", f, err)
		}
		paths[i] = p
	}

	for line := 1; ; line++ {
		record := jsonRecord{}

		err := jsonDecoder.Decode(&record)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not decode record on line %d: %w", line, err)
		}

		for i, p := range paths {
			if err := stripByJsonpath(record, p); err != nil {
				return fmt.Errorf("could not remove a field ('%s'): %w", fields[i], err)
			}
		}

		if err := jsonEncoder.Encode(record); err != nil {
			return fmt.Errorf("could not encode record on line %d: %w", line, err)
		}
	}

	if err := outBuffered.Flush(); err != nil {
		return fmt.Errorf("could not flush output file: %w", err)
	}

	return nil
}

func stripByJsonpath(r interface{}, path []string) error {
	var current interface{}

Loop:
	for ip, p := range path {
		switch p {
		case "$":
			current = r

		case "*":
			s, ok := current.([]interface{})
			if !ok {
				break Loop
			}

			for _, el := range s {
				if err := stripByJsonpath(el, path[ip+1:]); err != nil {
					return err
				}
			}

		case "..":
			return fmt.Errorf("jsonpath operator is not supported: %s", p)

		default:
			m, ok := current.(map[string]interface{})
			if !ok {
				break Loop
			}

			if _, found := m[p]; !found {
				break Loop
			}

			if ip == len(path)-1 {
				delete(m, p)
				break
			}

			current = m[p]
		}
	}

	return nil
}
