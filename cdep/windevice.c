#include <windows.h>
#include "windevice.h"

int keyDownUp(int e)
{
        return (GetAsyncKeyState(e) & 0x8000) ? 1:0;
}

bool setCursorPos(int x, int y)
{
        return SetCursorPos(x, y);
}

void mouseEvent(int dwFlags, int dx, int dy, int dwData, long long dwExtraInfo)
{
        mouse_event(dwFlags, dx, dy, dwData, dwExtraInfo);
}