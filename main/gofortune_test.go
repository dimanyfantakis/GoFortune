package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const expectedFortunesFilePath = "/usr/share/games/fortunes"

func TestFindFortuneFilesPath(t *testing.T) {
	actualFortuneFilesPath := findFortuneFilesPath()
	if actualFortuneFilesPath != expectedFortunesFilePath {
		t.Fatalf("The fortune's file path should be, %s", expectedFortunesFilePath)
	}
}

func TestCollectFortuneFiles(t *testing.T) {
	actualFortuneFiles := collectFortuneFiles(expectedFortunesFilePath)
	for i := range actualFortuneFiles {
		x, x1 := os.Stat(actualFortuneFiles[i])
		if x.IsDir() || x1 != nil || filepath.Ext(actualFortuneFiles[i]) == ".dat" || filepath.Ext(actualFortuneFiles[0]) == ".u8" {
			t.Fatalf("Invalid fortune file")
		}
	}
}

func TestReadFortunes(t *testing.T) {
	x := filepath.Join(expectedFortunesFilePath, "fortunes")
	readfile, err := os.Open(x)
	if err != nil {
		panic(err)
	}

	fortunes := readFortunes(readfile)

	if len(fortunes) == 0 {
		t.Fatalf("Failed to read fortunes from file")
	}

	for _, f := range fortunes {
		if strings.Contains(f, "%") {
			t.Fatalf("Invalid fortune")
		}

		if len(f) == 0 {
			t.Fatalf("Fortune can't be empty")
		}
	}
}

func TestPrintRandomFortune(t *testing.T) {
	x := filepath.Join(expectedFortunesFilePath, "fortunes")
	readfile, err := os.Open(x)
	if err != nil {
		panic(err)
	}

	fortuneFiles := readFortunes(readfile)

	stdOutR := os.Stdout
	e, w, _ := os.Pipe()
	os.Stdout = w

	printRandomFortune(fortuneFiles)

	w.Close()
	stdOut, _ := io.ReadAll(e)
	os.Stdout = stdOutR

	if len(stdOut) == 0 {
		t.Fatalf("Fortune should be printed to stdout")
	}

}
