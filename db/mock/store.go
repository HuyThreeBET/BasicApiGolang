// Code generated by MockGen. DO NOT EDIT.
// Source: simple_shop/db/sqlc (interfaces: Store)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	sqlc "simple_shop/db/sqlc"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddBankAccountBalance mocks base method.
func (m *MockStore) AddBankAccountBalance(arg0 context.Context, arg1 sqlc.AddBankAccountBalanceParams) (sqlc.BankAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBankAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(sqlc.BankAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBankAccountBalance indicates an expected call of AddBankAccountBalance.
func (mr *MockStoreMockRecorder) AddBankAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBankAccountBalance", reflect.TypeOf((*MockStore)(nil).AddBankAccountBalance), arg0, arg1)
}

// CreateAdmin mocks base method.
func (m *MockStore) CreateAdmin(arg0 context.Context, arg1 sqlc.CreateAdminParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockStoreMockRecorder) CreateAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockStore)(nil).CreateAdmin), arg0, arg1)
}

// CreateBankAccount mocks base method.
func (m *MockStore) CreateBankAccount(arg0 context.Context, arg1 sqlc.CreateBankAccountParams) (sqlc.BankAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBankAccount", arg0, arg1)
	ret0, _ := ret[0].(sqlc.BankAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBankAccount indicates an expected call of CreateBankAccount.
func (mr *MockStoreMockRecorder) CreateBankAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBankAccount", reflect.TypeOf((*MockStore)(nil).CreateBankAccount), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 sqlc.CreateProductParams) (sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreatePuschaseHistory mocks base method.
func (m *MockStore) CreatePuschaseHistory(arg0 context.Context, arg1 sqlc.CreatePuschaseHistoryParams) (sqlc.PurchaseHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePuschaseHistory", arg0, arg1)
	ret0, _ := ret[0].(sqlc.PurchaseHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePuschaseHistory indicates an expected call of CreatePuschaseHistory.
func (mr *MockStoreMockRecorder) CreatePuschaseHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePuschaseHistory", reflect.TypeOf((*MockStore)(nil).CreatePuschaseHistory), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 sqlc.CreateUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteBankAccount mocks base method.
func (m *MockStore) DeleteBankAccount(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBankAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBankAccount indicates an expected call of DeleteBankAccount.
func (mr *MockStoreMockRecorder) DeleteBankAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBankAccount", reflect.TypeOf((*MockStore)(nil).DeleteBankAccount), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockStore) DeleteProduct(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetBankAccountByUserNameAndCurrency mocks base method.
func (m *MockStore) GetBankAccountByUserNameAndCurrency(arg0 context.Context, arg1 sqlc.GetBankAccountByUserNameAndCurrencyParams) (sqlc.BankAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBankAccountByUserNameAndCurrency", arg0, arg1)
	ret0, _ := ret[0].(sqlc.BankAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBankAccountByUserNameAndCurrency indicates an expected call of GetBankAccountByUserNameAndCurrency.
func (mr *MockStoreMockRecorder) GetBankAccountByUserNameAndCurrency(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBankAccountByUserNameAndCurrency", reflect.TypeOf((*MockStore)(nil).GetBankAccountByUserNameAndCurrency), arg0, arg1)
}

// GetBankAccountByUserNameAndCurrencyForUpdate mocks base method.
func (m *MockStore) GetBankAccountByUserNameAndCurrencyForUpdate(arg0 context.Context, arg1 sqlc.GetBankAccountByUserNameAndCurrencyForUpdateParams) (sqlc.BankAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBankAccountByUserNameAndCurrencyForUpdate", arg0, arg1)
	ret0, _ := ret[0].(sqlc.BankAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBankAccountByUserNameAndCurrencyForUpdate indicates an expected call of GetBankAccountByUserNameAndCurrencyForUpdate.
func (mr *MockStoreMockRecorder) GetBankAccountByUserNameAndCurrencyForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBankAccountByUserNameAndCurrencyForUpdate", reflect.TypeOf((*MockStore)(nil).GetBankAccountByUserNameAndCurrencyForUpdate), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStore) GetProduct(arg0 context.Context, arg1 int32) (sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStore)(nil).GetProduct), arg0, arg1)
}

// GetPurchaseHistory mocks base method.
func (m *MockStore) GetPurchaseHistory(arg0 context.Context, arg1 int32) (sqlc.PurchaseHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPurchaseHistory", arg0, arg1)
	ret0, _ := ret[0].(sqlc.PurchaseHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPurchaseHistory indicates an expected call of GetPurchaseHistory.
func (mr *MockStoreMockRecorder) GetPurchaseHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPurchaseHistory", reflect.TypeOf((*MockStore)(nil).GetPurchaseHistory), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListBankAccountsByUsername mocks base method.
func (m *MockStore) ListBankAccountsByUsername(arg0 context.Context, arg1 string) ([]sqlc.BankAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBankAccountsByUsername", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.BankAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBankAccountsByUsername indicates an expected call of ListBankAccountsByUsername.
func (mr *MockStoreMockRecorder) ListBankAccountsByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBankAccountsByUsername", reflect.TypeOf((*MockStore)(nil).ListBankAccountsByUsername), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStore) ListProducts(arg0 context.Context) ([]sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0)
	ret0, _ := ret[0].([]sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreMockRecorder) ListProducts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStore)(nil).ListProducts), arg0)
}

// ListPuschaseHistories mocks base method.
func (m *MockStore) ListPuschaseHistories(arg0 context.Context) ([]sqlc.PurchaseHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPuschaseHistories", arg0)
	ret0, _ := ret[0].([]sqlc.PurchaseHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPuschaseHistories indicates an expected call of ListPuschaseHistories.
func (mr *MockStoreMockRecorder) ListPuschaseHistories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPuschaseHistories", reflect.TypeOf((*MockStore)(nil).ListPuschaseHistories), arg0)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context, arg1 sqlc.ListUsersParams) ([]sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0, arg1)
}

// PurchaseTransaction mocks base method.
func (m *MockStore) PurchaseTransaction(arg0 context.Context, arg1 sqlc.PurchaseTransactionPagrams) (sqlc.PurchaseTransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurchaseTransaction", arg0, arg1)
	ret0, _ := ret[0].(sqlc.PurchaseTransactionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurchaseTransaction indicates an expected call of PurchaseTransaction.
func (mr *MockStoreMockRecorder) PurchaseTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurchaseTransaction", reflect.TypeOf((*MockStore)(nil).PurchaseTransaction), arg0, arg1)
}

// UpdateHashedPasswordOfUser mocks base method.
func (m *MockStore) UpdateHashedPasswordOfUser(arg0 context.Context, arg1 sqlc.UpdateHashedPasswordOfUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHashedPasswordOfUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHashedPasswordOfUser indicates an expected call of UpdateHashedPasswordOfUser.
func (mr *MockStoreMockRecorder) UpdateHashedPasswordOfUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHashedPasswordOfUser", reflect.TypeOf((*MockStore)(nil).UpdateHashedPasswordOfUser), arg0, arg1)
}

// UpdateQuantityOfProduct mocks base method.
func (m *MockStore) UpdateQuantityOfProduct(arg0 context.Context, arg1 sqlc.UpdateQuantityOfProductParams) (sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantityOfProduct", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQuantityOfProduct indicates an expected call of UpdateQuantityOfProduct.
func (mr *MockStoreMockRecorder) UpdateQuantityOfProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantityOfProduct", reflect.TypeOf((*MockStore)(nil).UpdateQuantityOfProduct), arg0, arg1)
}
