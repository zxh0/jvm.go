# jvm.go
A JVM written in GO...
![jvm.go Logo](https://raw.githubusercontent.com/zxh0/jvm.go/master/jvmgo.png)

# Introduction
jvm.go is a new JVM(which is far from complete) programmed in Go. The main purpose of this project is learning Go and JVM. So, the number one goal of the project is readability of code. The basic idea is just implement the core JVM, and use rt.jar(from OpenJDK) as its class libraries. GC is directly using Go's. For me, the hardest part is Thread and ClassLoader.

# My dev environment:
  * Mac OS X 10.10.2
  * Java 1.8.0_31
  * Go 1.4

# Build jvm.go

```sh
go get github.com/zxh0/jvm.go/jvmgo
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

# Example
Create a java source file: Main.java
```java
public class Main {

    public static void main(String []args){
        String val = "hello world";
        if(args != null && args.length > 0){
            val = args[0];
        }
        System.out.println(val);
    }
}
```

Compile Main.java, it will generate Main.class
```sh
javac Main.java
``` 

Run with jvmgo
```sh
jvmgo Main
```

Output
```sh
hello world
```
