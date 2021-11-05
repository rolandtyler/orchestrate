package auth

import (
	"context"
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/sdk/client"
	clientutils "github.com/consensys/orchestrate/pkg/toolkit/app/http/client-utils"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/types/api"
	"github.com/consensys/orchestrate/tests/e2e"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
	"net/http"
)

func (s *authTestSuite) TestAPIKEY() {
	var schedule *api.ScheduleResponse
	var scheduleDefaultTenant *api.ScheduleResponse
	tenantID := "tenant-e2e"

	s.Run("should create resources successfully", func() {
		resp, err := s.env.Client.CreateSchedule(s.env.Ctx, &api.CreateScheduleRequest{})
		require.NoError(s.T(), err)
		scheduleDefaultTenant = resp

		assert.Equal(s.T(), multitenancy.DefaultTenant, resp.TenantID)
	})

	s.Run("should create resources successfully while impersonating", func() {
		headers := map[string]string{
			multitenancy.TenantIDHeader: tenantID,
		}
		resp, err := s.env.Client.CreateSchedule(context.WithValue(s.env.Ctx, clientutils.RequestHeaderKey, headers), &api.CreateScheduleRequest{})
		require.NoError(s.T(), err)
		schedule = resp

		assert.Equal(s.T(), resp.TenantID, tenantID)
	})

	s.Run("should succeed to retrieve the resource of an other tenant", func() {
		resp, err := s.env.Client.GetSchedule(s.env.Ctx, schedule.UUID)
		require.NoError(s.T(), err)

		assert.Equal(s.T(), resp.TenantID, schedule.TenantID)
	})

	s.Run("any tenant should retrieve the resources of the default tenant", func() {
		headers := map[string]string{
			multitenancy.TenantIDHeader: "different-tenant",
		}
		resp, err := s.env.Client.GetSchedule(context.WithValue(s.env.Ctx, clientutils.RequestHeaderKey, headers), scheduleDefaultTenant.UUID)
		require.NoError(s.T(), err)

		assert.Equal(s.T(), resp.TenantID, scheduleDefaultTenant.TenantID)
	})

	s.Run("should fail to retrieve the resource of an other tenant while impersonating with status Not Found", func() {
		headers := map[string]string{
			multitenancy.TenantIDHeader: "different-tenant",
		}
		_, err := s.env.Client.GetSchedule(context.WithValue(s.env.Ctx, clientutils.RequestHeaderKey, headers), schedule.UUID)
		require.True(s.T(), errors.IsNotFoundError(err))
	})

	s.Run("should fail to authenticate using invalid API Key with status Unauthorized", func() {
		orchestrateClient := client.NewHTTPClient(
			&http.Client{Transport: e2e.NewTestHttpTransport("", "invalidKey")},
			client.NewConfig(s.env.Config.ApiURL, nil),
		)

		_, err := orchestrateClient.GetSchedules(s.env.Ctx) // Call to random endpoint just to check we can't access it
		assert.True(s.T(), errors.IsUnauthorizedError(err))
	})

	// TODO: Delete accounts, currently not implemented
}
