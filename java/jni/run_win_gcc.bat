@echo off
set jdk_home="D:\Program Files\Java\jdk1.8.0_151"

javac JniTest.java
javah JniTest
gcc -I %jdk_home%\include -I %jdk_home%\include\win32 -shared -s -o Jni.dll JniTest.c
java JniTest