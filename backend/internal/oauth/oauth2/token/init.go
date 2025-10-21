/*
 * Copyright (c) 2025, WSO2 LLC. (https://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package token

import (
	"net/http"

	"github.com/asgardeo/thunder/internal/application"
	"github.com/asgardeo/thunder/internal/oauth/oauth2/granthandlers"
	"github.com/asgardeo/thunder/internal/oauth/scope"
	"github.com/asgardeo/thunder/internal/system/middleware"
)

// Initialize initializes the token handler and registers its routes.
func Initialize(
	mux *http.ServeMux,
	appService application.ApplicationServiceInterface,
	grantHandlerProvider granthandlers.GrantHandlerProviderInterface,
	scopeValidator scope.ScopeValidatorInterface,
) TokenHandlerInterface {
	tokenHandler := newTokenHandler(appService, grantHandlerProvider, scopeValidator)
	registerRoutes(mux, tokenHandler)
	return tokenHandler
}

// registerRoutes registers the routes for the TokenService.
func registerRoutes(mux *http.ServeMux, tokenHandler TokenHandlerInterface) {
	opts := middleware.CORSOptions{
		AllowedMethods:   "POST",
		AllowedHeaders:   "Content-Type, Authorization",
		AllowCredentials: true,
	}
	mux.HandleFunc(middleware.WithCORS("POST /oauth2/token", tokenHandler.HandleTokenRequest, opts))
}
