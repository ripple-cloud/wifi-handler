#Wifihandler

##Install

In your working directory

Install the library and command line program from [go-bindata] (https://github.com/jteeuwen/go-bindata)

```go
go get github.com/jteeuwen/go-bindata/...
```

Convert static files to Go source code:

```go
go-bindata templates public/... 
```

Cross-compile for raspberry pi

```go
GOARCH=arm GOOS=linux GOARM=5 go build
```

Copy to raspberry pi

```go
scp wifi-handler root@192.168.2.1:wifi-handler
```

Run 

```go
./wifi-handler [or name of binary]
```
