# jvm.go
A JVM written in Go.
![jvm.go Logo](https://raw.githubusercontent.com/zxh0/jvm.go/master/jvmgo.png)

# Introduction
jvm.go is a new JVM (which is far from complete) programmed in Go. The main purpose of this project is learning Go and the JVM. So the number one goal of the project is readability of code. The basic idea is to just implement the core JVM, and use `rt.jar` (from OpenJDK) as its class library. The garbage collector is implemented by directly using Go’s GC. For me, the hardest part is `Thread` and `ClassLoader`.

# My dev environment
  * Mac OS X 10.10.2
  * Java 1.8.0_31
  * Go 1.4

# Build jvm.go
```sh
go get github.com/zxh0/jvm.go/jvmgo
```

# Run jvm.go
```
Download [zulu1.8.0_31-8.5.0.1-macosx.zip](http://cdn.azulsystems.com/zulu/2015-01-8.5-bin/zulu1.8.0_31-8.5.0.1-macosx.zip) and unzip it to somewhere
Copy $GOPATH/bin/jvmgo into zulu1.8.0_31-8.5.0.1-macosx
```
```sh
cd path/to/zulu1.8.0_31-8.5.0.1-macosx
jvmgo -cp path/to/jars:path/to/classes HelloWorld
```

# Example
Create a Java source file, `Main.java`:
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

Compile `Main.java`, which will generate `Main.class`:
```sh
javac Main.java
``` 

Run with `jvmgo`:
```sh
jvmgo Main
```

Output:
```sh
hello world
```
