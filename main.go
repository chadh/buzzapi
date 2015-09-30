package main

import (
	"bufio"
	"os"
)

func main() {
	// parse command line arguments and populate op, res, params, and flags
	initializeArgs()

	// set up file handle for results
	if *output != "" {
		fo, err := os.Create(*output)
		if err != nil {
			log.Critical("%s", err)
		}
		defer func() {
			if err := fo.Close(); err != nil {
				log.Critical("%s", err)
			}
		}()
		outfile = bufio.NewWriter(fo)
	} else {
		outfile = bufio.NewWriter(os.Stdout)
	}
	defer outfile.Flush()

	// set up logging and initialize log
	initializeLogger()
	// parse config file and populate env data structure
	initializeConfig()

	// bundle up the request into a data structure that can be easily marshaled
	createJSONRequest()

	// POST the request to the api, receive the results, and process them
	postRequest()
}
