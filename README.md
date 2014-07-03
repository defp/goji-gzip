goji_gzip
=========

[goji framework](https://goji.io) gzip middleware


example:

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/lidashuang/goji_gzip"

	"github.com/zenazn/goji"
)


func main() {

	// use gzip handle for every request
	goji.Use(gzip.GzipHandler)

	goji.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "helloworld..........")
	})

	goji.Serve()
}
```

TODO: 

* add some test
