module github.com/antonioiubatti93/buckle

go 1.20

require (
	github.com/antonioiubatti93/buckle/people v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.3
)

// Try to comment out and see what happens...
replace github.com/antonioiubatti93/buckle/people => ./people

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
