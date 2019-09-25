package heap

const (
	ACC_PUBLIC       = 0x0001
	ACC_PRIVATE      = 0x0002
	ACC_PROTECTED    = 0x0004
	ACC_STATIC       = 0x0008
	ACC_FINAL        = 0x0010
	ACC_SUPER        = 0x0020
	ACC_SYNCHRONIZED = 0x0020
	ACC_VOLATILE     = 0x0040
	ACC_BRIDGE       = 0x0040
	ACC_TRANSIENT    = 0x0080
	ACC_VARARGS      = 0x0080
	ACC_NATIVE       = 0x0100
	ACC_INTERFACE    = 0x0200
	ACC_ABSTRACT     = 0x0400
	ACC_STRICT       = 0x0800
	ACC_SYNTHETIC    = 0x1000
	ACC_ANNOTATION   = 0x2000
	ACC_ENUM         = 0x4000
)

type AccessFlags struct {
	accessFlags uint16
}

func (flags *AccessFlags) GetAccessFlags() uint16 {
	return flags.accessFlags
}

func NewAccessFlags(accessFlags uint16) *AccessFlags {
	af := AccessFlags{accessFlags: accessFlags}
	return &af
}

func (flags *AccessFlags) IsPublic() bool {
	return 0 != flags.accessFlags&ACC_PUBLIC
}
func (flags *AccessFlags) IsPrivate() bool {
	return 0 != flags.accessFlags&ACC_PRIVATE
}
func (flags *AccessFlags) IsProtected() bool {
	return 0 != flags.accessFlags&ACC_PROTECTED
}
func (flags *AccessFlags) IsStatic() bool {
	return 0 != flags.accessFlags&ACC_STATIC
}
func (flags *AccessFlags) IsFinal() bool {
	return 0 != flags.accessFlags&ACC_FINAL
}
func (flags *AccessFlags) IsSuper() bool {
	return 0 != flags.accessFlags&ACC_SUPER
}
func (flags *AccessFlags) IsSynchronized() bool {
	return 0 != flags.accessFlags&ACC_SYNCHRONIZED
}
func (flags *AccessFlags) IsVolatile() bool {
	return 0 != flags.accessFlags&ACC_VOLATILE
}
func (flags *AccessFlags) IsBridge() bool {
	return 0 != flags.accessFlags&ACC_BRIDGE
}
func (flags *AccessFlags) IsTransient() bool {
	return 0 != flags.accessFlags&ACC_TRANSIENT
}
func (flags *AccessFlags) IsVarargs() bool {
	return 0 != flags.accessFlags&ACC_VARARGS
}
func (flags *AccessFlags) IsNative() bool {
	return 0 != flags.accessFlags&ACC_NATIVE
}
func (flags *AccessFlags) IsInterface() bool {
	return 0 != flags.accessFlags&ACC_INTERFACE
}
func (flags *AccessFlags) IsAbstract() bool {
	return 0 != flags.accessFlags&ACC_ABSTRACT
}
func (flags *AccessFlags) IsStrict() bool {
	return 0 != flags.accessFlags&ACC_STRICT
}
func (flags *AccessFlags) IsSynthetic() bool {
	return 0 != flags.accessFlags&ACC_SYNTHETIC
}
func (flags *AccessFlags) IsAnnotation() bool {
	return 0 != flags.accessFlags&ACC_ANNOTATION
}
func (flags *AccessFlags) IsEnum() bool {
	return 0 != flags.accessFlags&ACC_ENUM
}
