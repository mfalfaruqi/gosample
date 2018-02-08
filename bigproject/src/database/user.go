package database

import (
	_ "github.com/lib/pq" // Used for query
)

// GetUserByName for get user by name
func GetUser(name string, limit int) ([]User, error) {
	var userList []User

	query := `
		select 
			user_id,
			full_name,
			msisdn,
			user_email,
			COALESCE(to_char(birth_date, 'YYYY-MM-DD HH24:MI:SS'), '') as birth_date,
			COALESCE(to_char(create_time, 'YYYY-MM-DD HH24:MI:SS'), '') as create_time,
			COALESCE(to_char(update_time, 'YYYY-MM-DD HH24:MI:SS'), '') as update_time,
			case when birth_date is null then 0 else extract(year from age(now(), birth_date)) end as age
		from 
			ws_user 
		where 
			lower(full_name) like lower('%'||$1||'%')
		limit $2
	`
	err := dbUser.Select(&userList, query, name, limit)

	if err != nil {
		return userList, err
	}

	return userList, nil
}
