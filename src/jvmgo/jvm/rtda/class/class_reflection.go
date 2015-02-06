package class

func (self *Class) GetFields(publicOnly bool) ([]*Field) {
    if publicOnly {
        publicFields := make([]*Field, 0, len(self.fields))
        for _, field := range self.fields {
            if field.IsPublic() {
                n := len(publicFields)
                publicFields = publicFields[:n + 1]
                publicFields[n] = field
            }
        }
        return publicFields
    } else {
        return self.fields
    }
}
