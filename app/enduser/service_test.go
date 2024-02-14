package enduser_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
// 	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/enduser"
// 	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockUserRepo struct {
// 	mock.Mock
// }

// // func (m *MockUserRepo) BeginTx(ctx context.Context) (*sql.Tx, error) {
// // 	args := m.Called(ctx)
// // 	return args.Get(0).(*sql.Tx), args.Error(1)
// // }

// // func (m *MockUserRepo) GetLoginDetails() (map[string]string, error) {
// // 	args := m.Called()
// // 	return args.Get(0).(map[string]string), args.Error(1)
// // }

// // func (m *MockUserRepo) TokenDetails(username string) (int, string, error) {
// // 	args := m.Called(username)
// // 	return args.Int(0), args.String(1), args.Error(2)
// // }

// // func (m *MockUserRepo) HandleTransaction(ctx context.Context, tx repository.Transaction, err error) error {
// // 	args := m.Called(ctx, tx, err)
// // 	return args.Error(0)
// // }

// // func (m *MockUserRepo) AddUser(req dto.CreateUser) (dto.Response, error) {
// // 	args := m.Called(req)
// // 	return args.Get(0).(dto.Response), args.Error(1)
// // }

// // func (m *MockUserRepo) UpdateUser(req dto.UpdateUser, user_id int) (dto.UpdateUser, error) {
// // 	args := m.Called(req, user_id)
// // 	return args.Get(0).(dto.UpdateUser), args.Error(1)
// // }

// func TestAuthenticate(t *testing.T) {
// 	mockUserRepo := new(MockUserRepo)
// 	service := enduser.NewService(mockUserRepo)

// 	// Test case 1: Valid token
// 	validToken := "valid_token"
// 	expectedUserID := 1
// 	expectedUsername := "test_user"
// 	mockUserRepo.On("TokenDetails", "test_user").Return(expectedUserID, "test_role", nil)
// 	_, username, err := service.Authenticate(validToken)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUserID, 1)
// 	assert.Equal(t, expectedUsername, username)

// 	// Test case 2: Invalid token
// 	invalidToken := "invalid_token"
// 	mockUserRepo.On("TokenDetails", "test_user").Return(0, "", errors.New("error fetching token details"))
// 	_, _, err = service.Authenticate(invalidToken)
// 	assert.Error(t, err)
// }

// func TestCreateLogin(t *testing.T) {
// 	mockUserRepo := new(MockUserRepo)
// 	service := enduser.NewService(mockUserRepo)

// 	ctx := context.Background()

// 	// Test case 1: Valid login credentials
// 	validRequest := dto.CreateLoginRequest{
// 		Username: "test_user",
// 		Password: "test_password",
// 	}
// 	mockUserRepo.On("GetLoginDetails").Return(map[string]string{
// 		"test_user": "$2a$10$ajV0/wlS64E7vGh6ROhwOe0TlIYfxwmEJf2ZwYJf1H4H22VdZ17ze", // hashed password
// 	}, nil)
// 	mockUserRepo.On("TokenDetails", "test_user").Return(1, "test_role", nil)
// 	mockUserRepo.On("HandleTransaction", ctx, mock.Anything, nil).Return(nil)
// 	token, err := service.CreateLogin(ctx, validRequest)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, token)

// 	// Test case 2: Invalid login credentials
// 	invalidRequest := dto.CreateLoginRequest{
// 		Username: "test_user",
// 		Password: "wrong_password",
// 	}
// 	mockUserRepo.On("GetLoginDetails").Return(map[string]string{
// 		"test_user": "$2a$10$ajV0/wlS64E7vGh6ROhwOe0TlIYfxwmEJf2ZwYJf1H4H22VdZ17ze", // hashed password
// 	}, nil)
// 	_, err = service.CreateLogin(ctx, invalidRequest)
// 	assert.Error(t, err)
// }

// func TestCreateSignup(t *testing.T) {
// 	mockUserRepo := new(MockUserRepo)
// 	service := enduser.NewService(mockUserRepo)

// 	ctx := context.Background()

// 	// Test case 1: Valid signup request
// 	validRequest := dto.CreateUser{
// 		Name:     "demo",
// 		Address:  "PUNE",
// 		Email:    "abc@gmail.com",
// 		Password: "demoPass!1",
// 		Mobile:   "9598956231",
// 		Role:     "Customer",
// 	}
// 	mockUserRepo.On("BeginTx", ctx).Return(&repository.MockTx{}, nil)
// 	mockUserRepo.On("AddUser", validRequest).Return(dto.Response{}, nil)
// 	mockUserRepo.On("HandleTransaction", ctx, mock.Anything, nil).Return(nil)
// 	response, err := service.CreateSignup(ctx, validRequest)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response)

// 	// Test case 2: Error during signup
// 	mockUserRepo.On("BeginTx", ctx).Return(&repository.MockTx{}, nil)
// 	mockUserRepo.On("AddUser", validRequest).Return(dto.Response{}, errors.New("error adding user"))
// 	mockUserRepo.On("HandleTransaction", ctx, mock.Anything, mock.Anything).Return(nil)
// 	_, err = service.CreateSignup(ctx, validRequest)
// 	assert.Error(t, err)
// }

// func TestUpdateUser(t *testing.T) {
// 	mockUserRepo := new(MockUserRepo)
// 	service := enduser.NewService(mockUserRepo)

// 	ctx := context.Background()
// 	userID := 1

// 	// Test case 1: Valid update request
// 	validRequest := dto.UpdateUser{
// 		Name:     "demo_update",
// 		Address:  "PUNE,MH",
// 		Password: "demoPass!123",
// 		Mobile:   "9595959595",
// 	}
// 	mockUserRepo.On("BeginTx", ctx).Return(&repository.MockTx{}, nil)
// 	mockUserRepo.On("UpdateUser", validRequest, userID).Return(dto.UpdateUser{}, nil)
// 	mockUserRepo.On("HandleTransaction", ctx, mock.Anything, nil).Return(nil)
// 	response, err := service.UpdateUser(ctx, validRequest, userID)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response)

// 	// Test case 2: Error during update
// 	mockUserRepo.On("BeginTx", ctx).Return(&repository.MockTx{}, nil)
// 	mockUserRepo.On("UpdateUser", validRequest, userID).Return(dto.UpdateUser{}, errors.New("error updating user"))
// 	mockUserRepo.On("HandleTransaction", ctx, mock.Anything, mock.Anything).Return(nil)
// 	_, err = service.UpdateUser(ctx, validRequest, userID)
// 	assert.Error(t, err)
// }

// func Test_service_Authenticate(t *testing.T) {
// 	type fields struct {
// 		UserRepo repository.UserStorer
// 	}
// 	type args struct {
// 		tknStr string
// 	}
// 	tests := []struct {
// 		name         string
// 		fields       fields
// 		args         args
// 		wantUser_id  int
// 		wantResponse string
// 		wantErr      bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			us := &service{
// 				UserRepo: tt.fields.UserRepo,
// 			}
// 			gotUser_id, gotResponse, err := us.Authenticate(tt.args.tknStr)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotUser_id != tt.wantUser_id {
// 				t.Errorf("service.Authenticate() gotUser_id = %v, want %v", gotUser_id, tt.wantUser_id)
// 			}
// 			if gotResponse != tt.wantResponse {
// 				t.Errorf("service.Authenticate() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
// 			}
// 		})
// 	}
// }
