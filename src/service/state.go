package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hegade/go_address_API/models"
	"gorm.io/gorm"
)

type StateService interface {
	Create(StateRequest) (string, error)
	Get(string) (*StateRequest, error)
	Update(string, StateRequest) (*StateRequest, error)
	Delete(string) (string, error)
	List(string, string) ([]StateRequest, error)
	DeleteAll() (string, error)
}
type stateService struct {
	conn *gorm.DB
	ctx  *gin.Context
}

type StateRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type errorString struct {
	errorstring string
}

func (e *errorString) Error() string {
	return e.errorstring
}
func throwError(text string) error {
	return &errorString{
		text,
	}
}
func NewStateService(conn *gorm.DB) StateService {
	return &stateService{
		conn: conn,
	}
}

func (s *stateService) Create(req StateRequest) (string, error) {
	req.Name = strings.Trim(req.Name, " ")
	incomingData := models.Mst_state{
		Name: req.Name,
	}
	fmt.Println(req)
	if req.Name == "" {
		return "", throwError("State name cannot be empty")
	}
	result := s.conn.Create(&incomingData)
	fmt.Println(result)
	return "State Inserted", nil
}

func (s *stateService) List(page string, size string) ([]StateRequest, error) {
	stateList := []StateRequest{}
	// if page == "" || size == "" {
	// 	return stateList, throwError("Page and Size cannot be empty")
	// }

	// totalCount, _ := strconv.ParseInt(size, 6, 12)
	// toStart, _ := strconv.ParseInt(page, 6, 12)
	// offSet := (toStart - 1) * totalCount
	// limit := toStart * totalCount
	s.conn.Find(&models.Mst_state{}).Scan(&stateList)
	fmt.Println(stateList)

	return stateList, nil
}

func (s *stateService) Get(ID string) (*StateRequest, error) {
	_, err := strconv.ParseInt(ID, 0, 8)
	data := models.Mst_state{}
	if err != nil {
		return nil, throwError("Invalid ID")
	}
	result := map[string]interface{}{}

	stu := s.conn.Model(&data).Where("id = ?", ID).Scan(&result)
	fmt.Println(result)
	if stu.RowsAffected == 0 {
		return nil, throwError("State not found")

	}
	state := StateRequest{}
	state.ID = result["id"].(int64)
	state.Name = result["name"].(string)
	return &state, nil

}

func (s *stateService) Update(ID string, req StateRequest) (*StateRequest, error) {
	id, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		return nil, throwError("Invalid ID")
	}
	req.Name = strings.Trim(req.Name, " ")
	if req.Name == "" {
		return nil, throwError("State Name cannot be empty")
	}

	stateToUpdate := models.Mst_state{
		ID: id,
	}
	res := s.conn.Model(&stateToUpdate).Update("name", req.Name)
	if res.RowsAffected == 0 {
		return nil, throwError("State not found")
	}

	return &req, nil
}

func (s *stateService) Delete(ID string) (string, error) {
	id, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		return "", throwError("Invalid ID")
	}
	stateToDelete := models.Mst_state{
		ID: id,
	}
	res := s.conn.Delete(&stateToDelete)
	if res.RowsAffected == 0 {
		return "", throwError("State not found")
	}

	return "deleted", nil
}

func (s *stateService) DeleteAll() (string, error) {
	res := s.conn.Delete(&models.Mst_state{})
	if res.RowsAffected == 0 {
		return "", throwError("State not found")
	}

	return "All states deleted", nil
}
