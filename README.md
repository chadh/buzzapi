## Buzzapi go command line client

  This is just a simple command line client for interacting with Buzzapi.
  
    usage: gobuzzapi [<flags>] <command> [<args> ...]
    
    Buzzapi Command Line Interface
    
    Flags:
      --help               Show help (also see --help-long and --help-man).
      --timeout=15         Max time to wait (in seconds)
      -v, --verbose        Lots of output
      -q, --quiet          Very little output
      --handle=HANDLE      Client handle
      --client-handle=CLIENT-HANDLE
                          Client handle
      --show-json          Show JSON output
      --async              Asynchronous mode
      --globoff            Turn off globbing
      --log-level="warn"   Log Level
      -o, --output=OUTPUT  output file
      --environment="test"
                        Environment
    
    Commands:
      help [<command>...]
        Show help.
    
      ping <resource>
        Ping Operation
    
      create <resource> [<params>...]
        Create Operation
    
      read <resource> [<params>...]
        Read Operation
    
      update <resource> [<params>...]
        Update Operation
    
      delete <resource> [<params>...]
        Delete Operation
    
      search <resource> [<params>...]
        Search Operation
    
      documentation <resource> [<params>...]
        Documentation Operation

