goji_gzip
=========

[goji framework](https://goji.io) gzip middleware


example:

```golang
package main

import (
        "fmt"
        "net/http"
        "github.com/lidashuang/goji_gzip"

        "github.com/zenazn/goji"
)

func main() {

		goji.Use(gzip.GzipHandler)

    goji.Get("/", func(w http.ResponseWriter, r *http.Request){
			fmt.Fprint(w, "helloworld..........")
		})

    goji.Serve()
}
```
