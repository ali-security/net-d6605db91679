// Command create_mod builds a Go module zip in the canonical proxy.golang.org
// layout from a checkout directory, using golang.org/x/mod/zip.CreateFromDir.
//
// Usage: create_mod <module-path> <version> <source-dir> <output-zip>
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <source-dir> <output-zip>")
		os.Exit(2)
	}
	modPath, version, srcDir, outPath := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "create output dir: %v\n", err)
		os.Exit(1)
	}

	out, err := os.Create(outPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create output zip: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(out, m, srcDir); err != nil {
		fmt.Fprintf(os.Stderr, "create module zip: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("wrote %s for %s@%s\n", outPath, modPath, version)
}
