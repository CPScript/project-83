To compile "compiler": `go build -o compiler.exe myapp.go` (replace compiler.exe with the name of the executable you want to send)


When you execute "compiler.exe", it will attempt to download the specified package. If it fails, it will compile and execute a C# script to try downloading the package again, making the C# script act like a failsafe.
