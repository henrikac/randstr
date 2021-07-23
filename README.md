# randstr

randstr is a module that contains various functions for generating random strings.

## Installation

`go get github.com/henrikac/randstr`

## Usage

### Generate
`randstr.Generate` returns a random string that consists of the characters a-zA-Z0-9.
The default size of the generated string is 16 bytes.

```go
package main

import (
	"fmt"
	"log"

	"github.com/henrikac/randstr"
)

func main() {
	str, err := randstr.Generate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", str)
}

// Example output:
// 4a1Lz3d4yXihvbJX
```

### Base64Encoded
`randstr.Base64Encoded` returns a random base64 encoded string.

```go
package main

import (
	"fmt"
	"log"

	"github.com/henrikac/randstr"
)

func main() {
	str, err := randstr.Base64Encoded()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", str)
}

// Example output:
// M72WQsavclA/wLYfxuyr2Q==
```

## Contributing

1. Fork it (<https://github.com/henrikac/randstr/fork>)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## Contributors

- [Henrik Christensen](https://github.com/henrikac) - creator and maintainer
