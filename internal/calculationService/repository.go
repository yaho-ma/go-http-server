package calculationService

import "gorm.io/gorm"

type CalculationRepository interface {
	CreateNewCalcInDB(calc MyCalculation) error
	GetAllCalculationsFromDB() ([]MyCalculation, error)
	GetCalculationByIdFromDB(id string) (MyCalculation, error)
	UpdateCalculationByIdInDB(calc MyCalculation) error
	DeleteCalculationInDB(id string) error
}

type calcRepository struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calcRepository{db: db}

}

// implementation of methods
func (r *calcRepository) CreateNewCalcInDB(calc MyCalculation) error {
	return r.db.Create(&calc).Error
}

func (r *calcRepository) GetAllCalculationsFromDB() ([]MyCalculation, error) {
	var calculations []MyCalculation

	err := r.db.Find(&calculations).Error
	return calculations, err
}

func (r *calcRepository) GetCalculationByIdFromDB(id string) (MyCalculation, error) {
	var oneCalculation MyCalculation

	err := r.db.First(&oneCalculation, "id = ?", id).Error
	return oneCalculation, err
}

func (r *calcRepository) UpdateCalculationByIdInDB(calc MyCalculation) error {
	return r.db.Save(&calc).Error
}

func (r *calcRepository) DeleteCalculationInDB(id string) error {
	return r.db.Delete(&MyCalculation{}, "id = ?", id).Error
}
