package main

func main() {
	// parse command line arguments and populate op, res, params, and flags
	initializeArgs()
	// set up logging and initialize log
	initializeLogger()
	// parse config file and populate env data structure
	initializeConfig()

	// bundle up the request into a data structure that can be easily marshaled
	createJSONRequest()

	// POST the request to the api, receive the results, and process them
	postRequest()
}
