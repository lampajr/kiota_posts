package posts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620 "kiota_posts/client/models"
)

// PostsRequestBuilder builds and executes requests for operations under \posts
type PostsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PostsRequestBuilderGetQueryParameters get posts
type PostsRequestBuilderGetQueryParameters struct {
    // Filter results by title
    Title *string `uriparametername:"title"`
    // Filter results by user ID
    UserId *int32 `uriparametername:"userId"`
}
// PostsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PostsRequestBuilderGetQueryParameters
}
// PostsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPostsRequestBuilderInternal instantiates a new PostsRequestBuilder and sets the default values.
func NewPostsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostsRequestBuilder) {
    m := &PostsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/posts{?title*,userId*}", pathParameters),
    }
    return m
}
// NewPostsRequestBuilder instantiates a new PostsRequestBuilder and sets the default values.
func NewPostsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPostsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get posts
// returns a []Postable when successful
func (m *PostsRequestBuilder) Get(ctx context.Context, requestConfiguration *PostsRequestBuilderGetRequestConfiguration)([]i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.CreatePostFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable)
        }
    }
    return val, nil
}
// Post create post
// returns a []int32 when successful
func (m *PostsRequestBuilder) Post(ctx context.Context, body []int32, requestConfiguration *PostsRequestBuilderPostRequestConfiguration)([]int32, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    val := make([]int32, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*int32))
        }
    }
    return val, nil
}
// ToGetRequestInformation get posts
// returns a *RequestInformation when successful
func (m *PostsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PostsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create post
// returns a *RequestInformation when successful
func (m *PostsRequestBuilder) ToPostRequestInformation(ctx context.Context, body []int32, requestConfiguration *PostsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/posts", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetContentFromScalarCollection(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PostsRequestBuilder when successful
func (m *PostsRequestBuilder) WithUrl(rawUrl string)(*PostsRequestBuilder) {
    return NewPostsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
