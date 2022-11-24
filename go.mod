module github.com/antonioiubatti93/buckle

go 1.19

replace (
	github.com/antonioiubatti93/buckle/pkg/person/v2 => ./pkg/person/v2
	github.com/antonioiubatti93/buckle/sdk/person => ./sdk/person
)

require (
	github.com/antonioiubatti93/buckle/pkg/person v0.0.0
	github.com/antonioiubatti93/buckle/pkg/person/v2 v2.0.0
	github.com/antonioiubatti93/buckle/sdk/person v0.0.1
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
