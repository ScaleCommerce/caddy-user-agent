package user_agent_parse

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"

	ua "github.com/mileusna/useragent"
)

func init() {
	caddy.RegisterModule(UserAgentParse{})
	// Register the Caddyfile directive with order
	httpcaddyfile.RegisterDirectiveOrder("user_agent_parse", httpcaddyfile.Before, "basicauth")
	httpcaddyfile.RegisterHandlerDirective("user_agent_parse", parseCaddyfile)
	// Log that the module was registered
	caddy.Log().Named("user_agent_parse").Info("User Agent Parse module has been registered")
}

type UserAgentParse struct {
	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (UserAgentParse) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.user_agent_parse",
		New: func() caddy.Module { return new(UserAgentParse) },
	}
}

// parseCaddyfile parses the user_agent_parse directive.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m UserAgentParse

	// Parse any additional configuration if needed
	// For now, we don't expect any arguments
	for h.Next() {
		if h.NextArg() {
			return nil, h.ArgErr()
		}
	}

	return &m, nil
}

func (l *UserAgentParse) Provision(ctx caddy.Context) error {
	l.logger = ctx.Logger(l)

	// Log that the module was successfully provisioned
	l.logger.Info("User Agent Parse module has been successfully loaded and provisioned",
		zap.String("module_id", "http.handlers.user_agent_parse"))

	return nil
}

func (m *UserAgentParse) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	agent := r.Header.Get("User-Agent")
	ua := ua.Parse(agent)

	// Determine type: bot has priority over device type
	var userAgentType string
	if ua.Bot {
		userAgentType = "bot"
	} else if ua.Mobile {
		userAgentType = "mobile"
	} else if ua.Tablet {
		userAgentType = "tablet"
	} else {
		userAgentType = "desktop"
	}

	// Optional: Debug logging for each request (can be disabled for high traffic)
	m.logger.Debug("User Agent has been parsed",
		zap.String("user_agent", agent),
		zap.String("type", userAgentType))

	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)

	// Set single variable with type
	repl.Set("user_agent.type", userAgentType)

	return next.ServeHTTP(w, r)
}

// Interface guards.
var (
	_ caddy.Provisioner           = (*UserAgentParse)(nil)
	_ caddy.Module                = (*UserAgentParse)(nil)
	_ caddyhttp.MiddlewareHandler = (*UserAgentParse)(nil)
)
