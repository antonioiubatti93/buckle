# Buckle

Test repo in Go with multi-modules.

## Structure

Set of trivial (nested) modules providing a dummy constructor for a person with name and age:
- `pkg/person`
- `pkg/person/v2`: example of major version with breaking change
- `sdk/person`: exactly as `pkg/person`

## Issue

See https://github.com/antonioiubatti93/buckle/issues/1:
- if the main program is not using `person/sdk`, `gocov` fails with a go mod error
- as soon as the dependency enters the main module the proper way, the coverage can be produced
