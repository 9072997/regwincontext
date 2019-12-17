go get github.com/akavel/rsrc
rsrc -manifest regwincontext.exe.manifest -arch=amd64 -o regwincontext.syso
go build
