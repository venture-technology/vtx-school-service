package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-school/config"
	"github.com/venture-technology/vtx-school/internal/repository"
	"github.com/venture-technology/vtx-school/types"
	"github.com/venture-technology/vtx-school/utils"
)

type SchoolService struct {
	schoolrepository repository.SchoolRepositoryInterface
	kafkarepository  repository.IKafkaRepository
}

func NewSchoolService(repo repository.SchoolRepositoryInterface, producer repository.IKafkaRepository) *SchoolService {
	return &SchoolService{schoolrepository: repo, kafkarepository: producer}
}

func (s *SchoolService) CreateSchool(ctx context.Context, school *types.School) error {
	log.Printf("input received to create school -> name: %s, cnpj: %s, email: %s", school.Name, school.CNPJ, school.Email)
	school.Password = utils.HashPassword(school.Password)

	return s.schoolrepository.CreateSchool(ctx, school)
}

func (s *SchoolService) ReadSchool(ctx context.Context, cnpj *string) (*types.School, error) {
	log.Printf("param read school -> cnpj: %s", cnpj)
	return s.schoolrepository.ReadSchool(ctx, cnpj)
}

func (s *SchoolService) ReadAllSchools(ctx context.Context) ([]types.School, error) {
	return s.schoolrepository.ReadAllSchools(ctx)
}

func (s *SchoolService) UpdateSchool(ctx context.Context, school *types.School) error {
	log.Printf("input received to update school -> name: %s, cnpj: %s, email: %s", school.Name, school.CNPJ, school.Email)
	return s.schoolrepository.UpdateSchool(ctx, school)
}

func (s *SchoolService) DeleteSchool(ctx context.Context, cnpj *string) error {
	log.Printf("trying delete your infos --> %v", cnpj)
	return s.schoolrepository.DeleteSchool(ctx, cnpj)
}

func (s *SchoolService) AuthSchool(ctx context.Context, school *types.School) (*types.School, error) {
	school.Password = utils.HashPassword((school.Password))
	return s.schoolrepository.AuthSchool(ctx, school)
}

func (s *SchoolService) ParserJwtSchool(ctx *gin.Context) (interface{}, error) {

	cnpj, found := ctx.Get("cnpj")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cnpj, nil

}

func (s *SchoolService) CreateTokenJWTSchool(ctx context.Context, school *types.School) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cnpj": school.CNPJ,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil

}
