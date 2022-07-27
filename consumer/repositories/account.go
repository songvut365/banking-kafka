package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccounts []BankAccount, err error)
	FindByID(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.AutoMigrate(&BankAccount{})

	return accountRepository{db}
}

func (repository accountRepository) Save(bankAccount BankAccount) error {
	return repository.db.Save(bankAccount).Error
}

func (repository accountRepository) Delete(id string) error {
	return repository.db.Where("id = ?", id).Delete(&BankAccount{}).Error
}

func (repository accountRepository) FindAll() (bankAccounts []BankAccount, err error) {
	err = repository.db.Find(&bankAccounts).Error
	return bankAccounts, err
}

func (repository accountRepository) FindByID(id string) (bankAccount BankAccount, err error) {
	err = repository.db.Where("id = ?", id).Find(&bankAccount).Error
	return bankAccount, err
}
