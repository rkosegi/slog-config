# Go structured log config

This repository contains code to simplify configuration of [`slog`](https://pkg.go.dev/log/slog).

```go
package main

import xlog "github.com/rkosegi/slog-config"

func main() {
	lc := xlog.MustNew("debug", "json")
	lc.Logger().Info("Hello World")
}
```

Output
```
{"ts":"2024-11-27T05:45:19.555Z","level":"info","caller":"main.go(main.main):7","msg":"Hello World"}
```
