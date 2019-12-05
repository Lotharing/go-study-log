package mappers

//定义静态常量，用于标识SQL信息
const (
	USER_SELECT rawSQL = iota + 1
	USER_DELETE
	USER_UPDATE
	USER_INSERT
)

//定义全局变量，存放SQL信息
var (
	getUserInfo = `
		SELECT * FROM user;
	`
	addUserInfo = `
		INSERT INTO user (name,age,address) VALUES(?,?,?)
	`
	modifyUserInfo = `
		UPDATE user SET name = ? WHERE Id = ?
	`
	removeUserInfo = `
		DELETE FROM user WHERE id = ?
	`
)

//一般通过model去调用，然后参数直接传静态的SQL标识，然后返回SQL字符串，提供给model使用
func changeUserRawSQL(t rawSQL) string {
	switch t {
	case USER_SELECT:
		return getUserInfo
	case USER_DELETE:
		return removeUserInfo
	case USER_INSERT:
		return addUserInfo
	case USER_UPDATE:
		return modifyUserInfo
	default:
		return ""
	}
}
