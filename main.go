package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/experimental/gojs"
	"github.com/tetratelabs/wazero/sys"
)

func main() {
	log.SetOutput(os.Stderr)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cacheDir := filepath.Join(os.TempDir(), "wazero", "buildcache")
	cache, err := wazero.NewCompilationCacheWithDir(cacheDir)
	if err != nil {
		log.Fatal("Failed to create build cache:", err)
	}

	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig().
		WithCloseOnContextDone(true).
		WithCompilationCache(cache))

	binPath := "./wasm/" + os.Args[0] + ".wasm"
	bin, err := os.ReadFile(binPath)
	if err != nil {
		log.Fatalf("Failed to read binary %q: %v", binPath, err)
	}
	m, err := r.CompileModule(ctx, bin)
	if err != nil {
		log.Fatal("Failed to compile module:", err)
	}

	gojs.MustInstantiate(ctx, r)

	conf := wazero.NewModuleConfig().
		WithStdin(os.Stdin).
		WithStderr(os.Stderr).
		WithStdout(os.Stdout).
		WithArgs(os.Args...)

	if err := gojs.Run(ctx, r, m, gojs.NewConfig(conf)); err != nil {
		exitErr, ok := err.(*sys.ExitError)
		if !ok {
			log.Fatal("Failed to execute binary:", err)
		}
		if exitErr.ExitCode() != 0 {
			log.Fatal("Failed to execute binary:", err)
		}
	}
}
