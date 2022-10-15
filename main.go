package main

import (
	"vAdmin/cmd"
)

// @title DevOps API
// @version 0.0.1
// @description Rnotes系统
// @SecurityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @BasePath /

//func main() {
//	configName := "settings"
//
//
//	config.InitConfig(configName)
//
//	gin.SetMode(gin.DebugMode)
//	log.Println(config.MysqlConfig.Port)
//
//	err := gorm.AutoMigrate(orm.Eloquent)
//	if err != nil {
//		log.Fatalln("数据库初始化失败 err: %v", err)
//	}
//
//	if config.ApplicationConfig.IsInit {
//		if err := models.InitDb(); err != nil {
//			log.Fatal("数据库基础数据初始化失败！")
//		} else {
//			config.SetApplicationIsInit()
//		}
//	}
//
//	r := router.InitRouter()
//
//	defer orm.Eloquent.Close()
//
//	srv := &http.Server{
//		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
//		Handler: r,
//	}
//
//	go func() {
//		// 服务连接
//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			log.Fatalf("listen: %s\n", err)
//		}
//	}()
//	log.Println("Server Run ", config.ApplicationConfig.Host+":"+config.ApplicationConfig.Port)
//	log.Println("Enter Control + C Shutdown Server")
//	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt)
//	<-quit
//	log.Println("Shutdown Server ...")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Fatal("Server Shutdown:", err)
//	}
//	log.Println("Server exiting")
//}

func main() {

	/*
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("0 0 0 *", func() {
		log.Println("Run models.PullRelease...")
			models.PullRelease()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
	 */
	cmd.Execute()
}
