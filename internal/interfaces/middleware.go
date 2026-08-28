package interfaces

import (
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/infrastructure/validation"
	"golang.org/x/time/rate"
)

var FRONTEND_URL = os.Getenv("FRONTEND_URL")

const maxRequestBodyBytes int64 = 10 << 20

type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var clientLimiters = struct {
	sync.Mutex
	clients map[string]*clientLimiter
}{clients: make(map[string]*clientLimiter)}

func RequestSizeLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, maxRequestBodyBytes)
		ctx.Next()
	}
}

func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIP := ctx.ClientIP()
		if clientIP == "" {
			clientIP, _, _ = net.SplitHostPort(ctx.Request.RemoteAddr)
		}

		clientLimiters.Lock()
		now := time.Now()
		for ip, client := range clientLimiters.clients {
			if now.Sub(client.lastSeen) > 10*time.Minute {
				delete(clientLimiters.clients, ip)
			}
		}
		client, ok := clientLimiters.clients[clientIP]
		if !ok {
			client = &clientLimiter{limiter: rate.NewLimiter(rate.Every(time.Second), 30)}
			clientLimiters.clients[clientIP] = client
		}
		client.lastSeen = now
		allowed := client.limiter.Allow()
		clientLimiters.Unlock()

		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		ctx.Next()
	}
}

func SecurityHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("X-Frame-Options", "DENY")
		ctx.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		ctx.Header("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
		ctx.Header("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'none'")
		ctx.Next()
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken, ok := bearerToken(ctx.GetHeader("Authorization"))
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		token, err := validation.VerifyIDToken(ctx, idToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		if token.UID == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
			return
		}

		ctx.Set("userId", token.UID)
		ctx.Next()
	}
}

func authMiddlewareNoAbort() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken, ok := bearerToken(ctx.GetHeader("Authorization"))
		if !ok {
			ctx.Set("userId", "")
			ctx.Next()
			return
		}

		token, err := validation.VerifyIDToken(ctx, idToken)
		if err != nil {
			ctx.Set("userId", "")
		} else {
			if token.UID == "" {
				ctx.Set("userId", "")
			} else {
				ctx.Set("userId", token.UID)
			}
		}

		ctx.Next()
	}
}

func bearerToken(header string) (string, bool) {
	if !strings.HasPrefix(header, "Bearer ") {
		return "", false
	}
	token := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
	return token, token != ""
}

func SetCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			FRONTEND_URL,
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
			"Origin",
			"Access-Control-Request-Method",
			"Access-Control-Request-Methods",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Max-Age",
			"Access-Control-Allow-Credentials",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))
}
