package model

import (
	"os"
	"path/filepath"
	"strings"
)

type Language int

const (
	LanguageUnknown Language = iota
	LanguageGo
	LanguagePython
	LanguageJavaScript
	LanguageTypeScript
	LanguageJava
)

var languageExt = map[Language][]string{
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
	if n, ok := languageNames[l]; ok {
		return n
	}
	return "Unknown"
}

func (l Language) Extensions() []string {
	return languageExt[l]
}

const minFileThreshold = 3

func DetectLanguages(root string) []Language {
	counts := make(map[Language]int)
	total := 0

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			if info != nil && info.IsDir() && strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
				return filepath.SkipDir
			}
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		for lang, exts := range languageExt {
			for _, e := range exts {
				if ext == e {
					counts[lang]++
					total++
				}
			}
		}
		return nil
	})

	var result []Language
	for lang, count := range counts {
		if count >= minFileThreshold || (total > 0 && float64(count)/float64(total) >= 0.05) {
			result = append(result, lang)
		}
	}
	if len(result) == 0 {
		for lang, count := range counts {
			if count > 0 {
				result = append(result, lang)
			}
		}
	}
	return result
}

func DetectLanguage(root string) Language {
	counts := make(map[Language]int)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			if info != nil && info.IsDir() && strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
				return filepath.SkipDir
			}
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		for lang, exts := range languageExt {
			for _, e := range exts {
				if ext == e {
					counts[lang]++
				}
			}
		}
		return nil
	})
	var best Language
	maxCount := 0
	for lang, count := range counts {
		if count > maxCount {
			maxCount = count
			best = lang
		}
	}
	return best
}
