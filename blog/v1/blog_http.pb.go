// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: blog/v1/blog.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBlogServiceArticleDetail = "/blog.v1.BlogService/ArticleDetail"
const OperationBlogServiceChangeUserState = "/blog.v1.BlogService/ChangeUserState"
const OperationBlogServiceDeleteArticle = "/blog.v1.BlogService/DeleteArticle"
const OperationBlogServiceGetSystemConfig = "/blog.v1.BlogService/GetSystemConfig"
const OperationBlogServiceHitsArticle = "/blog.v1.BlogService/HitsArticle"
const OperationBlogServiceLikeArticle = "/blog.v1.BlogService/LikeArticle"
const OperationBlogServiceListArticle = "/blog.v1.BlogService/ListArticle"
const OperationBlogServiceListUser = "/blog.v1.BlogService/ListUser"
const OperationBlogServiceSaveArticle = "/blog.v1.BlogService/SaveArticle"
const OperationBlogServiceSaveSystemConfig = "/blog.v1.BlogService/SaveSystemConfig"
const OperationBlogServiceSaveUser = "/blog.v1.BlogService/SaveUser"
const OperationBlogServiceUserByAccount = "/blog.v1.BlogService/UserByAccount"
const OperationBlogServiceUserDetail = "/blog.v1.BlogService/UserDetail"

type BlogServiceHTTPServer interface {
	ArticleDetail(context.Context, *ReqArticleDetail) (*RespArticleDetail, error)
	ChangeUserState(context.Context, *ReqChangeUserState) (*RespChangeUserState, error)
	DeleteArticle(context.Context, *ReqDeleteArticle) (*RespDeleteArticle, error)
	GetSystemConfig(context.Context, *ReqGetSystemConfig) (*RespGetSystemConfig, error)
	HitsArticle(context.Context, *ReqHitsArticle) (*RespHitsArticle, error)
	LikeArticle(context.Context, *ReqLikeArticle) (*RespLikeArticle, error)
	// ListArticle 文章相关api
	ListArticle(context.Context, *ReqListArticle) (*RespListArticle, error)
	ListUser(context.Context, *ReqListUser) (*RespListUser, error)
	SaveArticle(context.Context, *ReqSaveArticle) (*RespSaveArticle, error)
	// SaveSystemConfig gconfig相关api
	SaveSystemConfig(context.Context, *ReqSaveSystemConfig) (*RespSaveSystemConfig, error)
	SaveUser(context.Context, *ReqSaveUser) (*RespSaveUser, error)
	UserByAccount(context.Context, *ReqUserByAccount) (*RespUserDetail, error)
	// UserDetail user相关api
	UserDetail(context.Context, *ReqUserDetail) (*RespUserDetail, error)
}

func RegisterBlogServiceHTTPServer(s *http.Server, srv BlogServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/user", _BlogService_UserDetail0_HTTP_Handler(srv))
	r.POST("/user/acc", _BlogService_UserByAccount0_HTTP_Handler(srv))
	r.POST("/user/save", _BlogService_SaveUser0_HTTP_Handler(srv))
	r.POST("/user/changestate", _BlogService_ChangeUserState0_HTTP_Handler(srv))
	r.POST("/user/list", _BlogService_ListUser0_HTTP_Handler(srv))
	r.POST("/config/save", _BlogService_SaveSystemConfig0_HTTP_Handler(srv))
	r.POST("/config", _BlogService_GetSystemConfig0_HTTP_Handler(srv))
	r.GET("/article/list", _BlogService_ListArticle0_HTTP_Handler(srv))
	r.POST("/article/save", _BlogService_SaveArticle0_HTTP_Handler(srv))
	r.POST("/article/del", _BlogService_DeleteArticle0_HTTP_Handler(srv))
	r.GET("/article/{article_id}", _BlogService_ArticleDetail0_HTTP_Handler(srv))
	r.POST("/article/like", _BlogService_LikeArticle0_HTTP_Handler(srv))
	r.POST("/article/hits", _BlogService_HitsArticle0_HTTP_Handler(srv))
}

