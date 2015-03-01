# jvm.go
A JVM writing in GO...
![jvm.go Logo](https://raw.githubusercontent.com/zxh0/jvm.go/master/jvmgo.png)

# My dev environment:
  * Mac OS X 10.10.2
  * Java 1.8.0_31
  * Go 1.4

Where to find rt.jar: /Library/Java/JavaVirtualMachines/jdk1.8.0_31.jdk/Contents/Home/jre/lib/rt.jar

# Build jvm.go
```
export GOPATH=path/to/jvm.go/
go install jvmgo
```

# Run jvm.go
```
cd path/to/jvm.go/bin
./jvmgo -cp path/to/rt.jar:path/to/classes HelloWorld
```
