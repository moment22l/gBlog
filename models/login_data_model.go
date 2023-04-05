package models

type LoginDataModel struct {
	Model
	UserID    uint      `json:"user_id"` // 登录的用户ID
	UserModel UserModel `gorm:"foreignKey:UserID" json:"user_model"`
	IP        string    `gorm:"size:20" json:"ip"`        // 登录ip
	NickName  string    `gorm:"size:42" json:"nick_name"` // 昵称
	Token     string    `gorm:"size:256" json:"token"`    // token
	Device    string    `gorm:"size:256" json:"device"`   // 登录设备
	Addr      string    `gorm:"64" json:"addr"`           // 地址
}
