package model

type SysUser struct {
	BaseModel
	Name string

	Coms []Company `gorm:"foreignKey:UserId"`
}
