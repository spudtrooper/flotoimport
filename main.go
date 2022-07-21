package main

import (
	"flag"
	"log"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spudtrooper/floto/floto"
)

var (
	dir    = flag.String("dir", "~/Desktop/raw", "The input directory")
	outdir = flag.String("outdir", "~/Desktop/floto", "Output dir")
)

func realMain() error {
	if *dir == "" {
		return errors.Errorf("--dir required")
	}
	fs, err := filepath.Glob(path.Join(*dir, "*.jpg"))
	if err != nil {
		return err
	}
	imprt := floto.MakeImporter()
	for _, f := range fs {
		if err := imprt.ImportImage(*outdir, f, "", ""); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if err := realMain(); err != nil {
		log.Fatalf(err.Error())
	}
}
