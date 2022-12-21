# Buckle

Test repo in Go with multi-modules.

## Use case

Zero-coupon bond pricing, given notional, maturity and interest rate.

## Components

### Term structure

Defines evolving functions of time described in different ways, e.g.:
- constant (flat)
- functional/parametric
- by tenor-value data point pairs

See the **module** [termstructure](./termstructure/go.mod) for the details.

### Interest rate

Defines interest rates based on any term structure to compute:
- short rates
- curve shifts
- discount factors

as functions of time.

See the **module** [interestrate](./interestrate/go.mod) for the details.

### Main

Just an example.

## Design

Inspired by [SOLID principles](https://en.wikipedia.org/wiki/SOLID) and [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) concepts, e.g.:
- high-level policies should not depend on low-level details
- dependencies point inwards, i.e. mechanisms depend on policies
- business rules/entities are the most stable and abstract components
- the main is the outest and latest detail, glue holding it all together

Notice that:
- no sub-module imports anything from outer circles (check out the go.mod)
- only the main (lowest-level detail) imports the sub-modules
