[![Build Status](https://travis-ci.org/connectordb/duck.svg)](https://travis-ci.org/connectordb/duck)[![Coverage Status](https://coveralls.io/repos/connectordb/duck/badge.svg?branch=master&service=github)](https://coveralls.io/github/connectordb/duck?branch=master)[![GoDoc](https://godoc.org/github.com/connectordb/duck?status.svg)](http://godoc.org/github.com/connectordb/duck)
# Duck
Sometimes in golang a function returns an interface (or perhaps something was marshalled into an interface). Duck allows you to manipulate and convert from interfaces in a very simple manner. Duck is *very* eager to fit the inputs to the wanted format.

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

var vptr interface{}

v := 6.0
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

Duck allows comparing interfaces with similar conversions as the type functions.

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

## Objects

Duck gets very liberal with its duck-typing for objects. The main function here is `Get` which gets an element from the object. Since duck just *loves* duck-typing, as long as the requested element
can be found in the given object, it is extracted - no matter if the object is a map, struct or array/slice.

```go
floatarray := []float32{0.0,1.0,2.0,3.0,4.0}

// 1.0, true
e,ok := duck.Get(floatarray,1)
e,ok = duck.Get(floatarray,1.0)
e,ok = duck.Get(floatarray,true)
e,ok = duck.Get(floatarray,"1.0")

//Python-style indexing!
//4.0,true
e, ok = duck.Get(floatarray,-1.0)

//Currently, only map[string] is supported - panics on other map types!
smap := map[string]string{"foo": "bar", "true": "wow", "1": "one"}

//"bar", true
s, ok := duck.Get(smap, "foo")
//"wow", true
s, ok = duck.Get(smap,true)
// "one", true
s, ok = duck.Get(smap,1)	//Note that string conversions ALWAYS convert 1.000 -> 1

st := struct {
	MyElement string
	SecondElement int
}{"woo",1337}

//"woo", true
sv, ok := duck.Get(st,"MyElement")
//nil, false
sv, ok = duck.Get(st,"Nonexisting")

```

#### Duck-Tags

Sometimes you want structs to be recognized by a special tag in Get. The `duck` tag allows you to do that.

```go
val := struct{
	A1 string `duck:"lol"`
	A2 string `duck:"-"`
}{"foo","bar"}

//nil,false
v,ok := duck.Get(val,"A1")
v,ok = duck.Get(val,"A2")

//"foo", true
v, ok = duck.Get(val,"lol")
```


## Set

Initial support for setting values is also available.

```go
var integer int
//true
ok := duck.Set(&integer,"54")
if (intger==54) {
	//true!
}

var mystring string
//true
ok = duck.Set(&mystring,13.0)
if (mystring=="13") {
	//true!
}

var iface interface{}

//true
ok = duck.Set(&iface, true)
//true!
_,ok = iface.(bool)

//Currently, only map[string]interface{} is supported for setting values
// reflect makes it very difficult to set map values. Structs and arrays work fine.
mymap := map[string]interface{}{"foo":"bar"}
ok = duck.Set(&mymap,1337,"foo")
//mymap["foo"]=1337 now

//Arbitrary object depth is supported
mysuperobject := struct{
	A1 []interface{}
	A2 string
}{
	A1: []interface{}{"hello","world"},
}
//world
duck.Get(mysuperobject,"A1",1)

//sets "world" to "not world anymore"
ok = duck.Set(&m,"not world anymore","A1",1)


```
