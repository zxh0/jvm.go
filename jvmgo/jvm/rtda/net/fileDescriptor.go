package rtda

type FileDescriptor struct {
	id          int
	inetAddress string
	port        int
}

func (self *FileDescriptor) GetId() int {
	return self.id
}

func (self *FileDescriptor) SetId(id int) {
	self.id = id
}

func (self *FileDescriptor) SetPort(port int) {
	self.port = port
}

func (self *FileDescriptor) GetPort() int {
	return self.port
}

func (self *FileDescriptor) GetInetAddress() string {
	return self.inetAddress
}

func (self *FileDescriptor) SetInetAddress(address string) {
	self.inetAddress = address
}
