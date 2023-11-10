#### SMARTY DISCLAIMER: Subject to the terms of the associated license agreement, this software is freely available for your use. This software is FREE, AS IN PUPPIES, and is a gift. Enjoy your new responsibility. This means that while we may consider enhancement requests, we may or may not choose to entertain requests at our sole and absolute discretion.

[![Build Status](https://travis-ci.org/smarty/gunit.svg?branch=master)](https://travis-ci.org/smarty/gunit)
[![Code Coverage](https://codecov.io/gh/smarty/gunit/branch/master/graph/badge.svg)](https://codecov.io/gh/smarty/gunit)
[![Go Report Card](https://goreportcard.com/badge/github.com/smarty/gunit)](https://goreportcard.com/report/github.com/smarty/gunit)
[![GoDoc](https://godoc.org/github.com/smarty/gunit?status.svg)](http://godoc.org/github.com/smarty/gunit/v2)

# gunit (v2)

## Installation

```
$ go get github.com/smarty/gunit/v2
```

## Usage

For users of JetBrains IDEs, here's LiveTemplate you can use for generating the scaffolding for a new fixture:

- Abbreviation: `fixture-v2`
- Description: `Generate gunit Fixture boilerplate`
- Template Text:

```
func Test$NAME$Fixture(t *testing.T) {
    gunit.Run(&$NAME$Fixture{T: gunit.New(t)})
}

type $NAME$Fixture struct {
    *gunit.Fixture
}

func (this *$NAME$Fixture) Setup() {
}

func (this *$NAME$Fixture) Test$END$() {
}

```

**NOTE:** _Be sure to specify that this LiveTemplate is applicable in Go files._
