package model

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Language int

const (
	LanguageGo Language = iota + 1
	LanguagePython
	LanguageJavaScript
	LanguageTypeScript
	LanguageJava
)

var languageExtensions = map[Language][]string{
	LanguageGo:         {".go"},
	LanguagePython:     {".py"},
	LanguageJavaScript: {".js", ".jsx", ".mjs", ".cjs"},
	LanguageTypeScript: {".ts", ".tsx", ".mts", ".cts"},
	LanguageJava:       {".java", ".kt", ".groovy"},
}

var languageNames = map[Language]string{
	LanguageGo:         "Go",
	LanguagePython:     "Python",
	LanguageJavaScript: "JavaScript",
	LanguageTypeScript: "TypeScript",
	LanguageJava:       "Java",
}

func (l Language) String() string {
	if name, ok := languageNames[l]; ok {
		return name
	}
	return "Unknown"
}

func (l Language) Extensions() []string {
	return languageExtensions[l]
}

func DetectLanguage(root string) Language {
	langScore := make(map[Language]int)

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))

		for lang, exts := range languageExtensions {
			for _, e := range exts {
				if ext == e {
					langScore[lang]++
					break
				}
			}
		}

		return nil
	})

	var detected Language
	maxCount := 0

	for lang, count := range langScore {
		if count > maxCount {
			maxCount = count
			detected = lang
		}
	}

	return detected
}