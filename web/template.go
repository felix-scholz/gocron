package web

import "embed"

//go:embed dist/index.html
var StaticPage embed.FS
