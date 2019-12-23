#ifndef WINDEVICE_H_
#define WINDEVICE_H_

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

// 键按下或者抬起
// 1-处于按下状态
// 0-处于抬起状态
int keyDownUp(int e);
// 设置屏幕坐标
bool setCursorPos(int x, int y);
// 鼠标事件
void mouseEvent(int dwFlags, int dx, int dy, int dwData, long long dwExtraInfo);

#ifdef __cplusplus
}
#endif

#endif // WINDEVICE_H_
