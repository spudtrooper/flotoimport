package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spudtrooper/floto/floto"
)

var (
	dir    = flag.String("dir", "", "The input directory, defaults to ~/Desktop/raw")
	outdir = flag.String("outdir", "", "Output dir, defaults to ~/Desktop/floto")
)

func realMain() error {
	dir := *dir
	outdir := *outdir
	if dir == "" || outdir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return errors.Errorf("getting home dir: %v", err)
		}
		log.Printf("home: %v", home)
		if dir == "" {
			dir = path.Join(home, "Desktop", "raw")
		}
		if outdir == "" {
			outdir = path.Join(home, "Desktop", "floto")
		}
	}

	fs, err := filepath.Glob(path.Join(dir, "*.jpg"))
	if err != nil {
		return err
	}

	imprt := floto.MakeImporter()
	for _, f := range fs {
		log.Printf("f:%v", f)
		if err := imprt.ImportImage(outdir, f, "", ""); err != nil {
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
