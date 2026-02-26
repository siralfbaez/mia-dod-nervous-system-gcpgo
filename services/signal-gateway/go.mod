module github.com/siralfbaez/mia-dod-nervous-system-gcpgo/services/signal-gateway

go 1.25.5

require (
	github.com/siralfbaez/mia-dod-nervous-system-gcpgo/pkg/resilience v0.0.0
	github.com/sony/gobreaker v0.5.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// This tells Go to use the local version of your package during development
replace github.com/siralfbaez/mia-dod-nervous-system-gcpgo/pkg/resilience => ../../pkg/resilience
