package autoload

import "github.com/PrashantRaj18198/goenvargs"

func init() {
	goenvargs.LoadEnvVars(nil, nil)	
}