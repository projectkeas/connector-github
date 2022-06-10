package eventBuilder

import (
	"fmt"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

type mapper func(map[string]interface{}, string) (cloudevents.Event, cee.ValidationError)

var (
	Version string            = "1.0"
	mappers map[string]mapper = map[string]mapper{
		// These rules aren't officially available so need to be put under a beta label
		//"branch_protection_rule": 		mapBranchProtectionRule,
		//"code_scanning_alert":            mapCodeScanningAlert,
		//"discussion":                     mapDiscussion,
		//"discussion_comment":             mapDiscussionComment,
		//"membership":                     mapMembership,
		//"pull_request_review_thread":     mapPullRequestReviewThread,
		//"repository_dispatch":            mapRepositoryDispatch,
		//"sponsorship":                    mapSponsorship,
		//"workflow_dispatch":              mapWorkflowDispatch,
		//"workflow_job":                   mapWorkflowJob,
		//"workflow_run":                   mapWorkflowRun,

		"check_run":                      mapCheckRun,
		"check_suite":                    mapCheckSuite,
		"commit_comment":                 mapCommitComment,
		"create":                         mapCreate,
		"delete":                         mapDelete,
		"deploy_key":                     mapDeployKey,
		"deployment":                     mapDeployment,
		"deployment_status":              mapDeploymentStatus,
		"fork":                           mapFork,
		"github_app_authorization":       mapGithubAppAuthorization,
		"gollum":                         mapGollum,
		"installation":                   mapInstallation,
		"installation_repositories":      mapInstallationRepositories,
		"issue_comment":                  mapIssueComment,
		"issues":                         mapIssues,
		"label":                          mapLabel,
		"marketplace_purchase":           mapMarketplacePurchase,
		"member":                         mapMember,
		"meta":                           mapMeta,
		"milestone":                      mapMilestone,
		"organization":                   mapOrganization,
		"org_block":                      mapOrgBlock,
		"package":                        mapPackage,
		"page_build":                     mapPageBuild,
		"project":                        mapProject,
		"project_card":                   mapProjectCard,
		"project_column":                 mapProjectColumn,
		"public":                         mapPublic,
		"pull_request":                   mapPullRequest,
		"pull_request_review":            mapPullRequestReview,
		"pull_request_review_comment":    mapPullRequestReviewComment,
		"push":                           mapPush,
		"release":                        mapRelease,
		"repository":                     mapRepository,
		"repository_import":              mapRepositoryImport,
		"repository_vulnerability_alert": mapVulnerabilityAlert,
		"security_advisory":              mapSecurityAdvisory,
		"star":                           mapStar,
		"status":                         mapStatus,
		"team":                           mapTeam,
		"team_add":                       mapTeamAdd,
		"watch":                          mapWatch,
	}
)

func NewCloudEventFromWebhook(payload map[string]interface{}, eventName string, id string) (cloudevents.Event, cee.ValidationError) {
	mapper, found := mappers[eventName]
	if found {
		return mapper(payload, id)
	}

	return cloudevents.Event{}, cee.ValidationError{
		"NotFound": fmt.Errorf("a mapper for the event '%s' was not found", eventName),
	}
}

// ============================
// 			Helpers
// ============================

func coalesce(strs ...string) string {
	for _, value := range strs {
		if value != "" {
			return value
		}
	}

	return ""
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t.UTC()
}

func lookupValueFromPayloadPath(payload map[string]interface{}, path ...string) string {
	for _, p := range path {
		item, found := payload[p]

		if !found {
			break
		}

		switch v := item.(type) {
		case string:
			return v
		case map[string]interface{}:
			payload = v
		default:
			return fmt.Sprint(v)
		}
	}

	return ""
}
