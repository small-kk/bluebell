package routes

import (
	"app/controller"
	"app/logger"
	"app/middleware"
	"app/pkg/validator"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	//初始化gin框架内置校验器的翻译器
	if err := validator.InitTrans("zh"); err != nil {
		zap.L().Error("init validator trans failed", zap.Error(err))
	}

	//新建一个引擎
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//加载html文件和静态文件
	r.Static("/static", "./static/")
	r.LoadHTMLFiles("./templates/index.html")

	// 首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := r.Group("/api/v1")

	//注册用户路由
	v1.POST("/signup", controller.SignUpHandler)

	//登录用户路由
	v1.POST("/login", controller.LoginHandler)

	//获取社区分类列表
	v1.GET("/community", controller.CommunityHandler)
	//根据社区id查找社区详情
	v1.GET("/community/:id", controller.CommunityDetailHandler)

	//根据帖子id查询帖子详情
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	//按最新创建时间获取帖子列表
	v1.GET("/post/list", controller.GetPostListHandler)
	//按照分数获取帖子列表
	v1.GET("/post/list2", controller.GetPostList2Handler)

	// 刷新access token
	//r.GET("/refresh_token", controller.RefreshTokenHandler)

	//注册中间件
	v1.Use(middleware.JWTAuthMiddleware(), middleware.RateLimitMiddleware(time.Millisecond, 1000))
	{
		//创建贴子
		v1.POST("/post", controller.CreatePostHandler)
		//帖子投票
		v1.POST("/vote", controller.PostVoteHandler)
	}

	//用户访问路径不存在时，返回一个404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})

	return r
}
