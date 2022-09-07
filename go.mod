module github.com/antonioiubatti93/buckle

go 1.19

replace (
	github.com/antonioiubatti93/buckle/pkg/person => ./pkg/person
	github.com/antonioiubatti93/buckle/pkg/person/v2 => ./pkg/person/v2
)

require (
	github.com/antonioiubatti93/buckle/pkg/person v0.0.0-00010101000000-000000000000
	github.com/antonioiubatti93/buckle/pkg/person/v2 v2.0.0-00010101000000-000000000000
)
