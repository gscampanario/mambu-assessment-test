package mapper

import (
	"testing"

	"com.github.gscampanario/mambu-assessment-test/application/dto/client"
	models "com.github.gscampanario/mambu-assessment-test/domain/client/model"
)

func TestInsertClientMapper(t *testing.T) {
	var tests = []struct {
		name  string
		input client.InsertClientRequestDTO
		want  models.Client
	}{
		{
			name: "#1 - Map InsertClientRequestDTO to Client",
			input: client.InsertClientRequestDTO{
				FirstName: "John",
				LastName:  "Doe",
			},
			want: models.Client{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper := ClientMapper{}
			got := mapper.MapInsertClientRequestDTOToClient(tt.input)
			if got != tt.want {
				t.Errorf("MapInsertClientRequestDTOToClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
