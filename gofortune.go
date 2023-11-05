package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	fortuneFilesPath := findFortuneFilesPath()
	fortuneFiles := collectFortuneFiles(fortuneFilesPath)
	readfile := readRandomFortuneFile(fortuneFiles)
	defer readfile.Close()

	fortunes := readFortunes(readfile)
	printRandomFortune(fortunes)
}

func findFortuneFilesPath() string {
	fortuneCommand := exec.Command("fortune", "-f")
	stdError, err := fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = fortuneCommand.Start()
	if err != nil {
		panic(err)
	}

	stdoutScanner := bufio.NewScanner(stdError)
	stdoutScanner.Scan()
	output := stdoutScanner.Text()
	fortuneFilesPath := strings.Split(output, " ")[1:][0]
	return fortuneFilesPath
}

func collectFortuneFiles(fortuneFilesPath string) []string {
	fortuneFiles := make([]string, 0)
	filepath.WalkDir(fortuneFilesPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if filepath.Ext(path) == ".dat" || filepath.Ext(path) == ".u8" {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		fortuneFiles = append(fortuneFiles, path)
		return nil
	})
	return fortuneFiles
}

func readRandomFortuneFile(fortuneFiles []string) *os.File {

	random := rand.Intn(len(fortuneFiles))
	randomFortuneFile := fortuneFiles[random]

	readfile, err := os.Open(randomFortuneFile)
	if err != nil {
		panic(err)
	}

	return readfile
}

func readFortunes(readfile *os.File) []string {
	const (
		breakingCharacter = "%"
	)

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanRunes)

	fortunes := make([]string, 0)
	sb := strings.Builder{}
	for fileScanner.Scan() {
		char := fileScanner.Text()
		if char == breakingCharacter {
			if sb.Len() != 0 {
				fortunes = append(fortunes, sb.String())
				sb.Reset()
			}
			continue
		}
		sb.WriteString(char)
	}
	return fortunes
}

func printRandomFortune(fortunes []string) {
	random := rand.Intn(len(fortunes))
	fortune := fortunes[random]
	fmt.Println(fortune)
}
