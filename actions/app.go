package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	"gobuff_realworld_example_app/models"

	"github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	i18n "github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/packr/v2"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_gobuff_realworld_app_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)

		//AuthMiddlewares
		app.Use(SetCurrentUserMiddleware)
		app.Use(AuthorizeMiddleware)

		app.Middleware.Skip(AuthorizeMiddleware, HomeHandler)

		//Routes for Auth
		auth := app.Group("/auth")
		auth.GET("/login", AuthLoginHandler)
		auth.POST("/", AuthCreateHandler)
		auth.GET("/logout", AuthLogoutHandler)
		auth.Middleware.Skip(AuthorizeMiddleware, AuthLoginHandler, AuthCreateHandler)

		//Routes for User registration
		users := app.Group("/users")
		users.GET("/register", UsersRegisterHandler)
		users.GET("/profile/{user_email}", UsersProfileHandler).Name("userProfilePath")
		users.POST("/register", UsersCreateHandler)
		users.Middleware.Remove(AuthorizeMiddleware)

		// Routes for Following
		app.POST("/follow", UsersFollow)

		// Routes for Articles
		articles := app.Group("/articles")
		articles.POST("/new", ArticlesCreateHandler)
		articles.GET("/new", ArticlesNewHandler)
		articles.POST("/{slug}/comment", ArticlesCommentHandler).Name("articleCommentPath")
		articles.GET("/{slug}/delete", ArticlesDeleteHandler).Name("deleteArticlePath")
		articles.GET("/{slug}/edit", ArticlesEditHandler).Name("editArticlePath")
		articles.PUT("/{slug}/edit", ArticlesUpdateHandler).Name("editArticlePath")
		articles.POST("/star", ArticlesStarHandler)
		articles.GET("/{slug}", ArticlesReadHandler).Name("articlePath")
		articles.Middleware.Skip(AuthorizeMiddleware, ArticlesReadHandler)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("app:locales", "../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
