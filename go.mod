module github.com/antonioiubatti93/buckle

go 1.19

replace (
	github.com/antonioiubatti93/buckle/interestrate => ./interestrate
	github.com/antonioiubatti93/buckle/termstructure => ./termstructure
)

require (
	github.com/antonioiubatti93/buckle/interestrate v0.0.0
	github.com/antonioiubatti93/buckle/termstructure v0.0.0
)

require (
	github.com/edgelaboratories/interpolator v0.3.0 // indirect
	golang.org/x/exp v0.0.0-20221217163422-3c43f8badb15 // indirect
)
