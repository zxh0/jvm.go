# jvm.go
A JVM writing in GO...
![jvm.go Logo](https://raw.githubusercontent.com/zxh0/jvm.go/master/jvmgo.png)

# Introduction
jvm.go is a new JVM(which is far from complete) programmed in Go. The main purpose of this project is learning Go and JVM. So, the number one goal of the project is readability of code. The basic idea is just implement the core JVM, and use rt.jar(from OpenJDK) as its class libraries. GC is directly using Go's. For me, the hardest part is Thread and ClassLoader.

# My dev environment:
  * Mac OS X 10.10.2
  * Java 1.8.0_31
  * Go 1.4

# Build `jvm.go`

```sh
go get -d github.com/zxh0/jvm.go/
go install github.com/zxh0/jvm.go/jvmgo
```

# Run jvm.go
Create the following folder and file structure:

```
$GOPATH/bin
├── jvmgo
└── jre/
    └── lib/
        ├── rt.jar
        ├── currency.data
        └── net.properties
```

```sh
jvmgo -cp path/to/jars:path/to/classes HelloWorld
```

# Where to find rt.jar: 

```sh
/Library/Java/JavaVirtualMachines/jdk1.8.0_31.jdk/Contents/Home/jre/lib/rt.jar
```
