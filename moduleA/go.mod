module github.com/example/workspace-module-a

go 1.23.0

// godebug here is IGNORED in workspace mode - go.work godebug takes precedence
godebug panicnil=1

require (
	github.com/samber/lo v1.47.0
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/text v0.16.0 // indirect
)

