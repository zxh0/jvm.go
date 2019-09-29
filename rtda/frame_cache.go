package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

type FrameCache struct {
	thread       *Thread
	cachedFrames []*Frame
	frameCount   uint
	maxFrame     uint
}

func newFrameCache(thread *Thread, maxFrame uint) *FrameCache {
	return &FrameCache{
		thread:       thread,
		maxFrame:     maxFrame,
		cachedFrames: make([]*Frame, maxFrame),
	}
}

func (cache *FrameCache) borrowFrame(method *heap.Method) *Frame {
	if cache.frameCount > 0 {
		for i, frame := range cache.cachedFrames {
			if frame != nil &&
				frame.maxLocals >= method.MaxLocals &&
				frame.maxStack >= method.MaxStack {

				cache.frameCount--
				cache.cachedFrames[i] = nil
				frame.reset(method)
				return frame
			}
		}
	}
	return newFrame(cache.thread, method)
}

func (cache *FrameCache) returnFrame(frame *Frame) {
	if cache.frameCount < cache.maxFrame {
		for i, cachedFrame := range cache.cachedFrames {
			if cachedFrame == nil {
				cache.cachedFrames[i] = frame
				cache.frameCount++
				return
			}
		}
	} else {
		for _, cachedFrame := range cache.cachedFrames {
			if frame.maxLocals > cachedFrame.maxLocals {
				cachedFrame.maxLocals = frame.maxLocals
				cachedFrame.LocalVars = frame.LocalVars
				frame.maxLocals = 0
			}
			if frame.maxStack > cachedFrame.maxStack {
				cachedFrame.maxStack = frame.maxStack
				cachedFrame.OperandStack = frame.OperandStack
				frame.maxStack = 0
			}
		}
	}
}
