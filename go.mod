module github.com/antonioiubatti93/buckle

go 1.19

replace (
	github.com/antonioiubatti93/buckle/pkg/person/v2 => ./pkg/person/v2
	github.com/antonioiubatti93/buckle/sdk/person => ./sdk/person
)

require (
	github.com/antonioiubatti93/buckle/pkg/person/v2 v2.0.0
	github.com/antonioiubatti93/buckle/sdk/person v0.0.0
)
