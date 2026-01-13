package main

import (
	"encoding/csv"
	"os"
)

func loadRecipient(path string, ch chan Recipient) error {
	defer close(ch)
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	file.Close()
	for _, record := range records[1:] {
		//channel will be block until somene is not consume this channel
		ch <- Recipient{ // Instance create
			Name:  record[0],
			Email: record[1],
		}

	}
	return nil
}
