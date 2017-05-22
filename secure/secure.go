package secure

import (
    "os"
    "github.com/gin-gonic/gin"
    "github.com/unrolled/secure"
)

func TLSRedirect() gin.HandlerFunc {
        return func(c *gin.Context) {
            secureMiddleware := secure.New(secure.Options{
                AllowedHosts:          []string{os.Getenv("DOMAIN")},
                HostsProxyHeaders:     []string{"X-Forwarded-Host"},
                SSLRedirect: true,
                SSLHost:     os.Getenv("SECURE_HOST"),
                STSSeconds:            315360000,
            		STSIncludeSubdomains:  true,
            		STSPreload:            true,
            		FrameDeny:             true,
            		ContentTypeNosniff:    true,
            		BrowserXssFilter:      true,
            		ContentSecurityPolicy: "default-src 'self'",
            		PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains"`,
            })
            err := secureMiddleware.Process(c.Writer, c.Request)

            // If there was an error, do not continue.
            if err != nil {
                return
            }

            c.Next()
        }
    }
