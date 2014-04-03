# Duplicate Detector (Golang)

[![Build Status](https://travis-ci.org/quipo/duplicatedetector.png?branch=master)](https://travis-ci.org/quipo/duplicatedetector) 
[![GoDoc](https://godoc.org/github.com/quipo/duplicatedetector?status.png)](http://godoc.org/github.com/quipo/duplicatedetector)

## Introduction

Memcached-based duplicate detector. 

Given a message id/hash, it can check whether it was previously seen by the system.


## Installation

```sh
# load dependency
$ go get github.com/bradfix/gomemcache/memcache

# install package
$ go get github.com/quipo/duplicatedetector
```

## Sample usage

```go
package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/quipo/duplicatedetector"
)

func main() {

	mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211")

	prefix := "myapp."
	dupedetector := duplicatedetector.NewChecker(mc, prefix, 3600) // cache for 1 hour

	for i := 0; i < 3; i++ {
		dupe, err := dupedetector.IsDuplicate("abc")
		if nil != err {
			fmt.Println("**** ERROR from memcache: ", err)
		}
		if dupe {
			fmt.Println("Entry is a duplicate: skipping")
			continue
		}
		fmt.Println("First time this entry is seen")
	}
}

```

## Author

Lorenzo Alberton

* Web: [http://alberton.info](http://alberton.info)
* Twitter: [@lorenzoalberton](https://twitter.com/lorenzoalberton)
* Linkedin: [/in/lorenzoalberton](https://www.linkedin.com/in/lorenzoalberton)


## Copyright

See [LICENSE](LICENSE) document
