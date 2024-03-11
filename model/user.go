package model

type User struct{
	tableName struct{} `pg:"users"`	
	Id int `pg:"type:serial"`
	FirstName string
	LastName string
	Email string	`pg:",unique"`
	Password string
	Posts []Post `pg:"rel:has-many"`
	Profile *Profile `pg:"rel:belongs-to"`
	Roles []Role `pg:"many2many:auth.user_role"`
}

type Profile struct{
	tableName struct{} `pg:"profile"`	
	Id     int	`pg:"type:serial"`
    Avatar string
	Description string
    UserId int	`pg:",unique"`
}

type Role struct{
	tableName struct{} `pg:"role"`	
	Id int	`pg:"type:serial"`
	Name string
}

type UserRole struct{
	tableName struct{} `pg:"user_role"`	
	UserId int
	RoleId int
}
