#include <stdio.h>
#include "windows.h"

BOOL CALLBACK MyEnumProc(HWND hwnd, LPARAM lParam){
	char classname[100];
	char title[100];
	GetClassName(hwnd, classname, 100);
	GetWindowText(hwnd, title, 100);
	printf("hwnd is %x, classname is %s, title is %s\n", hwnd, classname, title);
}

int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance,LPSTR lpCmdLine, int nCmdShow){
	EnumWindows(MyEnumProc, (LPARAM)0);
	return 0;
}