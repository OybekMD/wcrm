package v1

import (
	"api-gateway/api/handlers/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Security      ApiKeyAuth
// @Summary       Get list of policeis
// @Description   This API get list of policies
// @Tags          casbin
// @Accept        json
// @Produce       json
// @Param         role query string true "Role"
// @Succes        200 {object} models.ListPolePolicyResponse
// @Failure       404 {object} models.Error
// @Failure       500 {object} models.Error
// @Router /v1/rbac/list-role-policies [GET]
func (h *handlerV1) ListPolicies(ctx *gin.Context) {
	role := ctx.Query("role")
	var resp models.ListPolePolicyResponse

	for _, p := range h.enforcer.GetFilteredPolicy(0, role) {
		resp.Policies = append(resp.Policies, &models.Policy{
			Role:     p[0],
			Endpoint: p[1],
			Method:   p[2],
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Security      ApiKeyAuth
// @Summary       Get list of roles
// @Description   This API get list of roles
// @Tags          casbin
// @Accept        json
// @Produce       json
// @Param         limit query int false "limit"
// @Param         offset query int false "offset"
// @Succes        200 {object} []string
// @Failure       404 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/rbac/roles [GET]
func (h *handlerV1) ListRoles(ctx *gin.Context) {
	resp := h.enforcer.GetAllRoles()
	ctx.JSON(http.StatusOK, resp)
}

// @Security	ApiKeyAuth
// @Summary		Create new user role
// @Description	Create new user role
// @Tags        casbin
// @Accept     	json
// @Produce     json
// @Param       body body models.CreateUserRoleRequest true "body"
// @Success     200 {object} models.CreateUserRoleRequest
// @Failure     404 {object} models.Error
// @Failure     500 {object} models.Error
// @Router /v1/rbac/add-user-role [POST]
func (h *handlerV1) CreateRole(ctx *gin.Context) {
	var reqBody models.CreateUserRoleRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&reqBody); err != nil {
		h.log.Error("rbacHandler/CreateUserRole", zap.Error(err))
		return
	}

	if _, err := h.enforcer.AddPolicy(reqBody.Role, reqBody.Path, reqBody.Method); err != nil {
		h.log.Error("Error on grant access", zap.Error(err))
		return
	}
	h.enforcer.SavePolicy()
	ctx.JSON(http.StatusOK, reqBody)
}
