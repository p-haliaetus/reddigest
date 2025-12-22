package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
)

const maxDays = 30

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	files, err := filepath.Glob("pre-docs/*.md")
	if err != nil {
		return err
	}

	sort.Strings(files)
	if len(files) > maxDays {
		files = files[len(files)-maxDays:]
	}

	digests := make(map[string]string, len(files))
	for _, path := range files {
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		html := markdown.ToHTML(b, nil, nil)
		date := strings.TrimSuffix(filepath.Base(path), ".md")
		digests[date] = string(html)
	}

	data, err := json.MarshalIndent(digests, "    ", "  ")
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		return err
	}

	out, err := os.Create("pre-docs/index.html")
	if err != nil {
		return err
	}
	defer out.Close()

	return tmpl.Execute(out, map[string]string{
		"DigestsJSON": string(data),
	})
}
