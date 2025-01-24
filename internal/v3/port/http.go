package port

import (
	"context"

	"github.com/cromerc/projectBase/internal/v3/port/mapper"

	"github.com/cromerc/projectBase/internal/v3/port/api"
)

//go:generate go run -modfile=../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config/server.yaml ../../../api/v3.yaml
//go:generate go run -modfile=../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config/models.yaml ../../../api/v3.yaml
//go:generate go run -modfile=../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config/client.yaml ../../../api/v3.yaml

type server struct {
	mapper mapper.PortMapper
}

func (s server) FindPetsByStatus(ctx context.Context, request api.FindPetsByStatusRequestObject) (api.FindPetsByStatusResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) FindPetsByTags(ctx context.Context, request api.FindPetsByTagsRequestObject) (api.FindPetsByTagsResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) DeletePet(ctx context.Context, request api.DeletePetRequestObject) (api.DeletePetResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) GetPetByID(ctx context.Context, request api.GetPetByIDRequestObject) (api.GetPetByIDResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) UpdatePetWithForm(ctx context.Context, request api.UpdatePetWithFormRequestObject) (api.UpdatePetWithFormResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) UploadFile(ctx context.Context, request api.UploadFileRequestObject) (api.UploadFileResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) AddPet(ctx context.Context, request api.AddPetRequestObject) (api.AddPetResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) UpdatePet(ctx context.Context, request api.UpdatePetRequestObject) (api.UpdatePetResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) GetInventory(ctx context.Context, request api.GetInventoryRequestObject) (api.GetInventoryResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) PlaceOrder(ctx context.Context, request api.PlaceOrderRequestObject) (api.PlaceOrderResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) DeleteOrder(ctx context.Context, request api.DeleteOrderRequestObject) (api.DeleteOrderResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) GetOrderByID(ctx context.Context, request api.GetOrderByIDRequestObject) (api.GetOrderByIDResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) CreateUser(ctx context.Context, request api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) CreateUsersWithListInput(ctx context.Context, request api.CreateUsersWithListInputRequestObject) (api.CreateUsersWithListInputResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) LoginUser(ctx context.Context, request api.LoginUserRequestObject) (api.LoginUserResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) LogoutUser(ctx context.Context, request api.LogoutUserRequestObject) (api.LogoutUserResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) DeleteUser(ctx context.Context, request api.DeleteUserRequestObject) (api.DeleteUserResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) GetUserByName(ctx context.Context, request api.GetUserByNameRequestObject) (api.GetUserByNameResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) UpdateUser(ctx context.Context, request api.UpdateUserRequestObject) (api.UpdateUserResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func NewServer() api.StrictServerInterface {
	pm := mapper.PortMapper{}

	return server{mapper: pm}
}
