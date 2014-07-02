package rest

type User struct {
	id string
}

func (self *User) Id() string {
	return self.id
}

func (self *User) SetId(id string) {
	self.id = id
}
