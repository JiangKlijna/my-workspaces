
#include <stdio.h>
#include "JniTest.h"

JNIEXPORT jint JNICALL Java_JniTest_add(JNIEnv *e, jclass c, jint i, jint j) {
	return i + j;
}

JNIEXPORT void JNICALL Java_JniTest_println(JNIEnv *e, jclass c, jint i) {
	printf("%d\n", i);
}