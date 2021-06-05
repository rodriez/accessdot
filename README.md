# accessdot
A simple library to access an element using string expressions

## Installation

```bash
go get github.com/rodriez/accessdot
```

## Usage
[Go play!!! ](https://play.golang.org/p/BNhVTclefQ9)

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/rodriez/accessdot"
)

func main() {
	expression := "call.agent.id"
	text := `{"call":{"agent":{"id":"as23df"}}}`

	var source map[string]interface{}
	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	fmt.Println(result)
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

[Hey 👋 buy me a beer! ](https://www.buymeacoffee.com/rodriez)
