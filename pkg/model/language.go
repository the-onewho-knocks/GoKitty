package model

import (
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
	return languageExt[l]
}

func DetectLanguage(root string) Language {
	langScore := make(map[Language]int)
	var walk func(dir string)
	walk = func(dir string) {
		entries, _ := filepath.Glob(filepath.Join(dir, "*"))
		for _, entry := range entries {
			ext := strings.ToLower(filepath.Ext(entry))
			for lang, exts := range languageExt {
				for _, e := range exts {
					if ext == e {
						langs[lang]++
					}
				}
			}
		}
	}
	walk(root)

	var detected Language
	maxCount := 0
	for lang, count := range langs {
		if count > maxCount {
			maxCount = count
			detected = lang
		}
	}
	return detected
}
