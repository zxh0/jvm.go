# jvm.go
A JVM written in Go.
![jvm.go Logo](logo.png)

# Introduction
jvm.go is a toy JVM (which is far from complete) programmed in Go. The main purpose of this project is learning Go and the JVM. So the number one goal of the project is readability of code. The basic idea is to just implement the core JVM, and use `rt.jar` (from OpenJDK) as its class library. The garbage collector is implemented by directly using Go’s GC. 



# My dev environment
  * Mac OS X 10.14.6
  * Java 13.0.1 (installed through [SDKMAN](https://sdkman.io/))
  * Go 1.13



# Build jvm.go
```sh
git clone https://github.com/zxh0/jvm.go.git
cd jvm.go
git checkout jdk13

# build test module
cd test/hw_module
sh test.sh 

# run "Hello, World!"
go run github.com/zxh0/jvm.go/cmd/java \
  -p test/hw_module/out/ \
  -m hello.modules/hello.HelloWorld \
  -Xjre ~/.sdkman/candidates/java/13.0.1-open/
```

