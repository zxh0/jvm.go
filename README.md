# jvm.go
A JVM written in Go.
![jvm.go Logo](https://raw.githubusercontent.com/zxh0/jvm.go/master/jvmgo.png)

# Introduction
jvm.go is a toy JVM (which is far from complete) programmed in Go. The main purpose of this project is learning Go and the JVM. So the number one goal of the project is readability of code. The basic idea is to just implement the core JVM, and use `rt.jar` (from OpenJDK) as its class library. The garbage collector is implemented by directly using Go’s GC. 

# My dev environment
  * Mac OS X 10.10.2
  * Java 1.8.0_31
  * Go 1.4

# Build jvm.go
```sh
go get github.com/zxh0/jvm.go/jvmgo
```

# Run jvm.go using your Java installation
Ensure your Java version is 1.8.0_32 and JAVA_HOME env was set
```sh
jvmgo -XuseJavaHome -cp path/to/jars:path/to/classes HelloWorld
```

# Run jvm.go using Zulu
Download [zulu1.8.0_31-8.5.0.1-macosx.zip](http://www.azulsystems.com/products/zulu/downloads#mac) ([Zulu](http://www.azulsystems.com/products/zulu) is a certified build of OpenJDK that is fully compliant with the Java SE standard.) and unzip it to somewhere, Copy `jvmgo` from `$GOPATH/bin/` into unzipped folder 
```sh
cd path/to/zulu1.8.0_31-8.5.0.1-macosx
jvmgo -cp path/to/jars:path/to/classes HelloWorld
```
