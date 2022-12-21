module github.com/antonioiubatti93/buckle

go 1.19

replace (
	github.com/antonioiubatti93/buckle/interestrate => ./interestrate
	github.com/antonioiubatti93/buckle/termstructure => ./termstructure
)

require (
	github.com/antonioiubatti93/buckle/interestrate v0.0.0
	github.com/antonioiubatti93/buckle/termstructure v0.0.0
	github.com/stretchr/testify v1.8.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/edgelaboratories/interpolator v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20221217163422-3c43f8badb15 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
