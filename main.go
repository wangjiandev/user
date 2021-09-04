package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/wangjiandev/user/domain/repository"
	us "github.com/wangjiandev/user/domain/service"
	"github.com/wangjiandev/user/handler"
	pb "github.com/wangjiandev/user/proto/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	srv.Init()
	dns := "root:123456@tcp(127.0.0.1:3306)/imooc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// 初始化数据库
	userRepository := repository.NewUserRepository(db)
	//err = userRepository.InitTable()
	userService := us.NewUserService(userRepository)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), &handler.User{
		UserService: userService,
	})

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
