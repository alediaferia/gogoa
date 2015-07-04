#import "g_application.h"

void*
SharedApplication(void) {
	return [NSApplication sharedApplication];
}

void
Run(void *app) {
    @autoreleasepool {
        NSApplication* a = (NSApplication*)app;
        [a setActivationPolicy:NSApplicationActivationPolicyRegular];
        [a run];
	}
}
