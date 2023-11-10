package model

type User struct {
	Model                  //基础模型
	Usernmae        string `gorm:"not null;unique;type: varchar(50);index;comment: 用户名"` // 用户名
	Password        string `gorm:"type: varchar(255);not null;comment: 密码"`              //密码
	Avatar          string `gorm:"type: varchar(500);comment: 头像"`                       //头像
	ArticlesNum     int32  `gorm:"comment: 博文数"`                                         //博文数
	Signature       string `gorm:"type: varchar(150);comment:个性签名"`                      //个性签名
	Experience      int32  `gorm:"comment:个性签名"`                                         //经验值
	BackgroundImage string `gorm:"type: varchar(600);comment: 个人简介图片"`                   //个人简介图片
	Level           int32  `gorm:"comment: 等级"`                                          //等级
}
