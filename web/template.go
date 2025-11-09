package web

import "embed"

//go:embed dist/index.html
var StaticPage embed.FS

//go:embed all:dist/assets
var Assets embed.FS

//go:embed all:dist/static
var Statics embed.FS
