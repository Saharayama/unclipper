package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

const (
	excelPath = `C:\Program Files\Microsoft Office\root\Office16\EXCEL.EXE`
	wordPath  = `C:\Program Files\Microsoft Office\root\Office16\WINWORD.EXE`
)

var (
	invalidPathChars = regexp.MustCompile(`"|^ *-|^\t*-|<|>`)
	networkPathRegex = regexp.MustCompile(`^\\+`)
)

func formatPath(p string) string {
	p = invalidPathChars.ReplaceAllString(p, "")
	p = strings.TrimSpace(p)

	var path string
	if strings.Contains(p, ":") && strings.Contains(p, "/") {
		path = strings.ReplaceAll(p, "/", "\\")
	} else if strings.Contains(p, "/") {
		p = strings.ReplaceAll(p, "%20", " ")
		p = strings.ReplaceAll(p, "%23", "#")
		p = strings.ReplaceAll(p, "%5b", "[")
		p = strings.ReplaceAll(p, "%5d", "]")
		p = strings.ReplaceAll(p, "%5e", "^")
		p = strings.ReplaceAll(p, "%60", "`")
		p = strings.ReplaceAll(p, "%7b", "{")
		p = strings.ReplaceAll(p, "%7d", "}")
		p = strings.ReplaceAll(p, "%25", "%")
		path = "\\\\" + strings.ReplaceAll(p, "/", "\\")
	} else if strings.HasPrefix(p, "\\\\") {
		p = networkPathRegex.ReplaceAllString(p, "")
		path = "\\\\" + p
	} else {
		path = p
	}

	return path
}

func openPath(pathFormatted string) {
	if strings.HasPrefix(pathFormatted, "!--") {
		return
	}

	pathSuffix := filepath.Ext(pathFormatted)
	var cmd *exec.Cmd

	if !strings.HasPrefix(pathFormatted, "http") {
		if strings.HasPrefix(pathSuffix, ".xl") {
			if _, err := os.Stat(excelPath); err == nil {
				cmd = exec.Command(excelPath, "/x", pathFormatted)
			}
		} else if strings.HasPrefix(pathSuffix, ".do") {
			if _, err := os.Stat(wordPath); err == nil {
				cmd = exec.Command(wordPath, "/w", pathFormatted)
			}
		}
	}

	if cmd == nil {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", pathFormatted).Start()
	} else {
		if _, err := os.Stat(pathFormatted); err != nil {
			return
		}
		cmd.Start()
	}
}

func main() {
	cb, err := clipboard.ReadAll()
	if err != nil {
		return
	}

	lines := strings.SplitSeq(cb, "\n")
	for line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		openPath(formatPath(line))
	}
}
