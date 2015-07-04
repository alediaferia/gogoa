#include "g_window.h"

#include <stdlib.h>

void *NewWindow(int x, int y, int width, int height)
{
    NSWindow* window = [[NSWindow alloc] initWithContentRect:NSMakeRect(x, y, width, height)
                                              styleMask:NSTitledWindowMask
                                              backing:NSBackingStoreBuffered defer:NO];

    return window;
}

void MakeKeyAndOrderFront(void *self) {
    NSWindow *window = self;
    [window makeKeyAndOrderFront:nil];
}

void SetTitle(void *self, char *title) {
    NSWindow *window = self;
    NSString *nsTitle =  [NSString stringWithUTF8String:title];

    [window setTitle:nsTitle];
    free(title);
}
