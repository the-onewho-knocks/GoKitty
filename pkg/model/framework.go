package model

import (
	"os"
	"path/filepath"
	"strings"
)

type Framework int

const (
	FrameworkNone Framework = iota
	FrameworkGin
	FrameworkEcho
	FrameworkFiber
	FrameworkChi
	FrameworkDjango
	FrameworkFlask
	FrameworkFastAPI
	FrameworkExpress
	FrameworkNestJS
	FrameworkNextJS
	FrameworkReact
	FrameworkVue
	FrameworkAngular
	FrameworkSpring
	FrameworkSpringBoot
	FrameworkJavaEE
)

var frameworkNames = map[Framework]string{
	FrameworkNone:       "None",
	FrameworkGin:        "Gin",
	FrameworkEcho:       "Echo",
	FrameworkFiber:      "Fiber",
	FrameworkChi:        "Chi",
	FrameworkDjango:     "Django",
	FrameworkFlask:      "Flask",
	FrameworkFastAPI:    "FastAPI",
	FrameworkExpress:    "Express",
	FrameworkNestJS:     "NestJS",
	FrameworkNextJS:     "Next.js",
	FrameworkReact:      "React",
	FrameworkVue:        "Vue",
	FrameworkAngular:    "Angular",
	FrameworkSpring:     "Spring",
	FrameworkSpringBoot: "Spring Boot",
	FrameworkJavaEE:     "Java EE",
}

func (f Framework) String() string {
	if name, ok := frameworkNames[f]; ok {
		return name
	}
	return "Unknown"
}

type frameworkDetector struct {
	Framework Framework
	Check     func(root string) bool
}

var frameworkDetectors = []frameworkDetector{
	{FrameworkGin, hasGoModDep("gin-gonic/gin")},
	{FrameworkEcho, hasGoModDep("labstack/echo")},
	{FrameworkFiber, hasGoModDep("gofiber/fiber")},
	{FrameworkChi, hasGoModDep("go-chi/chi")},

	{FrameworkDjango, hasFile("manage.py")},
	{FrameworkFlask, hasAnyFile("app.py", "flask_app.py")},
	{FrameworkFastAPI, hasFile("main.py")},

	{FrameworkReact, hasPackageJSONDep("react")},
	{FrameworkVue, hasAnyFile("vue.config.js", "vite.config.js")},
	{FrameworkAngular, hasFile("angular.json")},
	{FrameworkNextJS, hasAnyFile("next.config.js", "next.config.mjs")},
	{FrameworkNestJS, hasPackageJSONDep("@nestjs/core")},

	{FrameworkSpring, hasAnyFile("pom.xml", "build.gradle", "build.gradle.kts")},
}

func hasGoModDep(dep string) func(string) bool {
	return func(root string) bool {
		data, err := os.ReadFile(filepath.Join(root, "go.mod"))
		if err != nil {
			return false
		}
		return strings.Contains(string(data), dep)
	}
}

func hasFile(name string) func(string) bool {
	return func(root string) bool {
		_, err := os.Stat(filepath.Join(root, name))
		return err == nil
	}
}

func hasAnyFile(names ...string) func(string) bool {
	return func(root string) bool {
		for _, name := range names {
			if _, err := os.Stat(filepath.Join(root, name)); err == nil {
				return true
			}
		}
		return false
	}
}

func hasPackageJSONDep(dep string) func(string) bool {
	return func(root string) bool {
		data, err := os.ReadFile(filepath.Join(root, "package.json"))
		if err != nil {
			return false
		}
		return strings.Contains(string(data), dep)
	}
}

func DetectFramework(root string) Framework {
	for _, detector := range frameworkDetectors {
		if detector.Check(root) {
			return detector.Framework
		}
	}
	return FrameworkNone
}