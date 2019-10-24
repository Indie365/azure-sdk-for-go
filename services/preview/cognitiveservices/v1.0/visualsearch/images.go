package visualsearch

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"io"
	"net/http"
)

// ImagesClient is the visual Search API lets you discover insights about an image such as visually similar images,
// shopping sources, and related searches. The API can also perform text recognition, identify entities (people,
// places, things), return other topical content for the user to explore, and more. For more information, see [Visual
// Search
// Overview](https://docs.microsoft.com/azure/cognitive-services/bing-visual-search/overview).<br/><br/><strong>NOTE:</strong>
// To comply with the new EU Copyright Directive in France, the Bing Visual Search API must omit some content from
// certain EU News sources for French users. The removed content may include thumbnail images and videos, video
// previews, and snippets which accompany search results from these sources. As a consequence, the Bing APIs may serve
// fewer results with thumbnail images and videos, video previews, and snippets to French users.
type ImagesClient struct {
	BaseClient
}

// NewImagesClient creates an instance of the ImagesClient client.
func NewImagesClient() ImagesClient {
	return ImagesClient{New()}
}

// VisualSearch sends the visual search request.
// Parameters:
// acceptLanguage - a comma-delimited list of one or more languages to use for user interface strings. The list
// is in decreasing order of preference. For additional information, including expected format, see
// [RFC2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html). This header and the
// [setLang](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-visual-search-api-v7-reference#setlang)
// query parameter are mutually exclusive; do not specify both. If you set this header, you must also specify
// the [cc](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-visual-search-api-v7-reference#cc)
// query parameter. To determine the market to return results for, Bing uses the first supported language it
// finds from the list and combines it with the cc parameter value. If the list does not include a supported
// language, Bing finds the closest language and market that supports the request or it uses an aggregated or
// default market for the results. To determine the market that Bing used, see the BingAPIs-Market header. Use
// this header and the cc query parameter only if you specify multiple languages. Otherwise, use the
// [mkt](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-visual-search-api-v7-reference#mkt)
// and
// [setLang](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-visual-search-api-v7-reference#setlang)
// query parameters. A user interface string is a string that's used as a label in a user interface. There are
// few user interface strings in the JSON response objects. Any links to Bing.com properties in the response
// objects apply the specified language.
// contentType - must be set to multipart/form-data and include a boundary parameter (for example,
// multipart/form-data; boundary=<boundary string>). For more details, see [Content form types](
// https://docs.microsoft.com/en-us/azure/cognitive-services/bing-visual-search/overview#content-form-types).
// userAgent - the user agent originating the request. Bing uses the user agent to provide mobile users with an
// optimized experience. Although optional, you are encouraged to always specify this header. The user-agent
// should be the same string that any commonly used browser sends. For information about user agents, see [RFC
// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html). The following are examples of user-agent
// strings. Windows Phone: Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0;
// ARM; Touch; NOKIA; Lumia 822). Android: Mozilla / 5.0 (Linux; U; Android 2.3.5; en - us; SCH - I500 Build /
// GINGERBREAD) AppleWebKit / 533.1 (KHTML; like Gecko) Version / 4.0 Mobile Safari / 533.1. iPhone: Mozilla /
// 5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X) AppleWebKit / 536.26 (KHTML; like Gecko) Mobile / 10B142
// iPhone4; 1 BingWeb / 3.03.1428.20120423. PC: Mozilla / 5.0 (Windows NT 6.3; WOW64; Trident / 7.0; Touch;
// rv:11.0) like Gecko. iPad: Mozilla / 5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit / 537.51.1 (KHTML,
// like Gecko) Version / 7.0 Mobile / 11A465 Safari / 9537.53.
// clientID - bing uses this header to provide users with consistent behavior across Bing API calls. Bing often
// flights new features and improvements, and it uses the client ID as a key for assigning traffic on different
// flights. If you do not use the same client ID for a user across multiple requests, then Bing may assign the
// user to multiple conflicting flights. Being assigned to multiple conflicting flights can lead to an
// inconsistent user experience. For example, if the second request has a different flight assignment than the
// first, the experience may be unexpected. Also, Bing can use the client ID to tailor web results to that
// client ID’s search history, providing a richer experience for the user. Bing also uses this header to help
// improve result rankings by analyzing the activity generated by a client ID. The relevance improvements help
// with better quality of results delivered by Bing APIs and in turn enables higher click-through rates for the
// API consumer. IMPORTANT: Although optional, you should consider this header required. Persisting the client
// ID across multiple requests for the same end user and device combination enables 1) the API consumer to
// receive a consistent user experience, and 2) higher click-through rates via better quality of results from
// the Bing APIs. Each user that uses your application on the device must have a unique, Bing generated client
// ID. If you do not include this header in the request, Bing generates an ID and returns it in the
// X-MSEdge-ClientID response header. The only time that you should NOT include this header in a request is the
// first time the user uses your app on that device. Use the client ID for each Bing API request that your app
// makes for this user on the device. Persist the client ID. To persist the ID in a browser app, use a
// persistent HTTP cookie to ensure the ID is used across all sessions. Do not use a session cookie. For other
// apps such as mobile apps, use the device's persistent storage to persist the ID. The next time the user uses
// your app on that device, get the client ID that you persisted. Bing responses may or may not include this
// header. If the response includes this header, capture the client ID and use it for all subsequent Bing
// requests for the user on that device. ATTENTION: You must ensure that this Client ID is not linkable to any
// authenticatable user account information. If you include the X-MSEdge-ClientID, you must not include cookies
// in the request.
// clientIP - the IPv4 or IPv6 address of the client device. The IP address is used to discover the user's
// location. Bing uses the location information to determine safe search behavior. Although optional, you are
// encouraged to always specify this header and the X-Search-Location header. Do not obfuscate the address (for
// example, by changing the last octet to 0). Obfuscating the address results in the location not being
// anywhere near the device's actual location, which may result in Bing serving erroneous results.
// location - a semicolon-delimited list of key/value pairs that describe the client's geographical location.
// Bing uses the location information to determine safe search behavior and to return relevant local content.
// Specify the key/value pair as <key>:<value>. The following are the keys that you use to specify the user's
// location. lat (required): The latitude of the client's location, in degrees. The latitude must be greater
// than or equal to -90.0 and less than or equal to +90.0. Negative values indicate southern latitudes and
// positive values indicate northern latitudes. long (required): The longitude of the client's location, in
// degrees. The longitude must be greater than or equal to -180.0 and less than or equal to +180.0. Negative
// values indicate western longitudes and positive values indicate eastern longitudes. re (required): The
// radius, in meters, which specifies the horizontal accuracy of the coordinates. Pass the value returned by
// the device's location service. Typical values might be 22m for GPS/Wi-Fi, 380m for cell tower triangulation,
// and 18,000m for reverse IP lookup. ts (optional): The UTC UNIX timestamp of when the client was at the
// location. (The UNIX timestamp is the number of seconds since January 1, 1970.) head (optional): The client's
// relative heading or direction of travel. Specify the direction of travel as degrees from 0 through 360,
// counting clockwise relative to true north. Specify this key only if the sp key is nonzero. sp (optional):
// The horizontal velocity (speed), in meters per second, that the client device is traveling. alt (optional):
// The altitude of the client device, in meters. are (optional): The radius, in meters, that specifies the
// vertical accuracy of the coordinates. Specify this key only if you specify the alt key. Although many of the
// keys are optional, the more information that you provide, the more accurate the location results are.
// Although optional, you are encouraged to always specify the user's geographical location. Providing the
// location is especially important if the client's IP address does not accurately reflect the user's physical
// location (for example, if the client uses VPN). For optimal results, you should include this header and the
// X-MSEdge-ClientIP header, but at a minimum, you should include this header.
// market - the market where the results come from. Typically, mkt is the country where the user is making the
// request from. However, it could be a different country if the user is not located in a country where Bing
// delivers results. The market must be in the form <language code>-<country code>. For example, en-US. The
// string is case insensitive. For a list of possible market values, see [Market
// Codes](https://docs.microsoft.com/en-us/azure/cognitive-services/bing-visual-search/supported-countries-markets).
// NOTE: If known, you are encouraged to always specify the market. Specifying the market helps Bing route the
// request and return an appropriate and optimal response. If you specify a market that is not listed in
// [Market
// Codes](https://docs.microsoft.com/en-us/azure/cognitive-services/bing-visual-search/supported-countries-markets),
// Bing uses a best fit market code based on an internal mapping that is subject to change.
// safeSearch - filter the image results in actions with type 'VisualSearch' for adult content. The following
// are the possible filter values. Off: May return images with adult content. Moderate: Do not return images
// with adult content. Strict: Do not return images with adult content. The default is Moderate. If the request
// comes from a market that Bing's adult policy requires that safeSearch is set to Strict, Bing ignores the
// safeSearch value and uses Strict. If you use the site: filter in the knowledge request, there is the chance
// that the response may contain adult content regardless of what the safeSearch query parameter is set to. Use
// site: only if you are aware of the content on the site and your scenario supports the possibility of adult
// content.
// setLang - the language to use for user interface strings. Specify the language using the ISO 639-1 2-letter
// language code. For example, the language code for English is EN. The default is EN (English). Although
// optional, you should always specify the language. Typically, you set setLang to the same language specified
// by mkt unless the user wants the user interface strings displayed in a different language. A user interface
// string is a string that's used as a label in a user interface. There are few user interface strings in the
// JSON response objects. Also, any links to Bing.com properties in the response objects apply the specified
// language.
// knowledgeRequest - the form data is a JSON object that identifies the image using an insights token or URL
// to the image. The object may also include an optional crop area that identifies an area of interest in the
// image. The insights token and URL are mutually exclusive – do not specify both. You may specify
// knowledgeRequest form data and image form data in the same request only if knowledgeRequest form data
// specifies the cropArea field only (it must not include an insights token or URL).
// imageParameter - the form data is an image binary. The Content-Disposition header's name parameter must be
// set to "image". You must specify an image binary if you do not use knowledgeRequest form data to specify the
// image; you may not use both forms to specify an image. You may specify knowledgeRequest form data and image
// form data in the same request only if knowledgeRequest form data specifies the cropArea field only  (it must
// not include an insights token or URL).
func (client ImagesClient) VisualSearch(ctx context.Context, acceptLanguage string, contentType string, userAgent string, clientID string, clientIP string, location string, market string, safeSearch SafeSearch, setLang string, knowledgeRequest string, imageParameter io.ReadCloser) (result ImageKnowledge, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImagesClient.VisualSearch")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.VisualSearchPreparer(ctx, acceptLanguage, contentType, userAgent, clientID, clientIP, location, market, safeSearch, setLang, knowledgeRequest, imageParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualsearch.ImagesClient", "VisualSearch", nil, "Failure preparing request")
		return
	}

	resp, err := client.VisualSearchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "visualsearch.ImagesClient", "VisualSearch", resp, "Failure sending request")
		return
	}

	result, err = client.VisualSearchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualsearch.ImagesClient", "VisualSearch", resp, "Failure responding to request")
	}

	return
}

