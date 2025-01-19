package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AppKit
#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>
#include <sys/types.h>
#include <signal.h>

 int killProcessByName(const char* name) {
     NSString *processName = [NSString stringWithUTF8String:name];
     NSWorkspace *workspace = [NSWorkspace sharedWorkspace];
     NSArray *runningApps = [workspace runningApplications];

     int killed = 0;

     for (NSRunningApplication *app in runningApps) {
         if ([[app bundleIdentifier] containsString:@"com.apple.dock"]) {
             pid_t pid = [app processIdentifier];
             if (kill(pid, SIGTERM) == 0) {
                 killed = 1;
             }
         }
     }

     return killed;
 }
*/
import "C"
import (
	"errors"
	"fmt"
)

func KillProcess(name ProcessName) (err error) {
	result := C.killProcessByName(C.CString(string(name)))
	if result != 0 {
		goto end
	}
	err = errors.Join(ErrFailedToKillProcess,
		fmt.Errorf("%s=%s", ProcessNameLogArg, name),
	)
end:
	return err
}
