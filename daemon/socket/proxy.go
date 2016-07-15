package socket

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/facebookgo/httpdown"
	"github.com/go-zoo/bone"

	"github.com/arigatomachine/cli/daemon/config"
	"github.com/arigatomachine/cli/daemon/routes"
	"github.com/arigatomachine/cli/daemon/session"
)

type AuthProxy struct {
	u    *url.URL
	l    net.Listener
	s    httpdown.Server
	c    *config.Config
	sess session.Session
}

func NewAuthProxy(c *config.Config, sess session.Session) (*AuthProxy, error) {
	l, err := MakeSocket(c.ProxySocketPath)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(c.API)
	if err != nil {
		return nil, err
	}

	return &AuthProxy{u: u, l: l, c: c, sess: sess}, nil
}

func (p *AuthProxy) Listen() {
	mux := bone.New()
	// XXX: We must validate certs, and figure something out for local dev
	// see https://github.com/arigatomachine/cli/issues/432
	t := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	proxy := &httputil.ReverseProxy{
		Transport: t,
		Director: func(r *http.Request) {
			r.URL.Scheme = p.u.Scheme
			r.URL.Host = p.u.Host
			r.Host = p.u.Host
			r.URL.Path = r.URL.Path[6:]

			tok := p.sess.GetToken()
			if tok != "" {
				r.Header["Authorization"] = []string{"Bearer " + tok}
			}
		},
	}

	mux.Handle("/proxy/", proxy)
	mux.SubRoute("/v1", routes.NewRouteMux(p.c, p.sess, t))

	h := httpdown.HTTP{}
	p.s = h.Serve(&http.Server{Handler: loggingHandler(mux)}, p.l)
}

func (p *AuthProxy) Close() error {
	return p.s.Stop()
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		next.ServeHTTP(w, r)
		log.Printf("%s %s", r.Method, p)
	})
}
