# Định nghĩa model

Go-pg sử dụng công nghệ ORM (tức Object-relation mapping) giúp ánh xạ bảng cơ sở dữ liệu vào trong struct<br>
Điều đấy có nghĩa là với mỗi struct trong golang có thể dùng làm đại diện để truy vấn đến bảng trong database và trả ra đối tượng struct với giá trị tương ứng.

## 1. Ví dụ về định nghĩa model
```go
type User struct{
	tableName struct{} `pg:"auth.users"`    //Tên schema auth, bảng users

	Id int `pg:"type:serial,pk"`    //Cột id đặt là primary key, với kiểu là serial

	FirstName string    //String sẽ tương ứng với kiểu text trong database, tên cột là first_name

	LastName string     //Tương tự với tên là last_name

	Email string	`pg:",unique"`  //Email điều kiện là không được trùng lặp

	Password string
}
```
Go-pg có thể tự động nhận biết tên trường, tên struct, kiểu biến để khởi tạo bảng database. Ví dụ:
- nếu tên struct là `Genre` sẽ mặc định có tên bảng là `genres`
- tên trường là `FullName` sẽ có tên cột là `full_name`
- Kiểu string sẽ tương ứng với kiểu text trong database, int thành bigint,...
- Trường có tên là Id tự động được thêm vào primary key.

Nếu không muốn để mặc định thì có thể ghi đè bằng những tags như: `pg:"type:kiểu"`,`pg:"tên cột"`,... hay thêm các ràng buộc cho cột như `pg:,unique`. Ví dụ trường id trong model User trên thay vì để kiểu mặc định là bigint, ta thay bằng serial. <br>
Có thể tham khảo thêm [tại đây](https://pg.uptrace.dev/models/)


## 2. Định nghĩa model quan hệ 1 - 1
Giả sử ta có 2 bảng User với Profile, với mỗi User chỉ có một Profile và ngược lại. Đây là quan hệ một một. Ta có thể định nghĩa model như sau

```go
type User struct{
	tableName struct{} `pg:"auth.users"`
    
    	Id int `pg:"serial"`

	Name string

    	Profile *Profile `pg:"rel:belongs-to"`  //Profile thuộc về User
}
```

```go
type Profile struct{
	tableName struct{} `pg:"auth.profile"`

	Id     int	`pg:"type:serial"`

    	Avatar string

    	UserId int  `pg:",unique"` //Id của User mà Profile thuộc về, đây là trường cần thiết để go-pg hiểu được quan hệ
}
```
Lưu ý: - Việc đặt unique ở UserId trong Profile không bắt buộc, chỉ nhằm để không có UserId nào giống nhau, như vậy sẽ đảm bảo quan hệ thực sự là 1 - 1 hơn. Nếu không có unique ta có thể có 2 Profile với cùng một UserId, như vậy sẽ mất đi bản chất quan hệ 1 - 1.

## 3. Định nghĩa model quan hệ 1 - nhiều
Một user có thể có nhiều bài post. Ta có thể định nghĩa model thể hiện quan hệ này như sau:
```go
type User struct{
	tableName struct{} `pg:"auth.users"`   

	Name string

    	Posts []Post `pg:"rel:has-many"` //Với mỗi user có thể lấy ra nhiều bài post của chính user đó
}
```

```go
type Post struct{
	tableName struct{} `pg:"blog.post"`

	Id int `pg:"type:serial" `

	Content string `pg:",notnull"`

	Title string	`pg:",notnull"`

	UserId int `pg:"type:integer"`  //Đây là trường cần thiết để go-pg hiểu được quan hệ
	
	User *User `pg:"rel:has-one"`   //Với mỗi bài post chỉ ứng với một user
}
```

## 4. Định nghĩa model quan hệ nhiều - nhiều
Một User có nhiều Role, như admin, teacher,... và ngược lại một Role cũng có thể có nhiều User. Đây là quan hệ nhiều - nhiều. Trong quan hệ nhiều - nhiều, cách thường thấy là tạo bảng thứ 3 chứa hai trường id của 2 bảng còn lại. Ta có thể định nghĩa model thể hiện quan hệ này như sau:

```go
type User struct{
	tableName struct{} `pg:"auth.users"`

	Id int `pg:"type:serial"`

	Name string
	
	Roles []Role `pg:"many2many:auth.user_role"`    //Quan hệ nhiều nhiều truy vấn tham chiếu đến bảng auth.user_role
}
```
```go
type Role struct{
	tableName struct{} `pg:"auth.role"`	

	Id int	`pg:"type:serial"`

	Name string
}
```
```go
type UserRole struct{   //Bảng thứ 3 tên auth.user_role thể hiện quan hệ giữa 2 bảng trên

	tableName struct{} `pg:"auth.user_role"`

	UserId int  // Id của User

	RoleId int  // Id của Role
}
```
Để ORM hiểu được bảng user_role là bảng thể hiện quan hệ nhiều - nhiều của 2 bảng user và role. Ta cần dùng hàm RegisterTable như sau:
```go
    orm.RegisterTable((*UserRole)(nil))
```  






