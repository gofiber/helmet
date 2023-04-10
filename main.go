// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io/
// 📝 Github Repository: https://github.com/gofiber/fiber

package helmet

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Config ...
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool
	// XSSProtection
	// Optional. Default value "0".
	XSSProtection string
	// ContentTypeNosniff
	// Optional. Default value "nosniff".
	ContentTypeNosniff string
	// XFrameOptions
	// Optional. Default value "SAMEORIGIN".
	// Possible values: "SAMEORIGIN", "DENY", "ALLOW-FROM uri"
	XFrameOptions string
	// HSTSMaxAge
	// Optional. Default value 0.
	HSTSMaxAge int
	// HSTSExcludeSubdomains
	// Optional. Default value false.
	HSTSExcludeSubdomains bool
	// ContentSecurityPolicy
	// Optional. Default value "".
	ContentSecurityPolicy string
	// CSPReportOnly
	// Optional. Default value false.
	CSPReportOnly bool
	// HSTSPreloadEnabled
	// Optional. Default value false.
	HSTSPreloadEnabled bool
	// ReferrerPolicy
	// Optional. Default value "no-referrer".
	ReferrerPolicy string
	// Permissions-Policy
	// Optional. Default value "".
	PermissionPolicy string
	// Cross-Origin-Embedder-Policy
	// Optional. Default value "require-corp".
	CrossOriginEmbedderPolicy string
	// Cross-Origin-Opener-Policy
	// Optional. Default value "same-origin".
	CrossOriginOpenerPolicy string
	// Cross-Origin-Resource-Policy
	// Optional. Default value "same-origin".
	CrossOriginResourcePolicy string
	// Origin-Agent-Cluster
	// Optional. Default value "?1".
	OriginAgentCluster string
	// X-DNS-Prefetch-Control
	// Optional. Default value "off".
	XDNSPrefetchControl string
	// X-Download-Options
	// Optional. Default value "noopen".
	XDownloadOptions string
	// X-Permitted-Cross-Domain-Policies
	// Optional. Default value "none".
	XPermittedCrossDomain string
}

// New ...
func New(config ...Config) fiber.Handler {
	// Init config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	// Set config default values
	if cfg.XSSProtection == "" {
		cfg.XSSProtection = "0"
	}
	if cfg.ContentTypeNosniff == "" {
		cfg.ContentTypeNosniff = "nosniff"
	}
	if cfg.XFrameOptions == "" {
		cfg.XFrameOptions = "SAMEORIGIN"
	}
	if cfg.ReferrerPolicy == "" {
		cfg.ReferrerPolicy = "no-referrer"
	}
	if cfg.CrossOriginEmbedderPolicy == "" {
		cfg.CrossOriginEmbedderPolicy = "require-corp"
	}
	if cfg.CrossOriginOpenerPolicy == "" {
		cfg.CrossOriginOpenerPolicy = "same-origin"
	}
	if cfg.CrossOriginResourcePolicy == "" {
		cfg.CrossOriginResourcePolicy = "same-origin"
	}
	if cfg.OriginAgentCluster == "" {
		cfg.OriginAgentCluster = "?1"
	}
	if cfg.XDNSPrefetchControl == "" {
		cfg.XDNSPrefetchControl = "off"
	}
	if cfg.XDownloadOptions == "" {
		cfg.XDownloadOptions = "noopen"
	}
	if cfg.XPermittedCrossDomain == "" {
		cfg.XPermittedCrossDomain = "none"
	}

	// Return middleware handler
	return func(c *fiber.Ctx) error {
		// Filter request to skip middleware
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		// Set headers
		if cfg.XSSProtection != "" {
			c.Set(fiber.HeaderXXSSProtection, cfg.XSSProtection)
		}
		if cfg.ContentTypeNosniff != "" {
			c.Set(fiber.HeaderXContentTypeOptions, cfg.ContentTypeNosniff)
		}
		if cfg.XFrameOptions != "" {
			c.Set(fiber.HeaderXFrameOptions, cfg.XFrameOptions)
		}
		if cfg.CrossOriginEmbedderPolicy != "" {
			c.Set("Cross-Origin-Embedder-Policy", cfg.CrossOriginEmbedderPolicy)
		}
		if cfg.CrossOriginOpenerPolicy != "" {
			c.Set("Cross-Origin-Opener-Policy", cfg.CrossOriginOpenerPolicy)
		}
		if cfg.CrossOriginResourcePolicy != "" {
			c.Set("Cross-Origin-Resource-Policy", cfg.CrossOriginResourcePolicy)
		}
		if cfg.OriginAgentCluster != "" {
			c.Set("Origin-Agent-Cluster", cfg.OriginAgentCluster)
		}
		if cfg.ReferrerPolicy != "" {
			c.Set("Referrer-Policy", cfg.ReferrerPolicy)
		}
		if cfg.XDNSPrefetchControl != "" {
			c.Set("X-DNS-Prefetch-Control", cfg.XDNSPrefetchControl)
		}
		if cfg.XDownloadOptions != "" {
			c.Set("X-Download-Options", cfg.XDownloadOptions)
		}
		if cfg.XPermittedCrossDomain != "" {
			c.Set("X-Permitted-Cross-Domain-Policies", cfg.XPermittedCrossDomain)
		}

		// Handle HSTS headers
		if c.Protocol() == "https" && cfg.HSTSMaxAge != 0 {
			subdomains := ""
			if !cfg.HSTSExcludeSubdomains {
				subdomains = "; includeSubDomains"
			}
			if cfg.HSTSPreloadEnabled {
				subdomains = fmt.Sprintf("%s; preload", subdomains)
			}
			c.Set(fiber.HeaderStrictTransportSecurity, fmt.Sprintf("max-age=%d%s", cfg.HSTSMaxAge, subdomains))
		}

		// Handle Content-Security-Policy headers
		if cfg.ContentSecurityPolicy != "" {
			if cfg.CSPReportOnly {
				c.Set(fiber.HeaderContentSecurityPolicyReportOnly, cfg.ContentSecurityPolicy)
			} else {
				c.Set(fiber.HeaderContentSecurityPolicy, cfg.ContentSecurityPolicy)
			}
		}

		// Handle Permissions-Policy headers
		if cfg.PermissionPolicy != "" {
			c.Set(fiber.HeaderPermissionsPolicy, cfg.PermissionPolicy)
		}

		return c.Next()
	}
}
