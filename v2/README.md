#### SMARTY DISCLAIMER: Subject to the terms of the associated license agreement, this software is freely available for your use. This software is FREE, AS IN PUPPIES, and is a gift. Enjoy your new responsibility. This means that while we may consider enhancement requests, we may or may not choose to entertain requests at our sole and absolute discretion.

[![GoDoc](https://godoc.org/github.com/smarty/gunit/v2?status.svg)](http://godoc.org/github.com/smarty/gunit/v2)

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
