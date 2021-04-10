# goenvargs

## Installation

As a library

```shell
go get github.com/PrashantRaj18198/goenvargs
```

## Usage

From your project

```bash
go build
./your_project FOO=bar ENV2=go\$
```

If your string has special characters like `$` remember to escape it with a `\`
before it.
Note: The default splitter is `=`.

## In your project

Autoload - parse all the vars, I don't care about missing env variables

```go
import (
  _ "github.com/PrashantRaj18198/goenvargs/autoload"
)

```

Show me errors!

```go

import (
  "github.com/PrashantRaj18198/goenvargs"
)

func init() {
  requiredEnvs := []string{"X_ENV1", "X_ENV2"}

  // throws error is any of required Envs are not present
	parsedEnvs, err := goenvargs.LoadEnvVars(requiredEnvs, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(parsedEnvs)
}

```

Use a custom splitter -

```go
import (
  "github.com/PrashantRaj18198/goenvargs"
)

func init() {
  splitter := ":"
	parsedEnvs, _ := goenvargs.LoadEnvVars(nil, &splitter)
	fmt.Print(parsedEnvs)
}

```
