#include <stdio.h>
#include "windows.h"

#define IsEmpty(str) ((str) == NULL || strlen(str) == 0)

void FillInPassword(HWND hwnd, const char *pwd) {
    HWND edit = FindWindowEx(hwnd, NULL, "Edit", NULL);
    SendMessage(edit, WM_SETFOCUS, 0, 0);
    SendMessage(edit, WM_SETTEXT, 0, (LPARAM) "");
    for (size_t i = 0, n = strlen(pwd); i < n; ++i) {
        keybd_event((BYTE) pwd[i], 0, 0, 0);
    }
    printf("Fill in the password : %s\n", pwd);
}

void ClickOnLogin(HWND hwnd) {
    const char LOGIN[] = {-75, -57, -62, -68}; // 登陆
    HWND button = FindWindowEx(hwnd, NULL, "Button", LOGIN);
    SendMessage(button, BM_CLICK, 0, 0);
}

void FindAndLogin(const char *pwd) {
    const char TITLE[] = {-47, -23, -42, -92, 32, 80, 73, 78, 32, -62, -21}; // 验证 PIN 码
    HWND hwnd = FindWindow(NULL, TITLE);
    if (hwnd == NULL) return;
    printf("Window found : %x\n", hwnd);
    SetWindowPos(hwnd, HWND_TOPMOST, 0, 0, 0, 0, SWP_NOSIZE | SWP_NOMOVE);
    FillInPassword(hwnd, pwd);
    ClickOnLogin(hwnd);
}

int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int nCmdShow) {
    const char *pwd = IsEmpty(lpCmdLine) ? "123456" : lpCmdLine;
    while (1) {
        FindAndLogin(pwd);
        Sleep(1000);
    }
}
