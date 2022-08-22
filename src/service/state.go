package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type StateService interface {
	Create(StateRequest) (string, error)
	Get(string) (*StateRequest, error)
	Update(string, StateRequest) (*StateRequest, error)
	Delete(string) (string, error)
	List(string, string) ([]StateRequest, error)
}
type stateService struct {
	conn *pgxpool.Pool
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
func errorData(text string) error {
	return &errorString{
		errorstring: text,
	}
}
func NewStateService(conn *pgxpool.Pool) StateService {
	return &stateService{
		conn: conn,
	}
}

func (s *stateService) Create(req StateRequest) (string, error) {
	req.Name = strings.Trim(req.Name, " ")
	if req.Name == "" {
		return "", errorData("State name cannot be empty")
	}
	cmdTag, err := s.conn.Exec(context.Background(), "insert into mst_state (name) values('"+req.Name+"')")
	if err != nil {
		return "", err
	}
	fmt.Println(cmdTag)
	return "State Inserted", nil
}

func (s *stateService) List(page string, size string) ([]StateRequest, error) {
	stateList := []StateRequest{}
	Rows, err := s.conn.Query(context.Background(), "select * from mst_state")
	if err != nil {
		return nil, err
	}
	defer Rows.Close()
	for Rows.Next() {
		stateArray := StateRequest{}
		err := Rows.Scan(&stateArray.ID, &stateArray.Name)
		if err != nil {
			return nil, err
		}
		stateList = append(stateList, stateArray)
	}
	if err != nil {
		return nil, err
	}
	return stateList, nil
}

func (s *stateService) Get(ID string) (*StateRequest, error) {
	_, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		return nil, errorData("Invalid ID")
	}
	Row, err := s.conn.Query(context.Background(), "select * from mst_state where id="+ID)
	if err != nil {
		return nil, err
	}
	defer Row.Close()
	if Row.CommandTag().RowsAffected() == 0 {
		return nil, errorData("ID Not Found")
	}
	stateData := StateRequest{}
	for Row.Next() {
		err = Row.Scan(&stateData.ID, &stateData.Name)
		if err != nil {
			return nil, err
		}
	}
	return &stateData, nil
}

func (s *stateService) Update(ID string, req StateRequest) (*StateRequest, error) {
	_, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		return nil, errorData("Invalid ID")
	}
	req.Name = strings.Trim(req.Name, " ")
	if req.Name == "" {
		return nil, errorData("State Name cannot be empty")
	}
	cmdTag, err := s.conn.Exec(context.Background(), "update mst_state set name='"+req.Name+"' where id="+ID)
	if err != nil {
		return nil, err
	}
	if cmdTag.RowsAffected() == 0 {
		return nil, errorData("ID Not Found")
	}
	req.ID, _ = strconv.ParseInt(ID, 0, 8)
	return &req, nil
}

func (s *stateService) Delete(ID string) (string, error) {
	_, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		return "", errorData("Invalid ID")
	}
	cmdTag, err := s.conn.Exec(context.Background(), "delete from mst_state where id="+ID)
	if err != nil {
		return "", err
	}
	if cmdTag.RowsAffected() == 0 {
		return "", errorData("ID Not Found")
	}
	return "deleted", nil
}
