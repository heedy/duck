[![Build Status](https://travis-ci.org/connectordb/duck.svg)](https://travis-ci.org/connectordb/duck)[![Coverage Status](https://coveralls.io/repos/connectordb/duck/badge.svg?branch=master&service=github)](https://coveralls.io/github/connectordb/duck?branch=master)[![GoDoc](https://godoc.org/github.com/connectordb/duck?status.svg)](http://godoc.org/github.com/connectordb/duck)
# Duck
Sometimes in golang a function returns an interface (or perhaps something was marshalled into an interface). Duck allows you to manipulate and convert from interfaces in a very simple manner.

## Type Conversions

The most basic actions are conversions between the standard built-in types.

```go
//true, true
b,ok := duck.Bool(1.0)
//false, true
b,ok = duck.Bool(0.0)

//"1337", true
s, ok := duck.String(1337)
//"true", true
s, ok := duck.String(true)

//56, true
i, ok := duck.Int("56.0")

//1.0, true
f, ok := duck.Float(true)

var v interface{}

v = 6.0
vptr := &v	//duck follows pointers

//"6.0",true
s, ok = duck.String(vptr)
//6,true
i,ok = duck.Int(vptr)
//6.0,true
f,ok = duck.Float(vptr)
//true,true
b,ok = duck.Bool(vptr)

```

## Comparisons

Duck allows comparing interfaces with similar conversion semantics as type conversions.

```go
//true, true ("34.5" < 35)
out, ok := duck.Lt("34.5",35)
//false, true
out, ok = duck.Lt("34.5",34)

//true,true
out, ok = duck.Eq("1.0", true)

//false,true - this one is unusual, true is defined as 1 (as in c)
out, ok = duck.Eq("2.0", true)

//false,true - false is defined as 0
out, ok = duck.Eq("2.0", false)

//true, true (34.6 >= 34)
out,ok = duck.Gte(34.6, 34)

if duck.Cmp(45,45.32)==duck.LessThan {
	//true!
}

if duck.Cmp(45," LOL ")==duck.CantCompare {
	//true!
}

if duck.Cmp(45,45.0)==duck.Equals {
	//true!
}

```

## Maps and Structs

Under construction...
