module github.com/KlyuchnikovV/edicode

go 1.19

require (
	github.com/KlyuchnikovV/simple_buffer v0.0.0-00010101000000-000000000000
	github.com/KlyuchnikovV/stack v0.0.0-20210427103552-f6f21f8b4227
	github.com/fsnotify/fsnotify v1.5.4
	github.com/mitchellh/mapstructure v1.4.3
	github.com/wailsapp/wails/v2 v2.1.0
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b
)

require (
	github.com/bep/debounce v1.2.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jchv/go-winloader v0.0.0-20210711035445-715c2860da7e // indirect
	github.com/labstack/echo/v4 v4.9.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/leaanthony/go-ansi-parser v1.0.1 // indirect
	github.com/leaanthony/gosod v1.0.3 // indirect
	github.com/leaanthony/slicer v1.6.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tkrajina/go-reflector v0.5.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/wailsapp/mimetype v1.4.1 // indirect
	golang.design/x/clipboard v0.6.2 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/exp/shiny v0.0.0-20221023144134-a1e5550cf13e // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/mobile v0.0.0-20210716004757-34ab1303b554 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace (
	github.com/KlyuchnikovV/buffer => ../../buffer
	github.com/KlyuchnikovV/gapbuf => ../../gapbuf
	github.com/KlyuchnikovV/linetree => ../../linetree
	github.com/KlyuchnikovV/simple_buffer => ../simple_buffer
)
