package models

import "BeegoBcRPCCode/db_mysql"

type User struct {
	Id        int    `form:"id"`
	Username string `form:"user_name"`
	Password  string `form:"password"`
	Email     string `form:"E_mail"`
}

//保存数据到数据库的方法
func (u User) SaveUser() (int64, error) {
	row, err := db_mysql.DB.Exec("insert into bcuser(user_name,E_mail,password) "+"values(?,?,?)", u.Username,u.Email,u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

//查询数据库中的数据
func (u User) QueryUser() (*User, error) {
	row := db_mysql.DB.QueryRow("select user_name from bcuser where E_mail= ? and password= ?", u.Email, u.Password)
	err := row.Scan(&u.Username)
	if err != nil {
		return nil, err

	}
	return &u, nil
}
