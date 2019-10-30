#!/bin/sh
set -e

alias javac='~/.sdkman/candidates/java/13.0.1-open/bin/javac'
alias jlink='~/.sdkman/candidates/java/13.0.1-open/bin/jlink'

javac -version | grep 'javac 13'
jlink --version | grep 13

javac --module-source-path src \
    --module-version 0.1  \
    -d out -m jvmgo.tests