func _BlogService_UserDetail0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqUserDetail
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceUserDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserDetail(ctx, req.(*ReqUserDetail))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespUserDetail)
		return ctx.Result(200, reply)
	}
}

func _BlogService_UserByAccount0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqUserByAccount
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceUserByAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserByAccount(ctx, req.(*ReqUserByAccount))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespUserDetail)
		return ctx.Result(200, reply)
	}
}

func _BlogService_SaveUser0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqSaveUser
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceSaveUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveUser(ctx, req.(*ReqSaveUser))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespSaveUser)
		return ctx.Result(200, reply)
	}
}

func _BlogService_ChangeUserState0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqChangeUserState
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceChangeUserState)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangeUserState(ctx, req.(*ReqChangeUserState))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespChangeUserState)
		return ctx.Result(200, reply)
	}
}

func _BlogService_ListUser0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqListUser
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ReqListUser))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespListUser)
		return ctx.Result(200, reply)
	}
}

func _BlogService_SaveSystemConfig0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqSaveSystemConfig
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceSaveSystemConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveSystemConfig(ctx, req.(*ReqSaveSystemConfig))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespSaveSystemConfig)
		return ctx.Result(200, reply)
	}
}

func _BlogService_GetSystemConfig0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqGetSystemConfig
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceGetSystemConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSystemConfig(ctx, req.(*ReqGetSystemConfig))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespGetSystemConfig)
		return ctx.Result(200, reply)
	}
}

func _BlogService_ListArticle0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqListArticle
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceListArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticle(ctx, req.(*ReqListArticle))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespListArticle)
		return ctx.Result(200, reply)
	}
}

func _BlogService_SaveArticle0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqSaveArticle
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceSaveArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveArticle(ctx, req.(*ReqSaveArticle))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespSaveArticle)
		return ctx.Result(200, reply)
	}
}

func _BlogService_DeleteArticle0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqDeleteArticle
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceDeleteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*ReqDeleteArticle))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespDeleteArticle)
		return ctx.Result(200, reply)
	}
}

func _BlogService_ArticleDetail0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqArticleDetail
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceArticleDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ArticleDetail(ctx, req.(*ReqArticleDetail))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespArticleDetail)
		return ctx.Result(200, reply)
	}
}

func _BlogService_LikeArticle0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqLikeArticle
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceLikeArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LikeArticle(ctx, req.(*ReqLikeArticle))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespLikeArticle)
		return ctx.Result(200, reply)
	}
}

func _BlogService_HitsArticle0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqHitsArticle
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBlogServiceHitsArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HitsArticle(ctx, req.(*ReqHitsArticle))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespHitsArticle)
		return ctx.Result(200, reply)
	}
}

type BlogServiceHTTPClient interface {
	ArticleDetail(ctx context.Context, req *ReqArticleDetail, opts ...http.CallOption) (rsp *RespArticleDetail, err error)
	ChangeUserState(ctx context.Context, req *ReqChangeUserState, opts ...http.CallOption) (rsp *RespChangeUserState, err error)
	DeleteArticle(ctx context.Context, req *ReqDeleteArticle, opts ...http.CallOption) (rsp *RespDeleteArticle, err error)
	GetSystemConfig(ctx context.Context, req *ReqGetSystemConfig, opts ...http.CallOption) (rsp *RespGetSystemConfig, err error)
	HitsArticle(ctx context.Context, req *ReqHitsArticle, opts ...http.CallOption) (rsp *RespHitsArticle, err error)
	LikeArticle(ctx context.Context, req *ReqLikeArticle, opts ...http.CallOption) (rsp *RespLikeArticle, err error)
	ListArticle(ctx context.Context, req *ReqListArticle, opts ...http.CallOption) (rsp *RespListArticle, err error)
	ListUser(ctx context.Context, req *ReqListUser, opts ...http.CallOption) (rsp *RespListUser, err error)
	SaveArticle(ctx context.Context, req *ReqSaveArticle, opts ...http.CallOption) (rsp *RespSaveArticle, err error)
	SaveSystemConfig(ctx context.Context, req *ReqSaveSystemConfig, opts ...http.CallOption) (rsp *RespSaveSystemConfig, err error)
	SaveUser(ctx context.Context, req *ReqSaveUser, opts ...http.CallOption) (rsp *RespSaveUser, err error)
	UserByAccount(ctx context.Context, req *ReqUserByAccount, opts ...http.CallOption) (rsp *RespUserDetail, err error)
	UserDetail(ctx context.Context, req *ReqUserDetail, opts ...http.CallOption) (rsp *RespUserDetail, err error)
}

type BlogServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewBlogServiceHTTPClient(client *http.Client) BlogServiceHTTPClient {
	return &BlogServiceHTTPClientImpl{client}
}

func (c *BlogServiceHTTPClientImpl) ArticleDetail(ctx context.Context, in *ReqArticleDetail, opts ...http.CallOption) (*RespArticleDetail, error) {
	var out RespArticleDetail
	pattern := "/article/{article_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBlogServiceArticleDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) ChangeUserState(ctx context.Context, in *ReqChangeUserState, opts ...http.CallOption) (*RespChangeUserState, error) {
	var out RespChangeUserState
	pattern := "/user/changestate"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceChangeUserState))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) DeleteArticle(ctx context.Context, in *ReqDeleteArticle, opts ...http.CallOption) (*RespDeleteArticle, error) {
	var out RespDeleteArticle
	pattern := "/article/del"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceDeleteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) GetSystemConfig(ctx context.Context, in *ReqGetSystemConfig, opts ...http.CallOption) (*RespGetSystemConfig, error) {
	var out RespGetSystemConfig
	pattern := "/config"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceGetSystemConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) HitsArticle(ctx context.Context, in *ReqHitsArticle, opts ...http.CallOption) (*RespHitsArticle, error) {
	var out RespHitsArticle
	pattern := "/article/hits"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceHitsArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) LikeArticle(ctx context.Context, in *ReqLikeArticle, opts ...http.CallOption) (*RespLikeArticle, error) {
	var out RespLikeArticle
	pattern := "/article/like"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceLikeArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) ListArticle(ctx context.Context, in *ReqListArticle, opts ...http.CallOption) (*RespListArticle, error) {
	var out RespListArticle
	pattern := "/article/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBlogServiceListArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) ListUser(ctx context.Context, in *ReqListUser, opts ...http.CallOption) (*RespListUser, error) {
	var out RespListUser
	pattern := "/user/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) SaveArticle(ctx context.Context, in *ReqSaveArticle, opts ...http.CallOption) (*RespSaveArticle, error) {
	var out RespSaveArticle
	pattern := "/article/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceSaveArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) SaveSystemConfig(ctx context.Context, in *ReqSaveSystemConfig, opts ...http.CallOption) (*RespSaveSystemConfig, error) {
	var out RespSaveSystemConfig
	pattern := "/config/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceSaveSystemConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) SaveUser(ctx context.Context, in *ReqSaveUser, opts ...http.CallOption) (*RespSaveUser, error) {
	var out RespSaveUser
	pattern := "/user/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceSaveUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) UserByAccount(ctx context.Context, in *ReqUserByAccount, opts ...http.CallOption) (*RespUserDetail, error) {
	var out RespUserDetail
	pattern := "/user/acc"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceUserByAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BlogServiceHTTPClientImpl) UserDetail(ctx context.Context, in *ReqUserDetail, opts ...http.CallOption) (*RespUserDetail, error) {
	var out RespUserDetail
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceUserDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
