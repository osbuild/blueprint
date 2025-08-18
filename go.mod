module github.com/osbuild/blueprint

go 1.23.9

toolchain go1.24.2

// temporary fix until images is upgraded out of toml 1.5.1 beta version: https://github.com/osbuild/blueprint/pull/30
replace github.com/osbuild/images v0.161.0 => github.com/osbuild/images v0.176.0

require (
	github.com/BurntSushi/toml v1.5.0
	github.com/coreos/go-semver v0.3.1
	github.com/google/uuid v1.6.0
	github.com/osbuild/images v0.161.0
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
