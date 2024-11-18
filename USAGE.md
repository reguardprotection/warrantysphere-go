<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/reguardprotection/warrantysphere"
	"log"
)

func main() {
	s := warrantysphere.New()

	ctx := context.Background()
	res, err := s.HealthControllerHealthCheck(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->