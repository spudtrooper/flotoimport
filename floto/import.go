package floto

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"github.com/spudtrooper/goutil/io"
)

func readImage(f string) (image.Image, error) {
	fd, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	img, err := jpeg.Decode(fd)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func writeImage(img image.Image, f string) error {
	fd, err := os.Create(f)
	if err != nil {
		return err
	}
	defer fd.Close()

	return jpeg.Encode(fd, img, &jpeg.Options{
		Quality: 100,
	})
}

func cropImage(img image.Image, new_width, new_height uint) (image.Image, error) {
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	var diff, left, right, top, bottom int
	if width > height {
		diff = width - height
		left = int(diff / 2)
		right = width - int(diff/2)
		top = 0
		bottom = height
	} else {
		diff = height - width
		top = int(diff / 2)
		bottom = height - int(diff/2)
		left = 0
		right = width
	}
	log.Printf("left=%d top=%d right=%d bottom=%d width=%d height=%d new width=%d new height=%d",
		left, top, right, bottom, width, height, (right - left), (bottom - top))
	return cutter.Crop(img, cutter.Config{
		Width:   width,
		Height:  height,
		Anchor:  image.Point{top, left},
		Options: cutter.Copy,
	})
}

func resizeImage(infile, outFile string, width, height uint) error {
	log.Printf("Resizing %s to %dx%d to %s", infile, width, height, outFile)
	img, err := readImage(infile)
	if err != nil {
		return err
	}
	croppedImg, err := cropImage(img, width, height)
	if err != nil {
		return err
	}
	resizedImg := resize.Resize(width, height, croppedImg, resize.Lanczos3)
	if writeImage(resizedImg, outFile); err != nil {
		return err
	}
	return nil
}

func writeFile(outFile, s string) error {
	return ioutil.WriteFile(outFile, []byte(s), 0755)
}

type Importer interface {
	ImportImage(outDir, infile, description, location string) error
}

type importer struct {
	start    time.Time
	nextSecs int64
}

func MakeImporter() Importer {
	now := time.Now().Local()
	return &importer{
		start:    now,
		nextSecs: now.UnixMicro() / 1000 / 1000,
	}
}

func (i *importer) ImportImage(outDir, infile, description, location string) error {
	baseDir := outDir
	if baseDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		baseDir = path.Join(home, "Desktop", "floto")
	}
	if _, err := io.MkdirAll(baseDir); err != nil {
		return err
	}
	sec := i.start.Format("20060102") + fmt.Sprintf("%d", i.nextSecs)
	i.nextSecs++
	log.Printf("Using infile    : %s", infile)
	log.Printf("Using base dir  : %s", baseDir)
	log.Printf("Using sec       : %s", sec)
	outFile := path.Join(baseDir, sec+".jpg")
	out_description := path.Join(baseDir, sec+".des")
	out_location := path.Join(baseDir, sec+".loc")
	outFile_med := path.Join(baseDir, sec+"-med.jpg")
	var width uint = 150
	var height uint = 150
	width_med := width * 4
	height_med := height * 4
	if err := resizeImage(infile, outFile, width, height); err != nil {
		return err
	}
	if err := resizeImage(infile, outFile_med, width_med, height_med); err != nil {
		return err
	}

	if err := writeFile(out_description, description); err != nil {
		return err
	}
	if err := writeFile(out_location, location); err != nil {
		return err
	}

	if err := io.Copy(outFile, path.Join(baseDir, "new.jpg")); err != nil {
		return err
	}
	return nil
}
