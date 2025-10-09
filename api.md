# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthUser">AuthUser</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthLoginResponse">AuthLoginResponse</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSignupResponse">AuthSignupResponse</a>

Methods:

- <code title="post /auth/login">client.Auth.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthLoginParams">AuthLoginParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthLoginResponse">AuthLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/signup">client.Auth.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthService.Signup">Signup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSignupParams">AuthSignupParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSignupResponse">AuthSignupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Setup

Response Types:

- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSetupNewResponse">AuthSetupNewResponse</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSetupGetResponse">AuthSetupGetResponse</a>

Methods:

- <code title="post /auth/setup">client.Auth.Setup.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSetupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSetupNewParams">AuthSetupNewParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSetupNewResponse">AuthSetupNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /auth/setup">client.Auth.Setup.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSetupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthSetupGetResponse">AuthSetupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Email

Response Types:

- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthEmailSendVerificationResponse">AuthEmailSendVerificationResponse</a>

Methods:

- <code title="post /auth/email/verify">client.Auth.Email.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthEmailService.SendVerification">SendVerification</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#AuthEmailSendVerificationResponse">AuthEmailSendVerificationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Posts

Response Types:

- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#Post">Post</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostNewResponse">PostNewResponse</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostListResponse">PostListResponse</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostGetBySlugResponse">PostGetBySlugResponse</a>

Methods:

- <code title="post /posts">client.Posts.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostNewParams">PostNewParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostNewResponse">PostNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /posts">client.Posts.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /posts/get">client.Posts.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostListParams">PostListParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostListResponse">PostListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /posts">client.Posts.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostDeleteParams">PostDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /posts/slug">client.Posts.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostService.GetBySlug">GetBySlug</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostGetBySlugParams">PostGetBySlugParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostGetBySlugResponse">PostGetBySlugResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Votes

Response Types:

- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#CurrentUserVote">CurrentUserVote</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#PostVote">PostVote</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#VoteAddResponse">VoteAddResponse</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#VoteRemoveResponse">VoteRemoveResponse</a>

Methods:

- <code title="post /votes">client.Votes.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#VoteService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#VoteAddParams">VoteAddParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#VoteAddResponse">VoteAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /votes">client.Votes.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#VoteService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#VoteRemoveParams">VoteRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#VoteRemoveResponse">VoteRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Boards

Response Types:

- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#Board">Board</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#BoardListResponse">BoardListResponse</a>

Methods:

- <code title="get /boards">client.Boards.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#BoardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#BoardListParams">BoardListParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#BoardListResponse">BoardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Roadmaps

Response Types:

- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#Roadmap">Roadmap</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapPrivate">RoadmapPrivate</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapListResponse">RoadmapListResponse</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapGetByURLResponse">RoadmapGetByURLResponse</a>
- <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapSearchByNameResponse">RoadmapSearchByNameResponse</a>

Methods:

- <code title="get /roadmaps">client.Roadmaps.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapListParams">RoadmapListParams</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapListResponse">RoadmapListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /roadmaps/{url}">client.Roadmaps.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapService.GetByURL">GetByURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, url <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapGetByURLResponse">RoadmapGetByURLResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /roadmaps/search/{name}">client.Roadmaps.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapService.SearchByName">SearchByName</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go">logchimp</a>.<a href="https://pkg.go.dev/github.com/logchimp/logchimp-go#RoadmapSearchByNameResponse">RoadmapSearchByNameResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
