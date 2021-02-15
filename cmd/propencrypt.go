package main

import (
	"gaelgirodon.fr/propencrypt/internal/app"
	"os"
)

// main is the application entrypoint.
func main() {
	os.Exit(app.Run())
}
