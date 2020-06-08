package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dimchansky/utfbom"
)

func handler(in string, out string, fields []string, version bool) error {
	if version {
		return handlerVersion()
	}

	inReader, err := os.Open(in)
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer inReader.Close()

	inReaderWithNoBom := utfbom.SkipOnly(inReader)

	_, err = os.Stat(out)
	if !os.IsNotExist(err) {
		return fmt.Errorf("output file already exists")
	}

	outWriter, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("could not create output file: %w", err)
	}
	defer outWriter.Close()

	ext := filepath.Ext(in)

	switch ext {
	case ".json":
		return handlerNdjson(inReaderWithNoBom, outWriter, fields)
	case ".csv":
		return handlerCsv(inReaderWithNoBom, outWriter, fields)
	}

	return fmt.Errorf("unsupported file type: %s", ext)
}
