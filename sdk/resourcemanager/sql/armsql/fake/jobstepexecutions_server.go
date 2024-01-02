//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// JobStepExecutionsServer is a fake server for instances of the armsql.JobStepExecutionsClient type.
type JobStepExecutionsServer struct {
	// Get is the fake for method JobStepExecutionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, stepName string, options *armsql.JobStepExecutionsClientGetOptions) (resp azfake.Responder[armsql.JobStepExecutionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByJobExecutionPager is the fake for method JobStepExecutionsClient.NewListByJobExecutionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByJobExecutionPager func(resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *armsql.JobStepExecutionsClientListByJobExecutionOptions) (resp azfake.PagerResponder[armsql.JobStepExecutionsClientListByJobExecutionResponse])
}

// NewJobStepExecutionsServerTransport creates a new instance of JobStepExecutionsServerTransport with the provided implementation.
// The returned JobStepExecutionsServerTransport instance is connected to an instance of armsql.JobStepExecutionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobStepExecutionsServerTransport(srv *JobStepExecutionsServer) *JobStepExecutionsServerTransport {
	return &JobStepExecutionsServerTransport{
		srv:                        srv,
		newListByJobExecutionPager: newTracker[azfake.PagerResponder[armsql.JobStepExecutionsClientListByJobExecutionResponse]](),
	}
}

// JobStepExecutionsServerTransport connects instances of armsql.JobStepExecutionsClient to instances of JobStepExecutionsServer.
// Don't use this type directly, use NewJobStepExecutionsServerTransport instead.
type JobStepExecutionsServerTransport struct {
	srv                        *JobStepExecutionsServer
	newListByJobExecutionPager *tracker[azfake.PagerResponder[armsql.JobStepExecutionsClientListByJobExecutionResponse]]
}

// Do implements the policy.Transporter interface for JobStepExecutionsServerTransport.
func (j *JobStepExecutionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "JobStepExecutionsClient.Get":
		resp, err = j.dispatchGet(req)
	case "JobStepExecutionsClient.NewListByJobExecutionPager":
		resp, err = j.dispatchNewListByJobExecutionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *JobStepExecutionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions/(?P<jobExecutionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/steps/(?P<stepName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	jobExecutionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobExecutionId")])
	if err != nil {
		return nil, err
	}
	stepNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("stepName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, jobExecutionIDParam, stepNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobExecution, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobStepExecutionsServerTransport) dispatchNewListByJobExecutionPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListByJobExecutionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByJobExecutionPager not implemented")}
	}
	newListByJobExecutionPager := j.newListByJobExecutionPager.get(req)
	if newListByJobExecutionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions/(?P<jobExecutionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/steps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		jobExecutionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobExecutionId")])
		if err != nil {
			return nil, err
		}
		createTimeMinUnescaped, err := url.QueryUnescape(qp.Get("createTimeMin"))
		if err != nil {
			return nil, err
		}
		createTimeMinParam, err := parseOptional(createTimeMinUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		createTimeMaxUnescaped, err := url.QueryUnescape(qp.Get("createTimeMax"))
		if err != nil {
			return nil, err
		}
		createTimeMaxParam, err := parseOptional(createTimeMaxUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		endTimeMinUnescaped, err := url.QueryUnescape(qp.Get("endTimeMin"))
		if err != nil {
			return nil, err
		}
		endTimeMinParam, err := parseOptional(endTimeMinUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		endTimeMaxUnescaped, err := url.QueryUnescape(qp.Get("endTimeMax"))
		if err != nil {
			return nil, err
		}
		endTimeMaxParam, err := parseOptional(endTimeMaxUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		isActiveUnescaped, err := url.QueryUnescape(qp.Get("isActive"))
		if err != nil {
			return nil, err
		}
		isActiveParam, err := parseOptional(isActiveUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsql.JobStepExecutionsClientListByJobExecutionOptions
		if createTimeMinParam != nil || createTimeMaxParam != nil || endTimeMinParam != nil || endTimeMaxParam != nil || isActiveParam != nil || skipParam != nil || topParam != nil {
			options = &armsql.JobStepExecutionsClientListByJobExecutionOptions{
				CreateTimeMin: createTimeMinParam,
				CreateTimeMax: createTimeMaxParam,
				EndTimeMin:    endTimeMinParam,
				EndTimeMax:    endTimeMaxParam,
				IsActive:      isActiveParam,
				Skip:          skipParam,
				Top:           topParam,
			}
		}
		resp := j.srv.NewListByJobExecutionPager(resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, jobExecutionIDParam, options)
		newListByJobExecutionPager = &resp
		j.newListByJobExecutionPager.add(req, newListByJobExecutionPager)
		server.PagerResponderInjectNextLinks(newListByJobExecutionPager, req, func(page *armsql.JobStepExecutionsClientListByJobExecutionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByJobExecutionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListByJobExecutionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByJobExecutionPager) {
		j.newListByJobExecutionPager.remove(req)
	}
	return resp, nil
}
