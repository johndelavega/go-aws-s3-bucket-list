module list

require (
	// v1.16.11 appears to be the minimum version as: 'go mod tidy' upgrades from lower version such as v1.15.64
	github.com/aws/aws-sdk-go v1.16.11
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/net v0.0.0-20190301231341-16b79f2e4e95 // indirect
)

replace github.com/aws/aws-sdk-go v1.16.11 => /home/johndlvg/egd/code/go/jd/aws-sdk-go-v1.15.64 // OK
