# jvm.go
A JVM writing in GO...
![jvm.go Logo](https://raw.githubusercontent.com/zxh0/jvm.go/master/jvmgo.png)

# My dev environment:
  * Mac OS X 10.10.2
  * Java 1.8.0_31
  * Go 1.4

# Build `jvm.go`

```sh
export GOPATH=path/to/jvm.go/
go install jvmgo
```

# Run jvm.go
Create the following folder and file structure:

```
path/to/jvm.go/bin
├── jvmgo
└── jre/
    └── lib/
        ├── rt.jar
        ├── currency.data
        └── net.properties
```

```sh
cd path/to/jvm.go/bin
./jvmgo -cp path/to/jars:path/to/classes HelloWorld
```

# Where to find rt.jar: 

```sh
/Library/Java/JavaVirtualMachines/jdk1.8.0_31.jdk/Contents/Home/jre/lib/rt.jar
```
