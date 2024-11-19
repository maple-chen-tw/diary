package controllers

import (
	models "diary/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JWTSecret = []byte("my-very-secret-key-that-is-random-and-secure-12345")

// 生成 JWT token
func generateJWT(user models.User) (string, error) {
	// 設置 JWT claims (payload)
	claims := jwt.MapClaims{
		"username": user.Username,
		"UserID":   user.UserID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 設定過期時間 24 小時
	}

	// 創建 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 用密鑰簽名 JWT token
	signedToken, err := token.SignedString(JWTSecret)
	if err != nil {
		log.Println("JWT signing error:", err)
		return "", err
	}
	log.Println("Generated JWT token:", signedToken)
	return signedToken, nil
}

// Register
func Register(c *gin.Context, db *gorm.DB) {
	var userInput struct {
		Username string `json:"username" db:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 檢查 email 是否已經被註冊
	existingUser, err := models.GetUserByEmail(db, userInput.Email)
	if err == nil && existingUser.Email != "" {
		// 如果 email 已經存在，返回錯誤
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	// 密碼hash處理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	user := models.User{
		Username:     userInput.Username,
		Email:        userInput.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := models.CreateUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	//c.JSON(http.StatusCreated, user)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    user,
	})
}

// login
func Login(c *gin.Context, db *gorm.DB) {

	var userInput struct {
		Username string `json:"username" db:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := models.GetUserByUsername(db, userInput.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// 檢查密碼

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userInput.Password)); err != nil {
		log.Println("Password comparison failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 密碼正確，生成 JWT
	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// 返回成功訊息和 JWT token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})

	//c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
