package class

const (
    ACC_PUBLIC          = 0x0001
    ACC_PRIVATE         = 0x0002
    ACC_PROTECTED       = 0x0004
    ACC_STATIC          = 0x0008
    ACC_FINAL           = 0x0010
    ACC_SUPER           = 0x0020
    ACC_SYNCHRONIZED    = 0x0020
    ACC_VOLATILE        = 0x0040
    ACC_BRIDGE          = 0x0040
    ACC_TRANSIENT       = 0x0080
    ACC_VARARGS         = 0x0080
    ACC_NATIVE          = 0x0100
    ACC_INTERFACE       = 0x0200
    ACC_ABSTRACT        = 0x0400
    ACC_STRICT          = 0x0800
    ACC_SYNTHETIC       = 0x1000
    ACC_ANNOTATION      = 0x2000
    ACC_ENUM            = 0x4000
)

type AccessFlags struct {
    accessFlags uint16
}

func (self *AccessFlags) isPublic() (bool) {
    return 0 != self.accessFlags & ACC_PUBLIC
}
func (self *AccessFlags) isPrivate() (bool) {
    return 0 != self.accessFlags & ACC_PRIVATE
}
func (self *AccessFlags) isProtected() (bool) {
    return 0 != self.accessFlags & ACC_PROTECTED
}
func (self *AccessFlags) isStatic() (bool) {
    return 0 != self.accessFlags & ACC_STATIC
}
func (self *AccessFlags) isFinal() (bool) {
    return 0 != self.accessFlags & ACC_FINAL
}
func (self *AccessFlags) isSuper() (bool) {
    return 0 != self.accessFlags & ACC_SUPER
}
func (self *AccessFlags) isSynchronized() (bool) {
    return 0 != self.accessFlags & ACC_SYNCHRONIZED
}
func (self *AccessFlags) isVolatile() (bool) {
    return 0 != self.accessFlags & ACC_VOLATILE
}
func (self *AccessFlags) isBridge() (bool) {
    return 0 != self.accessFlags & ACC_BRIDGE
}
func (self *AccessFlags) isTransient() (bool) {
    return 0 != self.accessFlags & ACC_TRANSIENT
}
func (self *AccessFlags) isVarargs() (bool) {
    return 0 != self.accessFlags & ACC_VARARGS
}
func (self *AccessFlags) isNative() (bool) {
    return 0 != self.accessFlags & ACC_NATIVE
}
func (self *AccessFlags) isInterface() (bool) {
    return 0 != self.accessFlags & ACC_INTERFACE
}
func (self *AccessFlags) isAbstract() (bool) {
    return 0 != self.accessFlags & ACC_ABSTRACT
}
func (self *AccessFlags) isStrict() (bool) {
    return 0 != self.accessFlags & ACC_STRICT
}
func (self *AccessFlags) isSynthetic() (bool) {
    return 0 != self.accessFlags & ACC_SYNTHETIC
}
func (self *AccessFlags) isAnnotation() (bool) {
    return 0 != self.accessFlags & ACC_ANNOTATION
}
func (self *AccessFlags) isEnum() (bool) {
    return 0 != self.accessFlags & ACC_ENUM
}
