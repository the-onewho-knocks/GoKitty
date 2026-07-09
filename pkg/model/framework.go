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

var frameworkDetectors = []struct {
	Framework Framework
	Check     func(root string) bool
}{
	{FrameworkGin, hasGoModDep("gin-gonic/gin")},
	{FrameworkEcho, hasGoModDep("labstack/echo")},
	{FrameworkFiber, hasGoModDep("gofiber/fiber")},
	{FrameworkChi, hasGoModDep("go-chi/chi")},
	{FrameworkDjango, hasFile("manage.py")},
	{FrameworkFlask, hasFile("app.py") || hasFile("flask_app.py")},
	{FrameworkFastAPI, hasFile("main.py")},
	{FrameworkReact, hasFile("package.json")},
	{FrameworkVue, hasFile("vue.config.js")},
	{FrameworkAngular, hasFile("angular.json")},
	{FrameworkNextJS, hasFile("next.config.js")},
	{FrameworkNestJS, hasGoModDep("nestjs")},
	{FrameworkSpring, hasFile("pom.xml") || hasFile("build.gradle")},
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
		_, err := os.Stat(filepath.Join(root, name))
		return err == nil
	}
}

func DetectFramework(root string) Framework {
	for _, d := range frameworkNames {
		if d.CheckFunc(root) {
			return d.Framework
		}
	}
	return FrameworkNone
}