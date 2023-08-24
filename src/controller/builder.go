// Copyright 2022 The ILLA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"net/http"

	"github.com/illacloud/builder-backend/pkg/builder"
	"github.com/illacloud/builder-backend/src/utils/accesscontrol"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BuilderRestHandler interface {
	GetTeamBuilderDesc(c *gin.Context)
}

type BuilderRestHandlerImpl struct {
	logger         *zap.SugaredLogger
	builderService builder.BuilderService
	AttributeGroup *accesscontrol.AttributeGroup
}

func NewBuilderRestHandlerImpl(logger *zap.SugaredLogger, builderService builder.BuilderService, attrg *accesscontrol.AttributeGroup) *BuilderRestHandlerImpl {
	return &BuilderRestHandlerImpl{
		logger:         logger,
		builderService: builderService,
		AttributeGroup: attrg,
	}
}

func (impl BuilderRestHandlerImpl) GetTeamBuilderDesc(c *gin.Context) {
	// fetch needed param
	teamID, errInGetTeamID := controller.GetMagicIntParamFromRequest(c, PARAM_TEAM_ID)
	userAuthToken, errInGetAuthToken := controller.GetUserAuthTokenFromHeader(c)
	if errInGetTeamID != nil || errInGetAuthToken != nil {
		return
	}

	// validate
	controller.AttributeGroup.Init()
	controller.AttributeGroup.SetTeamID(teamID)
	controller.AttributeGroup.SetUserAuthToken(userAuthToken)
	controller.AttributeGroup.SetUnitType(accesscontrol.UNIT_TYPE_BUILDER_DASHBOARD)
	controller.AttributeGroup.SetUnitID(accesscontrol.DEFAULT_UNIT_ID)
	canAccess, errInCheckAttr := controller.AttributeGroup.CanAccess(accesscontrol.ACTION_ACCESS_VIEW)
	if errInCheckAttr != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "error in check attribute: "+errInCheckAttr.Error())
		return
	}
	if !canAccess {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "you can not access this attribute due to access control policy.")
		return
	}

	// fetch data
	ret, err := controller.builderService.GetTeamBuilderDesc(teamID)
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_BUILDER_DESCRIPTION, "get builder description error: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, ret)
}
