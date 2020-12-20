package main

import (
	"github.com/miraclecoco/gomirror/pkg/subpkg1"
	"github.com/miraclecoco/gomirror/pkg/subpkg2"
	"github.com/miraclecoco/gomirror/pkg/subpkg2/subpkg4"
	"github.com/miraclecoco/gomirror/pkg/subpkg3"
)

func main() {
	subpkg1.Subpkg1Func()
	subpkg2.Subpkg2Func()
	subpkg3.Subpkg3Func()
	subpkg4.Subpkg2Func()
}
