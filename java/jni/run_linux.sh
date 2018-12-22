
echo $JAVA_HOME

$JAVA_HOME/bin/javac JniTest.java
$JAVA_HOME/bin/javac -h . JniTest.java
$JAVA_HOME/bin/javah JniTest
gcc -I $JAVA_HOME/include -I $JAVA_HOME/include/linux -shared -s -o Jni.so JniTest.c
$JAVA_HOME/bin/java JniTest
