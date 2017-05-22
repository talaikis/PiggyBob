package auth

improt (
  "github.com/gin-gonic/gin"
  "../main"
)

func AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
    session, _ := main.store.Get(c.Request, "_pig_")
  	a := session.Values["UserID"].(string)

		if a == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.Abort()
		}
		c.Next()
	}
}
