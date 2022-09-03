module github.com/KlyuchnikovV/edicode

go 1.16

require (
	github.com/KlyuchnikovV/simple_buffer v0.0.0-20220818203734-248703ddc7ea
	github.com/KlyuchnikovV/stack v0.0.0-20210427103552-f6f21f8b4227
	github.com/mitchellh/mapstructure v1.4.3
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/wailsapp/wails v1.16.9
	github.com/wailsapp/wails/v2 v2.0.0-beta.44.2
	golang.org/x/crypto v0.0.0-20211215165025-cf75a172585e // indirect
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/net v0.0.0-20211215060638-4ddde0e984e9
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/KlyuchnikovV/buffer => ../buffer
	github.com/KlyuchnikovV/gapbuf => ../gapbuf
	github.com/KlyuchnikovV/linetree => ../linetree
	github.com/KlyuchnikovV/simple_buffer => ../simple_buffer
)
