package database

// User for model ws_user
type User struct {
	ID          int    `db:"user_id" json:"id"`
	Name        string `db:"full_name" json:"name"`
	Msisdn      string `db:"msisdn" json:"msisdn"`
	Email       string `db:"user_email" json:"email"`
	BirthDate   string `db:"birth_date" json:"birth_date"`
	CreatedTime string `db:"create_time" json:"created_time"`
	UpdatedTime string `db:"update_time" json:"updated_time"`
	Age         int    `db:"age" json:"age"`
}
