package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Argument Processing stuff
var (
	app          = kingpin.New("gobuzzapi", "Buzzapi Command Line Interface")
	timeout      = app.Flag("timeout", "Max time to wait (in seconds)").Default("15").Int()
	verbose      = app.Flag("verbose", "Lots of output").Short('v').Bool()
	quiet        = app.Flag("quiet", "Very little output").Short('q').Bool()
	handle       = app.Flag("handle", "Client handle").String()
	clienthandle = app.Flag("client-handle", "Client handle").String()
	showjson     = app.Flag("show-json", "Show JSON output").Bool()
	async        = app.Flag("async", "Asynchronous mode").Bool()
	globoff      = app.Flag("globoff", "Turn off globbing").Bool()
	loglevel     = app.Flag("log-level", "Log Level").Default("warn").String()
	output       = app.Flag("output", "output file").Short('o').String()
	environment  = app.Flag("environment", "Environment").Default("test").String()

	ping    = app.Command("ping", "Ping Operation")
	pingRes = ping.Arg("resource", "Resource").Required().String()

	create       = app.Command("create", "Create Operation")
	createRes    = create.Arg("resource", "Resource").Required().String()
	createParams = create.Arg("params", "additional parameters").StringMap()

	read       = app.Command("read", "Read Operation")
	readRes    = read.Arg("resource", "Resource").Required().String()
	readParams = read.Arg("params", "additional parameters").StringMap()

	update       = app.Command("update", "Update Operation")
	updateRes    = update.Arg("resource", "Resource").Required().String()
	updateParams = update.Arg("params", "additional parameters").StringMap()

	delete       = app.Command("delete", "Delete Operation")
	deleteRes    = delete.Arg("resource", "Resource").Required().String()
	deleteParams = delete.Arg("params", "additional parameters").StringMap()

	search       = app.Command("search", "Search Operation")
	searchRes    = search.Arg("resource", "Resource").Required().String()
	searchParams = search.Arg("params", "additional parameters").StringMap()

	documentation       = app.Command("documentation", "Documentation Operation")
	documentationRes    = documentation.Arg("resource", "Resource").Required().String()
	documentationParams = documentation.Arg("params", "additional parameters").StringMap()
)

// Operation is buzzapi operation operand
type Operation int

const (
	// UNKNOWN is buzzapi unknown operation
	UNKNOWN Operation = iota
	// PING is basic "up" test
	PING Operation = iota
	// READ is buzzapi read operation
	READ Operation = iota
	// SEARCH is buzzapi search operation
	SEARCH Operation = iota
	// UPDATE is buzzapi update operation
	UPDATE Operation = iota
	// CREATE is buzzapi create operation
	CREATE Operation = iota
	// DELETE is buzzapi delete operation
	DELETE Operation = iota
	// DOCUMENTATION is buzzapi documentation operation
	DOCUMENTATION Operation = iota
)

func (o Operation) String() string {
	switch o {
	case PING:
		return "ping"
	case READ:
		return "read"
	case SEARCH:
		return "search"
	case UPDATE:
		return "update"
	case CREATE:
		return "create"
	case DELETE:
		return "delete"
	case DOCUMENTATION:
		return "documentation"
	case UNKNOWN:
		return "unknown"
	}

	return "OOPS"
}

var op Operation
var res string
var params *map[string]string

func initializeArgs() {
	//
	// process command line arguments
	//
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("1.0").Author("Alec Thomas")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case ping.FullCommand():
		op = PING
		res = *pingRes
	case create.FullCommand():
		op = CREATE
		res = *createRes
		params = createParams
	case read.FullCommand():
		op = READ
		res = *readRes
		params = readParams
	case update.FullCommand():
		op = UPDATE
		res = *updateRes
		params = updateParams
	case delete.FullCommand():
		op = DELETE
		res = *deleteRes
		params = deleteParams
	case search.FullCommand():
		op = SEARCH
		res = *searchRes
		params = searchParams
	case documentation.FullCommand():
		op = DOCUMENTATION
		res = *documentationRes
		params = documentationParams
	}
}
