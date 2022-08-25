module github.com/KlyuchnikovV/edicode

go 1.16

require (
	github.com/KlyuchnikovV/simple_buffer v0.0.0-20220818203734-248703ddc7ea
	github.com/KlyuchnikovV/stack v0.0.0-20210427103552-f6f21f8b4227
	github.com/mitchellh/mapstructure v1.4.3
	github.com/stretchr/testify v1.8.0
	github.com/wailsapp/wails v1.16.9
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d // indirect
	golang.org/x/net v0.0.0-20211215060638-4ddde0e984e9
	golang.org/x/sys v0.0.0-20211214234402-4825e8c3871d // indirect
)

replace (
	github.com/KlyuchnikovV/buffer => ../buffer
	github.com/KlyuchnikovV/gapbuf => ../gapbuf
	github.com/KlyuchnikovV/linetree => ../linetree
	github.com/KlyuchnikovV/simple_buffer => ../simple_buffer
)