// VisualSearchPreparer prepares the VisualSearch request.
func (client ImagesClient) VisualSearchPreparer(ctx context.Context, acceptLanguage string, contentType string, userAgent string, clientID string, clientIP string, location string, market string, safeSearch SafeSearch, setLang string, knowledgeRequest string, imageParameter io.ReadCloser) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	queryParameters := map[string]interface{}{}
	if len(market) > 0 {
		queryParameters["mkt"] = autorest.Encode("query", market)
	}
	if len(string(safeSearch)) > 0 {
		queryParameters["safeSearch"] = autorest.Encode("query", safeSearch)
	}
	if len(setLang) > 0 {
		queryParameters["setLang"] = autorest.Encode("query", setLang)
	}

	formDataParameters := map[string]interface{}{}
	if len(knowledgeRequest) > 0 {
		formDataParameters["knowledgeRequest"] = knowledgeRequest
	}
	if imageParameter != nil {
		formDataParameters["image"] = imageParameter
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/bing/v7.0", urlParameters),
		autorest.WithPath("/images/visualsearch"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithMultiPartFormData(formDataParameters),
		autorest.WithHeader("X-BingApis-SDK", "true"))
	if len(acceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(acceptLanguage)))
	}
	if len(contentType) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Content-Type", autorest.String(contentType)))
	}
	if len(userAgent) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("User-Agent", autorest.String(userAgent)))
	}
	if len(clientID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-MSEdge-ClientID", autorest.String(clientID)))
	}
	if len(clientIP) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-MSEdge-ClientIP", autorest.String(clientIP)))
	}
	if len(location) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-Search-Location", autorest.String(location)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// VisualSearchSender sends the VisualSearch request. The method will close the
// http.Response Body if it receives an error.
func (client ImagesClient) VisualSearchSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// VisualSearchResponder handles the response to the VisualSearch request. The method always
// closes the http.Response Body.
func (client ImagesClient) VisualSearchResponder(resp *http.Response) (result ImageKnowledge, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
