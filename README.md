# jvm.go
A JVM writing in GO...
![jvm.go Logo](https://raw.githubusercontent.com/zxh0/jvm.go/master/jvmgo.png)

# My dev environment:
  * Mac OS X 10.10.2
  * Java 1.8.0_31
  * Go 1.4

# Build jvm.go
```
export GOPATH=path/to/jvm.go/
go install jvmgo
```

# Run jvm.go
make folder structures like this:
```
path/to/jvm.go/bin
  jre/
    lib/
      rt.jar
      currency.data
      net.properties
```
```
cd path/to/jvm.go/bin
./jvmgo -cp path/to/3rd.jars:path/to/classes HelloWorld
```

# Where to find rt.jar: 
```
/Library/Java/JavaVirtualMachines/jdk1.8.0_31.jdk/Contents/Home/jre/lib/rt.jar
```
