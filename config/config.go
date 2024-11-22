package config

import "github.com/manifoldco/promptui"

// For adding styles to console output.
var (
	ErrMsg     = promptui.Styler(promptui.FGRed)
	SuccessMsg = promptui.Styler(promptui.FGGreen, promptui.FGBold)
	NormalMsg  = promptui.Styler(promptui.FGWhite)
	FaintMsg   = promptui.Styler(promptui.FGFaint)
)

// UILibrary represents an UI Library and `DependsOn`
// which means it can depend on any chosen CSS Strategy (framework).
type (
	ProjectCtx map[string]any

	UILibrary struct {
		// Name of the UI Library
		Name string

		// An UI Library can depend on any chosen CSS Strategy.
		// If it's independent, `DependsOn` should be an empty string.
		DependsOn string
	}
)

// Prompt options.
var (
	WebFrameworkOpts = []string{
		"Echo",
	}
	ExtraOpts = []string{
		"HTMX",
	}

	CssStrategyOpts = []string{
		"Tailwind",
		"Vanilla CSS",
	}
	UILibraryOpts = map[string][]string{
		"Preline": {"Tailwind"},
	}
)

// Project file structure
var (
	ProjectBaseFiles = map[string]string{
		"config/env.go":          "base/env.go.tmpl",
		"web/styles/globals.css": "base/globals.css.tmpl",
		".gitignore":             "base/gitignore.tmpl",
		"Makefile":               "base/makefile.tmpl",
		"README.md":              "base/readme.md.tmpl",
		"esbuild.config.js":      "base/esbuild.config.js.tmpl",
		"package.json":           "base/package.json.tmpl",
		"tailwind.config.js":     "base/tailwind.config.js.tmpl",
		"main.go":                "base/main.go.tmpl",
	}

	// Template path is not required anymore for pages.
	// We're processing these as raw files.
	ProjectPageFiles = map[string]string{
		"web/Home.html":  "",
		"web/Error.html": "",
	}

	ProjectAPIFiles = map[string]string{
		"api/api.go":     "api/api.go.echo.tmpl",
		"api/route.go":   "api/route.go.echo.tmpl",
		"api/handler.go": "api/handler.go.echo.tmpl",
	}
)
