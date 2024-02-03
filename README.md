# TSIS1 Example

## Run project

Main package has a few files, so you should run whole folder like on example bellow.

```bash
go run ./cmd/abr-plus
```

**not particular .go file, you need to run whole package folder**

## Description

Module name `github.com/codev0/go-dev-tsis1`. Remember, we use links as go module name becuase they are unuque and can be used to identify the package, it's not real link, it's just a unique identifier.

Structure:

- `/cmd` folder with executable app named abr-plus.
  - `abr-plus`
    - `main.go` main executable file
    - `handlers.go` all handlers moved here
- `/pkg` folder with `abrplus` package encapsulates all related things to my app.
  - `abr-plus`
    - `abr-plus.go` feneric file with no specific purpose
    - `model.go` file related to data. E.g., types, processing, etc.

**Package import pattern** is `module_name/path_to_package`, for example I want to import `abrplus` package in handlers, thus I write `github.com/codev0/go-dev-tsis1/pkg/abr-plus`.

## Debugging
`.vscode` folder contains configuration to run this project in debug mode in VS Code editor.
