package flag

import (
	"gBlog/global"
	"gBlog/models"
)

// MakeMigrations 迁移数据库(缺点: 只能增加表不会删除表)
func MakeMigrations() {
	var err error
	_ = global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.User2Collects{})
	_ = global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.AdvertModel{},
		&models.ArticleModel{},
		&models.BannerModel{},
		&models.CommentModel{},
		&models.FadeBackModel{},
		&models.MenuBannerModel{},
		&models.MenuModel{},
		&models.MessageModel{},
		&models.TagModel{},
		&models.UserModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("迁移数据库表结构失败")
		return
	}
	global.Log.Infoln("迁移数据库表结构成功")
}
