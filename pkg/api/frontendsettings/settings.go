// Package frontendsettings builds the FrontendSettingsDTO from static configuration.
// It is shared between the API handler (pkg/api) and the frontend service
// (pkg/services/frontend) to avoid duplicating the cfg → DTO mapping.
package frontendsettings

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/featuremgmt"
	"github.com/grafana/grafana/pkg/services/folder"
	"github.com/grafana/grafana/pkg/services/licensing"
	"github.com/grafana/grafana/pkg/setting"
)

// GetFrontendSettings returns the frontend settings that can be derived from static
// configuration without per-org or per-user context.
//
// Fields that require runtime services (datasources, panels, apps, OAuth providers,
// renderer availability, managed plugins, namespace) are left at their zero values;
// callers that have those services should augment the result afterwards.
func GetFrontendSettings(ctx context.Context, cfg *setting.Cfg, license licensing.Licensing, features featuremgmt.FeatureToggles) *dtos.FrontendSettingsDTO {
	version := setting.BuildVersion
	commit := setting.BuildCommit
	commitShort := getShortCommitHash(setting.BuildCommit, 10)
	buildstamp := setting.BuildStamp
	versionString := fmt.Sprintf(`%s v%s (%s)`, setting.ApplicationName, version, commitShort)

	trustedTypesDefaultPolicyEnabled := (cfg.CSPEnabled && strings.Contains(cfg.CSPTemplate, "require-trusted-types-for")) ||
		(cfg.CSPReportOnlyEnabled && strings.Contains(cfg.CSPReportOnlyTemplate, "require-trusted-types-for"))
	isCloudMigrationTarget := cfg.CloudMigration.Enabled && cfg.CloudMigration.IsTarget

	featureToggles := features.GetEnabled(ctx)
	// needed for backwards compatibility with external plugins
	featureToggles["topnav"] = true

	s := &dtos.FrontendSettingsDTO{
		MinRefreshInterval:   cfg.MinRefreshInterval,
		AppUrl:               cfg.AppURL,
		AppSubUrl:            cfg.AppSubURL,
		AuthProxyEnabled:     cfg.AuthProxy.Enabled,
		LdapEnabled:          cfg.LDAPAuthEnabled,
		JwtHeaderName:        cfg.JWTAuth.HeaderName,
		JwtUrlLogin:          cfg.JWTAuth.URLLogin,
		LiveEnabled:          cfg.LiveMaxConnections != 0,
		LiveMessageSizeLimit: cfg.LiveMessageSizeLimit,
		LiveNamespaced:       true,
		AutoAssignOrg:        cfg.AutoAssignOrg,
		VerifyEmailEnabled:   cfg.VerifyEmailEnabled,
		SigV4AuthEnabled:     cfg.SigV4AuthEnabled,
		AzureAuthEnabled:     cfg.AzureAuthEnabled,
		RbacEnabled:          true,
		ExploreEnabled:       cfg.ExploreEnabled,
		HelpEnabled:          cfg.HelpEnabled,
		ProfileEnabled:       cfg.ProfileEnabled,
		NewsFeedEnabled:      cfg.NewsFeedEnabled,
		QueryHistoryEnabled:  cfg.QueryHistoryEnabled,

		GoogleAnalyticsId:                   cfg.GoogleAnalyticsID,
		GoogleAnalytics4Id:                  cfg.GoogleAnalytics4ID,
		GoogleAnalytics4SendManualPageViews: cfg.GoogleAnalytics4SendManualPageViews,

		RudderstackWriteKey:        cfg.RudderstackWriteKey,
		RudderstackDataPlaneUrl:    cfg.RudderstackDataPlaneURL,
		RudderstackSdkUrl:          cfg.RudderstackSDKURL,
		RudderstackV3SdkUrl:        cfg.RudderstackV3SDKURL,
		RudderstackConfigUrl:       cfg.RudderstackConfigURL,
		RudderstackIntegrationsUrl: cfg.RudderstackIntegrationsURL,

		AnalyticsConsoleReporting:            cfg.FrontendAnalyticsConsoleReporting,
		DashboardPerformanceMetrics:          cfg.DashboardPerformanceMetrics,
		PanelSeriesLimit:                     cfg.PanelSeriesLimit,
		FeedbackLinksEnabled:                 cfg.FeedbackLinksEnabled,
		ApplicationInsightsConnectionString:  cfg.ApplicationInsightsConnectionString,
		ApplicationInsightsEndpointUrl:       cfg.ApplicationInsightsEndpointUrl,
		ApplicationInsightsAutoRouteTracking: cfg.ApplicationInsightsAutoRouteTracking,
		DisableLoginForm:                     cfg.DisableLoginForm,
		DisableUserSignUp:                    !cfg.AllowUserSignUp,
		LoginHint:                            cfg.LoginHint,
		PasswordHint:                         cfg.PasswordHint,
		ExternalUserMngInfo:                  cfg.ExternalUserMngInfo,
		ExternalUserMngLinkUrl:               cfg.ExternalUserMngLinkUrl,
		ExternalUserMngLinkName:              cfg.ExternalUserMngLinkName,
		ExternalUserMngAnalytics:             cfg.ExternalUserMngAnalytics,
		ExternalUserMngAnalyticsParams:       cfg.ExternalUserMngAnalyticsParams,
		ExternalUserUpgradeLinkUrl:           cfg.ExternalUserUpgradeLinkUrl,
		//nolint:staticcheck // ViewersCanEdit is deprecated but still used for backward compatibility
		ViewersCanEdit:                   cfg.ViewersCanEdit,
		DisableSanitizeHtml:              cfg.DisableSanitizeHtml,
		TrustedTypesDefaultPolicyEnabled: trustedTypesDefaultPolicyEnabled,
		CSPReportOnlyEnabled:             cfg.CSPReportOnlyEnabled,
		DateFormats:                      cfg.DateFormats,
		QuickRanges:                      cfg.QuickRanges,
		SecureSocksDSProxyEnabled:        cfg.SecureSocksDSProxy.Enabled && cfg.SecureSocksDSProxy.ShowUI,
		EnableFrontendSandboxForPlugins:  cfg.EnableFrontendSandboxForPlugins,
		PluginRestrictedAPIsAllowList:    cfg.PluginRestrictedAPIsAllowList,
		PluginRestrictedAPIsBlockList:    cfg.PluginRestrictedAPIsBlockList,
		PublicDashboardsEnabled:          cfg.PublicDashboardsEnabled,
		CloudMigrationEnabled:            cfg.CloudMigration.Enabled,
		CloudMigrationIsTarget:           isCloudMigrationTarget,
		CloudMigrationPollIntervalMs:     int(cfg.CloudMigration.FrontendPollInterval.Milliseconds()),
		SharedWithMeFolderUID:            folder.SharedWithMeFolderUID,
		RootFolderUID:                    accesscontrol.GeneralFolderUID,
		LocalFileSystemAvailable:         cfg.LocalFileSystemAvailable,
		ReportingStaticContext:           cfg.ReportingStaticContext,
		ExploreDefaultTimeOffset:         cfg.ExploreDefaultTimeOffset,
		ExploreHideLogsDownload:          cfg.ExploreHideLogsDownload,

		DefaultDatasourceManageAlertsUIToggle:          cfg.DefaultDatasourceManageAlertsUIToggle,
		DefaultAllowRecordingRulesTargetAlertsUIToggle: cfg.DefaultAllowRecordingRulesTargetAlertsUIToggle,

		BuildInfo: dtos.FrontendSettingsBuildInfoDTO{
			Version:       version,
			VersionString: versionString,
			Commit:        commit,
			CommitShort:   commitShort,
			Buildstamp:    buildstamp,
			Edition:       license.Edition(),
			Env:           cfg.Env,
		},

		LicenseInfo: dtos.FrontendSettingsLicenseInfoDTO{
			Expiry:          license.Expiry(),
			StateInfo:       license.StateInfo(),
			Edition:         license.Edition(),
			EnabledFeatures: license.EnabledFeatures(),
		},

		FeatureToggles:       featureToggles,
		AnonymousEnabled:     cfg.Anonymous.Enabled,
		AnonymousDeviceLimit: cfg.Anonymous.DeviceLimit,

		RendererDefaultImageWidth:           cfg.RendererDefaultImageWidth,
		RendererDefaultImageHeight:          cfg.RendererDefaultImageHeight,
		RendererDefaultImageScale:           cfg.RendererDefaultImageScale,
		Http2Enabled:                        cfg.Protocol == setting.HTTP2Scheme || cfg.Protocol == setting.SocketHTTP2Scheme,
		GrafanaJavascriptAgent:              cfg.GrafanaJavascriptAgent,
		PluginCatalogURL:                    cfg.PluginCatalogURL,
		PluginAdminEnabled:                  cfg.PluginAdminEnabled,
		PluginAdminExternalManageEnabled:    cfg.PluginAdminEnabled && cfg.PluginAdminExternalManageEnabled,
		PluginCatalogHiddenPlugins:          cfg.PluginCatalogHiddenPlugins,
		PluginCatalogPreinstalledPlugins:    append(cfg.PreinstallPluginsAsync, cfg.PreinstallPluginsSync...),
		PluginCatalogPreinstalledAutoUpdate: cfg.PreinstallAutoUpdate,
		ExpressionsEnabled:                  cfg.ExpressionsEnabled,
		AwsAllowedAuthProviders:             cfg.AWSAllowedAuthProviders,
		AwsAssumeRoleEnabled:                cfg.AWSAssumeRoleEnabled,
		AwsPerDatasourceHTTPProxyEnabled:    cfg.AWSPerDatasourceHTTPProxyEnabled,
		SupportBundlesEnabled:               isSupportBundlesEnabled(cfg),

		Azure: dtos.FrontendSettingsAzureDTO{
			Cloud:                                  cfg.Azure.Cloud,
			Clouds:                                 cfg.Azure.CustomClouds(),
			ManagedIdentityEnabled:                 cfg.Azure.ManagedIdentityEnabled,
			WorkloadIdentityEnabled:                cfg.Azure.WorkloadIdentityEnabled,
			UserIdentityEnabled:                    cfg.Azure.UserIdentityEnabled,
			UserIdentityFallbackCredentialsEnabled: cfg.Azure.UserIdentityFallbackCredentialsEnabled,
			AzureEntraPasswordCredentialsEnabled:   cfg.Azure.AzureEntraPasswordCredentialsEnabled,
		},

		Caching: dtos.FrontendSettingsCachingDTO{
			Enabled:           cfg.SectionWithEnvOverrides("caching").Key("enabled").MustBool(true),
			CleanCacheEnabled: cfg.SectionWithEnvOverrides("caching").Key("clean_cache_enabled").MustBool(true),
			DefaultTTLMs:      cfg.SectionWithEnvOverrides("caching").Key("ttl").MustDuration(time.Minute * 5).Milliseconds(),
		},
		RecordedQueries: dtos.FrontendSettingsRecordedQueriesDTO{
			Enabled: cfg.SectionWithEnvOverrides("recorded_queries").Key("enabled").MustBool(true),
		},
		Reporting: dtos.FrontendSettingsReportingDTO{
			Enabled: cfg.SectionWithEnvOverrides("reporting").Key("enabled").MustBool(true),
		},
		Analytics: dtos.FrontendSettingsAnalyticsDTO{
			Enabled: cfg.SectionWithEnvOverrides("analytics").Key("enabled").MustBool(true),
		},

		UnifiedAlerting: dtos.FrontendSettingsUnifiedAlertingDTO{
			MinInterval: cfg.UnifiedAlerting.MinInterval.String(),
		},

		TokenExpirationDayLimit: cfg.SATokenExpirationDayLimit,
		SnapshotEnabled:         cfg.SnapshotEnabled,

		SqlConnectionLimits: dtos.FrontendSettingsSqlConnectionLimitsDTO{
			MaxOpenConns:    cfg.SqlDatasourceMaxOpenConnsDefault,
			MaxIdleConns:    cfg.SqlDatasourceMaxIdleConnsDefault,
			ConnMaxLifetime: cfg.SqlDatasourceMaxConnLifetimeDefault,
		},
		OpenFeatureContext: cfg.OpenFeature.ContextAttrs,
	}

	if cfg.UnifiedAlerting.StateHistory.Enabled {
		s.UnifiedAlerting.StateHistory = &dtos.FrontendSettingsUnifiedAlertingStateHistoryDTO{
			Backend: cfg.UnifiedAlerting.StateHistory.Backend,
			Primary: cfg.UnifiedAlerting.StateHistory.MultiPrimary,
		}
		if cfg.UnifiedAlerting.StateHistory.PrometheusTargetDatasourceUID != "" {
			s.UnifiedAlerting.StateHistory.PrometheusTargetDatasourceUID = cfg.UnifiedAlerting.StateHistory.PrometheusTargetDatasourceUID
		}
		if cfg.UnifiedAlerting.StateHistory.PrometheusMetricName != "" {
			s.UnifiedAlerting.StateHistory.PrometheusMetricName = cfg.UnifiedAlerting.StateHistory.PrometheusMetricName
		}
		s.UnifiedAlerting.AlertStateHistoryBackend = cfg.UnifiedAlerting.StateHistory.Backend
		s.UnifiedAlerting.AlertStateHistoryPrimary = cfg.UnifiedAlerting.StateHistory.MultiPrimary
	}

	s.UnifiedAlerting.RecordingRulesEnabled = cfg.UnifiedAlerting.RecordingRules.Enabled
	s.UnifiedAlerting.DefaultRecordingRulesTargetDatasourceUID = cfg.UnifiedAlerting.RecordingRules.DefaultDatasourceUID

	if cfg.UnifiedAlerting.Enabled != nil {
		s.UnifiedAlertingEnabled = *cfg.UnifiedAlerting.Enabled
	}

	if cfg.GeomapDefaultBaseLayerConfig != nil {
		s.GeomapDefaultBaseLayerConfig = &cfg.GeomapDefaultBaseLayerConfig
	}
	if !cfg.GeomapEnableCustomBaseLayers {
		s.GeomapDisableCustomBaseLayer = true
	}

	// cfg-based auth fields; per-provider OAuth skip-sync fields are filled by the caller
	s.Auth = dtos.FrontendSettingsAuthDTO{
		AuthProxyEnableLoginToken:     cfg.AuthProxy.EnableLoginToken,
		SAMLSkipOrgRoleSync:           cfg.SAMLSkipOrgRoleSync,
		LDAPSkipOrgRoleSync:           cfg.LDAPSkipOrgRoleSync,
		JWTAuthSkipOrgRoleSync:        cfg.JWTAuth.SkipOrgRoleSync,
		DisableLogin:                  cfg.DisableLogin,
		BasicAuthStrongPasswordPolicy: cfg.BasicAuthStrongPasswordPolicy,
		DisableSignoutMenu:            cfg.DisableSignoutMenu,
	}

	return s
}

func isSupportBundlesEnabled(cfg *setting.Cfg) bool {
	return cfg.SectionWithEnvOverrides("support_bundles").Key("enabled").MustBool(true)
}

func getShortCommitHash(commitHash string, maxLength int) string {
	if len(commitHash) > maxLength {
		return commitHash[:maxLength]
	}
	return commitHash
}
