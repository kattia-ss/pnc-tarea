package csv

import (
	"encoding/csv"
	"errors"
	"layersapi/entities"
	"os"
	"time"
)

type UserRepository struct{}

// Implementa UserRepository con una estructura en CSV. Ideal para pruebas.

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetAll() ([]entities.User, error) {
	file, err := os.Open("data.csv")
	if err != nil {
		return []entities.User{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return []entities.User{}, err
	}

	var result []entities.User

	for i, record := range records {
		if i == 0 {
			continue
		}

		createdAt, _ := time.Parse(time.RFC3339, record[3])
		updatedAt, _ := time.Parse(time.RFC3339, record[4])
		meta := entities.Metadata{
			CreatedAt: createdAt.String(),
			UpdatedAt: updatedAt.String(),
			CreatedBy: record[5],
			UpdatedBy: record[6],
		}
		result = append(result, entities.NewUser(record[0], record[1], record[2], meta))
	}

	return result, nil
}

func (u UserRepository) GetById(id string) (entities.User, error) {
	file, err := os.Open("data.csv")
	if err != nil {
		return entities.User{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return entities.User{}, err
	}

	for i, record := range records {
		if i == 0 {
			continue
		} else if record[0] == id {

			createdAt, _ := time.Parse(time.RFC3339, record[3])
			updatedAt, _ := time.Parse(time.RFC3339, record[4])
			meta := entities.Metadata{
				CreatedAt: createdAt.String(),
				UpdatedAt: updatedAt.String(),
				CreatedBy: record[5],
				UpdatedBy: record[6],
			}
			return entities.NewUser(record[0], record[1], record[2], meta), nil
		}

	}

	return entities.User{}, errors.New("user not found")
}

func (u UserRepository) Create(user entities.User) error {
	file, err := os.Open("data.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	newUser := []string{
		user.Id,
		user.Name,
		user.Email,
		user.Metadata.CreatedAt,
		user.Metadata.UpdatedAt,
		"webapp",
		"webapp",
	}

	if err := writer.Write(newUser); err != nil {
		return err
	}

	return nil
}

// completar update
func (u *UserRepository) Update(id, name, email string) error {
	file, err := os.Open("data.csv")
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		file.Close()
		return err
	}
	file.Close()

	userFound := false
	for i, record := range records {
		if i > 0 && record[0] == id {
			records[i][1] = name
			records[i][2] = email
			records[i][4] = time.Now().Format(time.RFC3339) // Actualizar UpdatedAt
			userFound = true
			break
		}
	}

	if !userFound {
		return errors.New("user not found")
	}

	// Reescribir todo el archivo
	outputFile, err := os.Create("data.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// hacer delete

func (u *UserRepository) Delete(id string) error {
	file, err := os.Open("data.csv")
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		file.Close()
		return err
	}
	file.Close()

	userFound := false
	var newRecords [][]string

	// Mantener el encabezado y filtrar el usuario que queremos eliminar
	newRecords = append(newRecords, records[0])

	for i, record := range records {
		if i > 0 {
			if record[0] == id {
				userFound = true
				// No añadimos este registro a newRecords
			} else {
				newRecords = append(newRecords, record)
			}
		}
	}

	if !userFound {
		return errors.New("user not found")
	}

	// Reescribir todo el archivo sin el usuario eliminado
	outputFile, err := os.Create("data.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, record := range newRecords {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
