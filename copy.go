package duck

import (
	"bytes"
	"encoding/json"
	"reflect"
)

// Copy performs a deep copy of the given interface. The copy makes sure that no pointers/maps are shared
//	between the origginal data and new points.
// WARNING: This is not a real deep copy. It totally cheats at the moment. Do not use when you need
//	to preserve structs in output.
// WARNING 2: Converts int to float. Ultimately, this function is not yet good for normal use. It is just
//	barely useful for our own internal usage. Feel free to fix it.
func Copy(copyfrom interface{}) (interface{}, error) {
	//	Here's the idea: I need to have this working NOW, so I am just saying f it,
	//	and let gob encoding/decoding do the hard work. This is relatively inefficient,
	//	but I don't have the time ATM to make it better, since I have way too much on my head.
	//	TODO: Make this not use gob

	v := reflect.ValueOf(copyfrom)

	//First check that the value isn't something simple like int/bool/etc

	// Create the "output" variable by copying the interface type of input.
	copyto := reflect.Zero(v.Type()).Interface()

	// Copied from: http://stackoverflow.com/a/28579297

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	dec := json.NewDecoder(&buf)

	err := enc.Encode(copyfrom)
	if err != nil {
		return nil, err
	}

	err = dec.Decode(&copyto)

	return copyto, err
}
