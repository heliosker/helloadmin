package handler

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"helloadmin/api"
	"helloadmin/internal/handler"
	"helloadmin/internal/middleware"
	"helloadmin/internal/model"
	"helloadmin/test/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var roleId int64 = 1

func TestRoleHandler_StoreRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := api.RoleCreateRequest{
		Name:     "test role",
		Slug:     "test",
		Describe: "this is test role",
	}
	mockRoleService := mock_service.NewMockRoleService(ctrl)
	mockRoleService.EXPECT().CreateRole(gomock.Any(), &params).Return(nil)

	roleHandler := handler.NewRoleHandler(hdl, mockRoleService)
	router.Use(middleware.NoStrictAuth(jwt, logger))
	router.POST("/role", roleHandler.StoreRole)

	paramsJson, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", "/role", bytes.NewBuffer(paramsJson))
	req.Header.Set("Authorization", "Bearer "+genToken(t))

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}

func TestRoleHandler_GetRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	params := api.RoleFindRequest{
		Name: "",
		Slug: "",
		Meta: api.Meta{
			Page: 1,
			Size: 10,
		},
	}
	mockRoleService := mock_service.NewMockRoleService(ctrl)
	mockRoleService.EXPECT().SearchRole(gomock.Any(), &params).Return(&[]model.Role{}, nil)

	roleHandler := handler.NewRoleHandler(hdl, mockRoleService)
	router.Use(middleware.NoStrictAuth(jwt, logger))
	router.GET("/role", roleHandler.GetRole)

	req, _ := http.NewRequest("GET", "/role"+"?page=1&size=10", nil)
	req.Header.Set("Authorization", "Bearer "+genToken(t))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}

//func TestRoleHandler_ShowRole(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRoleService := mock_service.NewMockRoleService(ctrl)
//	mockRoleService.EXPECT().GetRoleById(gomock.Any(), roleId).Return(&api.RoleResponseItem{
//		Slug:      "test",
//		Describe:  "this is test role",
//		CreatedAt: "2023-12-27 19:01:00",
//		UpdateAt:  "2023-12-27 19:01:00",
//		Name:      "test role",
//	}, nil)
//
//	roleHandler := handler.NewRoleHandler(hdl, mockRoleService)
//	//router.Use(middleware.NoStrictAuth(jwt, logger))
//	router.GET("/role/:id", roleHandler.ShowRole)
//
//	req, _ := http.NewRequest("GET", "/role/1", nil)
//	req.Header.Set("Authorization", "Bearer "+genToken(t))
//
//	resp := httptest.NewRecorder()
//	router.ServeHTTP(resp, req)
//	assert.Equal(t, resp.Code, http.StatusOK)
//}

//
//func TestRoleHandler_UpdateRole(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	params := api.RoleUpdateRequest{
//		Name:     "test role update",
//		Slug:     "test",
//		Describe: "this is test role",
//	}
//	mockRoleService := mock_service.NewMockRoleService(ctrl)
//	mockRoleService.EXPECT().UpdateRole(gomock.Any(), roleId, &params).Return(nil)
//
//	roleHandler := handler.NewRoleHandler(hdl, mockRoleService)
//	router.Use(middleware.NoStrictAuth(jwt, logger))
//	router.PUT("/role/1", roleHandler.UpdateRole)
//
//	paramsJson, _ := json.Marshal(params)
//	req, _ := http.NewRequest("PUT", "/role/1", bytes.NewBuffer(paramsJson))
//	req.Header.Set("Authorization", "Bearer "+genToken(t))
//	resp := httptest.NewRecorder()
//	router.ServeHTTP(resp, req)
//	assert.Equal(t, resp.Code, http.StatusOK)
//}

//func TestRoleHandler_DeleteRole(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRoleService := mock_service.NewMockRoleService(ctrl)
//	mockRoleService.EXPECT().DeleteRole(gomock.Any(), roleId).Return(nil)
//	//router.Use(middleware.NoStrictAuth(jwt, logger))
//	//router.DELETE("/role/:id", handler.NewRoleHandler(hdl, mockRoleService).DeleteRole)
//
//	req, _ := http.NewRequest("DELETE", "/role/1", nil)
//	req.Header.Set("Authorization", "Bearer "+genToken(t))
//	resp := httptest.NewRecorder()
//	router.ServeHTTP(resp, req)
//	assert.Equal(t, resp.Code, http.StatusOK)
//}
