package main

import "github.com/ArshadKhan-007/tutorial/auth"

// Pacakages in Go is a collection of Go source files in the same directory.
// Every file in a Go project belongs to a package.
// The 'package' keyword at the top defines which package it belongs to.
// Packages help organize code and make it reusable across projects.

// WHY DO WE USE PACKAGES?
//- Code organization (keeps your project modular).
// -Reusability (write once, import anywhere).
// -Namespace separation (avoid naming conflicts).
// -Easier testing and maintenance.

// WHAT IS go.mod?
//The 'go.mod' file defines the module (project) name and its dependencies.
// - It’s like a manifest that tells Go:
//   👉 what your project is called (module path)
//   👉 which external packages your project depends on
// - Created using:  `go mod init <module-name>`
// - Tracks versions of imported packages for reproducibility.

func main() {
	auth.LoginWithCredentials("ArshadKhan-007", "secret")
}
