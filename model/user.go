package model

type UserModel struct {
	ID         int    `xorm:"'id' int pk notnull autoincr" json:"-"`
	Username   string `xorm:"'user' varchar(128) notnull "`
	PasswdHash string `xorm:"'passwdhash' varchar(512) notnull "`
}

func (model *UserModel) TableName() string {
	return "users"
}



func (model *UserModel) SetPasswdHash(pwd string) (pwdhash string)  {
	
	return
	
}
