package util

import (
	"context"
	"encoding/csv"
	"os"

	"github.com/rs/zerolog/log"
)

type CsvType1 struct {
	Filter [][]string
	Header []string
	Body   [][]string
}

// DownloadCsvType1 generate csv file with format:
// include 3 sections: titles/filters, header of the table, body of the table
func DownloadCsvType1(_ context.Context, data CsvType1) (*os.File, error) {
	f, err := os.CreateTemp("", "tempfile-")
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Filter
	for _, filter := range data.Filter {
		err := writer.Write(filter)
		if err != nil {
			log.Err(err).Send()
			return nil, err
		}
	}

	// Header
	err = writer.Write(data.Header)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	// Body
	for _, record := range data.Body {
		err = writer.Write(record)
		if err != nil {
			log.Err(err).Send()
			return nil, err
		}
	}
	return f, err
}
