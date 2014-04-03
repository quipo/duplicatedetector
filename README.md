# Duplicate Detector (Golang)

## Introduction

Memcached-based duplicate detector. Given a message id/hash, it can check whether it was previously seen by the system.



## Installation

    go get github.com/bradfix/gomemcache/memcache
    go get github.com/quipo/duplicatedetector


## Configuration


## Sample usage

```go
package main

import (
	"github.com/quipo/duplicatedetector"
)

func main() {

	mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212")

	prefix := "myapp."
	dupedetector := duplicatedetector.NewChecker(mc, prefix, 3600) // cache for 1 hour

	for i:=0; i<3; i++ {
		if dupedetector.IsDuplicate("abc") {
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
