# valf

[![GoDoc](https://godoc.org/github.com/pamburus/valf?status.svg)](https://godoc.org/github.com/pamburus/valf)
[![Build Status](https://travis-ci.org/pamburus/valf.svg?branch=master)](https://travis-ci.org/pamburus/valf)
[![Go Report Status](https://goreportcard.com/badge/github.com/pamburus/valf)](https://goreportcard.com/report/github.com/pamburus/valf)
[![Coverage Status](https://coveralls.io/repos/github/pamburus/valf/badge.svg?branch=master&service=github)](https://coveralls.io/github/pamburus/valf?branch=master)

This package provides means for dealing with typified values with snapshotting capabilities uniformly and a way to get its type and value back using visitor pattern.

This package was originally a part of github.com/ssgreg/logf package.

It has become a separate package to get richer flexibility and make this technology available for any other purpose which may not be related to logging.

## Example

The following example creates a new `valf` value and gets its type back using visitor pattern.

```go
package main

import (
	"fmt"

	"github.com/pamburus/valf"
)

type testVisitor struct {
	valf.IgnoringVisitor
}

func (v testVisitor) VisitString(value string) {
	fmt.Printf("string: %#v\n", value)
}

func (v testVisitor) VisitBytes(value []byte) {
	fmt.Printf("bytes: %q\n", string(value))
}

func main() {
	s := valf.String("some string value")
	s.AcceptVisitor(testVisitor{})

	bv := []byte("some bytes value")
	b := valf.Bytes(bv)
	b.AcceptVisitor(testVisitor{})

	bv[1] = 'a'
	b.AcceptVisitor(testVisitor{})

	bs := b.Snapshot()
	bv[1] = 'o'
	bs.AcceptVisitor(testVisitor{})
}
```

The output is the following:

```
string: "some string value"
bytes: "some bytes value"
bytes: "same bytes value"
bytes: "same bytes value"
```

