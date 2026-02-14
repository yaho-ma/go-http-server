package calculationService

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (MyCalculation, error)
	GetAllCalculations() ([]MyCalculation, error)
	GetCalculationById(id string) (MyCalculation, error)
	UpdateCalculation(id, expression string) (MyCalculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationRepository
}

/////////////////////////////////////////////////////////////////////////////////////

func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), err
}

func NewCalculationService(r CalculationRepository) CalculationService { // method with a receiver
	return &calcService{repo: r} // Create a new calcService struct and set its field repo to the value repo.”
}

func (s *calcService) CreateCalculation(expression string) (MyCalculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return MyCalculation{}, err
	}
	calc := MyCalculation{
		Expression: expression,
		Id:         uuid.NewString(),
		Result:     result,
	}
	if err := s.repo.CreateNewCalcInDB(calc); err != nil {
		return MyCalculation{}, err
	}
	return calc, err
}

func (s *calcService) GetAllCalculations() ([]MyCalculation, error) {
	return s.repo.GetAllCalculationsFromDB()
}

func (s *calcService) GetCalculationById(id string) (MyCalculation, error) {
	return s.repo.GetCalculationByIdFromDB(id)
}

func (s *calcService) UpdateCalculation(id, expression string) (MyCalculation, error) {
	calc, err := s.GetCalculationById(id)
	if err != nil {
		return MyCalculation{}, err
	}
	result, err := s.calculateExpression(expression)
	if err != nil {
		return MyCalculation{}, err
	}
	// update fields in the structure
	calc.Result = result
	calc.Expression = expression

	if err := s.repo.UpdateCalculationByIdInDB(calc); err != nil {
		return MyCalculation{}, err
	}
	return calc, nil

}

func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculationInDB(id)
}
