## Buzzapi go command line client

This is just a simple command line client for interacting with Buzzapi.

### Usage

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

## Configuration

The program requires a TOML formatted `config.toml` file like this:

```
[environments]
  [environments.common]
  api_user = 'YOUR_ACCT_HERE'
  api_passwd = 'YOUR_PW_HERE'

  [environments.prod]
  api_url = 'https://api.gatech.edu'
  api_user = 'YOUR_PROD_ACCT_HERE'
  api_passwd = 'YOUR_PROD_PW_HERE'

  [environments.test]
  api_url = 'https://test.api.gatech.edu'
```

The common environment settings serve as defaults.  Anything set in that
environment will be used in all other environments unless overridden (e.g. the
`api_user` and `api_passwd` settings in the prod environment)

Note that you specify the environment with just the string after the period.  So to use the settings in the `environments.prod` section, you specify `--environment prod` on the command line.

`config.toml` should either in the local directory or at `~/.config.toml`

## Example

```
$ buzzapi search central.email.aliases alias_regex=foo@dept.gatech.edu --environment prod
[
  {
    "date_to_expire": null,
    "destination": "foo@gatech.edu",
    "email_alias": "foo@dept.gatech.edu",
    "gtid": "08675309",
    "mage_person_index": null,
    "mage_role_index": 4242,
    "required_because": null,
    "tracks_primary_email": false
  }
]
```
