package testdata

const ok = 1

const a = iota // want "iota usage found"

const b = iota << 1 // want "iota usage found"

const c = 1 + iota // want "iota usage found"

type MyInt int

const d MyInt = iota // want "iota usage found"

const (
	e = iota // want "iota usage found"
	f

	g = iota // want "iota usage found"
)
