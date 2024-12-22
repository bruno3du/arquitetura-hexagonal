package cli_test

import (
	"testing"

	"github.com/bruno3du/hexagonal/adapters/cli"
	"github.com/bruno3du/hexagonal/application"
	mock_application "github.com/bruno3du/hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.5
	productStatus := application.ENABLED
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)

	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	result, err := cli.Run(service, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, "Product: "+productName+" has been created with success", result)

	result, err = cli.Run(service, "get", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, "Product: "+productName+" has been found with success", result)

	result, err = cli.Run(service, "enable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, "Product: "+productName+" has been enabled with success", result)

	result, err = cli.Run(service, "disable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, "Product: "+productName+" has been disabled with success", result)
}
