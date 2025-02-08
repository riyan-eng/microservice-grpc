package datastruct

type PermissionRoleAccessList struct {
	Id                string       `db:"id" json:"id"`
	Name              string       `db:"name" json:"name"`
	PermissionsString string       `db:"perms" json:"-"`
	Permissions       []Permission `db:"-" json:"permissions"`
}

type Permission struct {
	Id         string `db:"id" json:"id"`
	Code       string `db:"code" json:"code"`
	Name       string `db:"name" json:"name"`
	Parent     string `db:"parent" json:"parent"`
	FullMethod string `db:"full_method" json:"full_method"`
}
