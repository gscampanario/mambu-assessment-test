package service

import (
	"context"
	"errors"
	"testing"

	"com.github.gscampanario/mambu-assessment-test/domain/client/model"
	"com.github.gscampanario/mambu-assessment-test/mock"
	"github.com/golang/mock/gomock"
)

func TestClientService_Insert(t *testing.T) {
	service := ClientService{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var tests = []struct {
		name      string
		input     model.Client
		wantError bool
		mock      func(client model.Client)
	}{
		{
			name: "#1 - Insert client successfully",
			input: model.Client{
				FirstName: "Jane",
				LastName:  "Smith",
			},
			wantError: false,
			mock: func(client model.Client) {
				mockRepo := mock.NewMockIClientRepository(ctrl)
				mockRepo.EXPECT().
					Insert(gomock.Any(), client).
					Return(nil)

				service.ClientRepository = mockRepo
			},
		},
		{
			name: "#2 - Insert client with error",
			input: model.Client{
				FirstName: "Jane",
				LastName:  "Smith",
			},
			wantError: true,
			mock: func(client model.Client) {
				mockRepo := mock.NewMockIClientRepository(ctrl)
				mockRepo.EXPECT().
					Insert(gomock.Any(), client).
					Return(errors.New("mocked error"))

				service.ClientRepository = mockRepo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)
			err := service.Insert(context.Background(), tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("wantError mismatch, should be %v, got %v", tt.wantError, !tt.wantError)
			}
		})
	}
}
