package rtda

import rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"

type FrameCache struct {
	cachedFrames []*Frame
	frameCount   uint
}

func (self *FrameCache) borrowFrame(thread *Thread, method *rtc.Method) *Frame {
	if self.frameCount > 0 {
		for i, frame := range self.cachedFrames {
			if uint(len(frame.localVars.slots)) > method.MaxLocals() &&
				uint(len(frame.operandStack.slots)) > method.MaxStack() {

				self.cachedFrames[i] = nil
				frame.reset(thread, method)
				return frame
			}
		}
	}
	return newFrame(thread, method)
}

func (self *FrameCache) returnFrame() {

}
