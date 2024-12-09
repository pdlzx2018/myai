package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/your-username/chatnio/internal/service"
)

// loginRequest 定义登录请求的参数结构
type loginRequest struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// registerRequest 定义注册请求的参数结构
type registerRequest struct {
	Username string `json:"username" binding:"required"`    // 用户名
	Password string `json:"password" binding:"required"`    // 密码
	Email    string `json:"email" binding:"required,email"` // 邮箱
}

// userService 用户服务实例
var userService = service.NewUserService()

// Login 处理用户登录请求
// @Summary 用户登录
// @Description 验证用户名密码并返回JWT token
// @Accept json
// @Produce json
// @Param request body loginRequest true "登录参数"
// @Success 200 {object} gin.H{"token": "string"}
// @Failure 400 {object} gin.H{"error": "string"}
// @Failure 401 {object} gin.H{"error": "string"}
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var req loginRequest
	// 解析请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}

	// 调用登录服务
	token, err := userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// 返回token
	c.JSON(200, gin.H{
		"token": token,
	})
}

// Register 处理用户注册请求
// @Summary 用户注册
// @Description 创建新用户
// @Accept json
// @Produce json
// @Param request body registerRequest true "注册参数"
// @Success 200 {object} gin.H{"message": "string"}
// @Failure 400 {object} gin.H{"error": "string"}
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var req registerRequest
	// 解析请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}

	// 调用注册服务
	if err := userService.Register(req.Username, req.Password, req.Email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
