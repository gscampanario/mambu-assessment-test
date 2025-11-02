package mapper

import (
	"com.github.gscampanario/mambu-assessment-test/application/dto/client"
	models "com.github.gscampanario/mambu-assessment-test/domain/client"
)

type ClientMapper struct{}

func NewClientMapper() *ClientMapper {
	return &ClientMapper{}
}

// MapInsertClientRequestDTOToClient maps InsertClientRequestDTO to Client domain model
func (m *ClientMapper) MapInsertClientRequestDTOToClient(dto client.InsertClientRequestDTO) models.Client {
	return models.Client{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
}
