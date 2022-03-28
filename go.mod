module github.com/KlyuchnikovV/edicode

go 1.16

require (
	github.com/KlyuchnikovV/buffer v0.0.0-20211215134012-1613b11b850a
	github.com/KlyuchnikovV/stack v0.0.0-20210427103552-f6f21f8b4227
	github.com/mitchellh/mapstructure v1.4.3
	github.com/stretchr/testify v1.7.0
	github.com/wailsapp/wails v1.16.9
	golang.org/x/net v0.0.0-20211215060638-4ddde0e984e9
	golang.org/x/sys v0.0.0-20211214234402-4825e8c3871d // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/KlyuchnikovV/buffer => ../buffer
	github.com/KlyuchnikovV/gapbuf => ../gapbuf
	github.com/KlyuchnikovV/linetree => ../linetree
)
