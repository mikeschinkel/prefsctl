package stdlibex

import (
	"log/slog"
	"runtime"
)

type Frame runtime.Frame

type CallStack struct {
	Frames []Frame
}

func Caller(depth int) (frame Frame) {
	if depth <= 0 {
		panic("Must specify a callstack size of 1 or greater.")
	}
	pc := make([]uintptr, depth+1)
	n := runtime.Callers(depth, pc)
	frames := runtime.CallersFrames(pc[:n])
	f, _ := frames.Next()
	return Frame(f)
}

func Callers(depth int) (callStack CallStack) {
	if depth <= 0 {
		panic("Must specify a callstack size of 1 or greater.")
	}
	pc := make([]uintptr, depth+1)
	n := runtime.Callers(depth, pc)
	callStack.Frames = make([]Frame, depth)
	frames := runtime.CallersFrames(pc[:n])
	frame, more := frames.Next()
	slog.Info("Frame:", "frame", frame)
	if !more {
		goto end
	}
	for i := 0; i < depth; i++ {
		frame, more := frames.Next()
		callStack.Frames[i] = Frame(frame)
		slog.Info("Func:", "name", frame.Function)
		if !more {
			break
		}
	}
end:
	return callStack
}
