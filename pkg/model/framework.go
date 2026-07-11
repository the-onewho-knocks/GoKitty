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
)

type frameworkDetector struct {
	Framework Framework
	Name      string
	Check     func(root string) bool
}

func hasGoModDep(dep string) func(root string) bool {
	return func(root string) bool {
		data, err := os.ReadFile(filepath.Join(root, "go.mod"))
		if err != nil {
			return false
		}
		return strings.Contains(string(data), dep)
	}
}

func hasFile(name string) func(root string) bool {
	return func(root string) bool {
		info, err := os.Stat(filepath.Join(root, name))
		return err == nil && !info.IsDir()
	}
}

func hasDir(name string) func(root string) bool {
	return func(root string) bool {
		info, err := os.Stat(filepath.Join(root, name))
		return err == nil && info.IsDir()
	}
}

var frameworkDetectors = []frameworkDetector{
	{FrameworkGin, "Gin", hasGoModDep("gin-gonic/gin")},
	{FrameworkEcho, "Echo", hasGoModDep("labstack/echo")},
	{FrameworkFiber, "Fiber", hasGoModDep("gofiber/fiber")},
	{FrameworkChi, "Chi", hasGoModDep("go-chi/chi")},
	{FrameworkDjango, "Django", hasFile("manage.py")},
	{FrameworkFlask, "Flask", hasFile("app.py")},
	{FrameworkFastAPI, "FastAPI", hasFile("main.py")},
	{FrameworkReact, "React", hasDir("node_modules/react")},
	{FrameworkVue, "Vue", hasFile("vue.config.js")},
	{FrameworkAngular, "Angular", hasFile("angular.json")},
	{FrameworkNextJS, "Next.js", hasFile("next.config.js")},
	{FrameworkNestJS, "NestJS", hasGoModDep("nestjs")},
	{FrameworkExpress, "Express", hasFile("node_modules/express")},
	{FrameworkSpring, "Spring", hasFile("pom.xml")},
	{FrameworkSpringBoot, "Spring Boot", hasFile("build.gradle")},
	{FrameworkJavaEE, "Java EE", hasDir("src/main/java")},
}

var frameworkNames = map[Framework]string{
	FrameworkNone:      "None",
	FrameworkGin:      "Gin",
	FrameworkEcho:     "Echo",
	FrameworkFiber:    "Fiber",
	FrameworkChi:      "Chi",
	FrameworkDjango:   "Django",
	FrameworkFlask:    "Flask",
	FrameworkFastAPI:  "FastAPI",
	FrameworkExpress:  "Express",
	FrameworkNestJS:   "NestJS",
	FrameworkNextJS:   "Next.js",
	FrameworkReact:    "React",
	FrameworkVue:      "Vue",
	FrameworkAngular:  "Angular",
	FrameworkSpring:   "Spring",
	FrameworkSpringBoot: "Spring Boot",
	FrameworkJavaEE:   "Java EE",
}

func (f Framework) String() string {
	if n, ok := frameworkNames[f]; ok {
		return n
	}
	return "None"
}

func DetectFramework(root string) Framework {
	for _, d := range frameworkDetectors {
		if d.Check(root) {
			return d.Framework
		}
	}
	return FrameworkNone
}
