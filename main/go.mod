module main

go 1.12

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/someuser/someproject v0.0.0
)

replace github.com/someuser/someproject => ../src
