#!/bin/sh
set -e

alias javac='~/.sdkman/candidates/java/13.0.0-open/bin/javac'
alias jlink='~/.sdkman/candidates/java/13.0.0-open/bin/jlink'

javac -version | grep 'javac 13'
jlink --version | grep 13

OUT=myjre
rm -rf $OUT
javac --module-source-path src -d out -m hello.modules
jlink --module-path out --add-modules hello.modules,java.base --output $OUT
./$OUT/bin/java -m hello.modules/hello.HelloWorld
