# Modules and packages

## Module vs package
- **Module** = the unit of distribution. One `go.mod` per module. Roughly equivalent to a C# assembly or NuGet package.
- **Package** = the unit of code organisation. One folder = one package. Roughly equivalent to a C# namespace.
- A module contains one or more packages. Usually one module per repo.

## `go.mod` essentials
- First line declares the module name: `module github.com/user/repo`.
- The module name is the **import prefix** for every package inside it.
- It has no automatic link to the folder or repo name — whatever string you pass to `go mod init` becomes the identity.
- Check it: `go list -m` or just read line 1 of `go.mod`.

## Naming the module
- Convention: GitHub-style path — `github.com/<user>/<repo>` — even if you never publish. Works locally and stays valid if you ever push.
- Short bare names (`igposter`) work locally but can't be `go get`-ed.
- Avoid passing a filename to `go mod init` by accident (`go mod init main.go` → module literally named `main.go`).

## Renaming a module
- No dedicated rename command.
- Edit `go.mod` line 1 by hand, OR run `go mod edit -module=<newname>`.
- Neither updates `import` statements in `.go` files — find-and-replace those yourself.
- Last resort: delete `go.mod` + `go.sum`, re-run `go mod init <newname>`, then `go mod tidy`. Throws away any `require`/`replace` directives.

## Packages
- Every `.go` file must start with `package <name>` — not optional. It's how Go knows which package the file belongs to.
- All `.go` files in the same folder must declare the **same** package name. Together they form one package.
- The folder name and package name usually match, but don't strictly have to.
- `package main` is special: it produces an executable. Every other package is a library.

## Importing your own packages
- Import path = `<module path> + <subfolder path>`.
- Example: module `github.com/user/repo`, folder `repo/foo/bar/` → `import "github.com/user/repo/foo/bar"`.
- The URL-looking import isn't fetched from the network when it's your own module — it's just a string identifier.
- Use exported (capitalised) names from the imported package: `bar.DoThing()`.

## Folder nesting ≠ package nesting
- Go has no concept of sub-packages. `foo/bar/` is just a package called `bar` that happens to live inside `foo/`.
- They're siblings to Go, not parent/child. Import each independently.
- Folder depth only affects the import path string.

## Nested `go.mod` is (almost always) a mistake
- A `go.mod` in a subfolder creates a **separate module**. The parent can no longer import its packages directly.
- Symptom: "cannot find package" / "no required module provides package" when importing your own subfolder.
- Fix: delete the nested `go.mod` so the subfolder collapses back into a regular package of the parent module.
- Only keep a nested `go.mod` if you genuinely want to publish that subfolder as its own versioned library.

## Project layout
- One `go.mod` at the project root (or wherever you want the module to start).
- Group code by **feature/responsibility**, one folder per package: `config/`, `ig/`, `drive/`, etc.
- Keep `main.go` thin — wire packages together, push logic into the packages.
- Putting `go.mod` inside a `src/` subfolder works but is unidiomatic; most Go projects keep it at the repo root.

## What to commit
- Commit `go.mod` and `go.sum`. Both are required for reproducible builds.
- Don't commit `vendor/` unless you've deliberately chosen vendoring.

## Useful commands
- `go mod init <name>` — create a new module.
- `go mod tidy` — add missing / drop unused dependencies, rewrite `go.sum`.
- `go mod edit -module=<name>` — scripted rename.
- `go list -m` — print current module name.
- `go build ./...` — build every package in the module.
- `go run .` — build and run the `main` package in the current directory.
