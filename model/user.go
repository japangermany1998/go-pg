package model

type User struct{
	tableName struct{} `pg:"public.users"`	
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
	tableName struct{} `pg:"auth.profile"`	
	Id     int	`pg:"type:serial"`
    Avatar string
	Description string
    UserId int	`pg:",unique"`
}

type Role struct{
	tableName struct{} `pg:"auth.role"`	
	Id int	`pg:"type:serial"`
	Name string
}

type UserRole struct{
	tableName struct{} `pg:"auth.user_role"`	
	UserId int
	RoleId int
}
