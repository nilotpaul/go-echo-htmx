---
title: Go + Echo + Templates
---

# Go + Echo + Templates

This is a minimal project template designed to be highly configurable for your requirements. It includes only two essential dependency.

# Prerequisites

- Go
- Node.js with your preferred package manager (e.g., npm, yarn, or pnpm)
- [wgo](https://github.com/bokwoon95/wgo) for live server reload.

# Installation

**Run: `gospur init [project-name]`**

## Post Installation

```sh
# Needed for live reload
go install github.com/bokwoon95/wgo@latest
# Install node Deps
npm install
# Install Go Deps
go mod tidy
```

**To start dev server run:**

```sh
make dev
```

**To start prod server run:**

```
make
```

# Deployment

You only need:

- The built binary in `bin` folder.
- The bundled assets in `public` folder.
- Dockerfile coming soon 🔜.

# How easy is it to use?

```go
func handleGetHome(c echo.Context) error {
	return c.Render(http.StatusOK, "Home", map[string]any{
		"Title": "GoSpur",
		"Desc":  "Best for building Full-Stack Applications with minimal JavaScript",
	})
}
```

```html
{{ define "Home" }}
<h1 class="text-4xl">{{ .Ctx.Title }}</h1>
<p class="mt-4">{{ .Ctx.Desc }}</p>
{{ end }}
```

Only this much code is needed to render a page.

# Styling

- If you've selected tailwind, then no extra configuration is needed, start adding classes in any html file it'll just work.
- You can always use plain css (even with tailwind), again it'll just work.

# Quick Tips

- **HTML Routes:** Render templates using handlers like the example above.
- **JSON Routes:** Prefix API endpoints with `/api/json`. The configuration ensures JSON responses even on errors.

For example, `/api/json/example` will always return a JSON response, whereas `/example` would render a template or custom HTML error pages.

# Advanced Usage

**You can also install any npm library and use it.**

1.  Install the library you want.
2.  Update the esbuild configuration:

    ```js
    build({
      // Add the main entrypoint
      entryPoints: ["node_modules/some-library/index.js"],
    });
    ```

3.  Include the bundled script in your templates:
    your lib will be bundled and store in `public/bundle`, find the exact path and include in your templates.

    ```html
    <script src="/public/bundle/some-library.js"></script>
    ```

# Links to Documentation

- [Echo](https://echo.labstack.com)
- [Esbuild](https://esbuild.github.io)