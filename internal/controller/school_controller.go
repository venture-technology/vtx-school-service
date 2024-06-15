package controllers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-school/config"
	"github.com/venture-technology/vtx-school/internal/service"
	"github.com/venture-technology/vtx-school/types"
	"github.com/venture-technology/vtx-school/utils"
)

type ClaimsSchool struct {
	CNPJ string `json:"cnpj"`
	jwt.StandardClaims
}

type SchoolController struct {
	schoolservice *service.SchoolService
}

func NewSchoolController(schoolservice *service.SchoolService) *SchoolController {
	return &SchoolController{
		schoolservice: schoolservice,
	}
}

func (ct *SchoolController) RegisterRoutes(router *gin.Engine) {

	conf := config.Get()

	schoolMiddleware := func(c *gin.Context) {

		secret := []byte(conf.Server.Secret)

		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sem cookie de sessão"})
			c.Abort()
			return
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &ClaimsSchool{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*ClaimsSchool)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("cnpj", claims.CNPJ)
		c.Set("isAuthenticated", true)
		c.Next()

	}

	api := router.Group("api/v1")

	api.GET("/ping", ct.Ping)                                // pingar rota
	api.POST("/school", ct.CreateSchool)                     // criar uma escola
	api.GET("/school/:cnpj", ct.ReadSchool)                  // buscar uma escola em especifico
	api.GET("/school", ct.ReadAllSchools)                    // buscar todas as escolas
	api.PATCH("/school", schoolMiddleware, ct.UpdateSchool)  // atualizar algum dado especifico
	api.DELETE("/school", schoolMiddleware, ct.DeleteSchool) // deletar propria conta
	api.POST("/login/school", ct.AuthSchool)                 // logar como escola
}

func (ct *SchoolController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func (ct *SchoolController) CreateSchool(c *gin.Context) {

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.schoolservice.CreateSchool(c, &input)

	if err != nil {
		log.Printf("error to create school: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating school"})
		return
	}

	log.Print("school create was successful")

	c.JSON(http.StatusCreated, input)

}

func (ct *SchoolController) ReadSchool(c *gin.Context) {

	cnpj := c.Param("cnpj")

	school, err := ct.schoolservice.ReadSchool(c, &cnpj)

	if err != nil {
		log.Printf("error while found school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"school": school})

}

func (ct *SchoolController) ReadAllSchools(c *gin.Context) {

	schools, err := ct.schoolservice.ReadAllSchools(c)

	if err != nil {
		log.Printf("error while found schools: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "schools don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"schools": schools})

}

func (ct *SchoolController) UpdateSchool(c *gin.Context) {

	cnpjInterface, err := ct.schoolservice.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cnpj of cookie don't found"})
		return
	}

	cnpj, err := utils.InterfaceToString(cnpjInterface)

	log.Printf("trying change your infos --> %v", cnpj)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the value isn't string"})
		return
	}

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	input.CNPJ = *cnpj

	err = ct.schoolservice.UpdateSchool(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error at update"})
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (ct *SchoolController) DeleteSchool(c *gin.Context) {

	cnpjInterface, err := ct.schoolservice.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cnpj of cookie don't found"})
		return
	}

	cnpj, err := utils.InterfaceToString(cnpjInterface)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the value isn't string"})
		return
	}

	err = ct.schoolservice.DeleteSchool(c, cnpj)

	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cnpj)

	c.JSON(http.StatusOK, gin.H{"message": "school deleted w successfully"})

}

func (ct *SchoolController) AuthSchool(c *gin.Context) {

	var input types.School

	log.Printf("%s --> doing login --> %s", input.Email)

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "invalid body content"})
		return
	}

	school, err := ct.schoolservice.AuthSchool(c, &input)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := ct.schoolservice.CreateTokenJWTSchool(c, school)

	log.Printf("token returned --> %v", jwt)

	if err != nil {
		log.Printf("error to create jwt token: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to create jwt token"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", c.Request.Host, false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"school": school,
		"token":  jwt,
	})
}
