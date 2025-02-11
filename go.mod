module github.com/mikeschinkel/prefsctl

go 1.23.4

require (
	github.com/hashicorp/go-version v1.7.0
	github.com/micromdm/plist v0.2.1
	github.com/mikeschinkel/go-diffator v0.0.0-20240114054559-18c353d6899c
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.6
	golang.org/x/text v0.21.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mikeschinkel/go-lib v0.0.0-20240730075321-74d570fe6a1a // indirect
)

replace github.com/micromdm/plist v0.2.1 => github.com/mikeschinkel/go-plist v0.0.0-20250129020546-50b570acb5e1
