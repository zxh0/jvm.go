#!/usr/bin/make -f

JAVA_HOME=/Users/zxh/.sdkman/candidates/java/8.0.222-zulu

java: fix-java

# TODO
fix-java: clean build-java
	install_name_tool -add_rpath $(JAVA_HOME)/jre/lib bin/java
	install_name_tool -add_rpath $(JAVA_HOME)/jre/lib/server bin/java

build-java: build-wrapper
	CGO_CFLAGS="-I$(JAVA_HOME)/include -I$(JAVA_HOME)/include/darwin" \
	CGO_LDFLAGS="-L$(JAVA_HOME)/jre/lib -ljava" \
	go build -o bin/java github.com/zxh0/jvm.go/cmd/java

build-wrapper:
	mkdir -p bin
	gcc -I$(JAVA_HOME)/include \
		-I$(JAVA_HOME)/include/darwin \
		-c -o bin/libjniwrapper.dylib jni/wrapper.c

clean:
	rm -rf bin

.PHONY: test
test:
	go test github.com/zxh0/jvm.go/classfile
	go test github.com/zxh0/jvm.go/classpath
	go test github.com/zxh0/jvm.go/instructions/...
	go test github.com/zxh0/jvm.go/jimage
	go test github.com/zxh0/jvm.go/rtda/...
	go test github.com/zxh0/jvm.go/vmutils
