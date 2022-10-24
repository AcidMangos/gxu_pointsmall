package dao

func Migration() {
	// 自动迁移模式
	DBClient.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Product{}, &Order{}, &Admin{})

}
