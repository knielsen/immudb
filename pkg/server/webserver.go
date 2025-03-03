/*
Copyright 2022 Codenotary Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/codenotary/immudb/pkg/logger"
	"github.com/codenotary/immudb/webconsole"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func StartWebServer(addr string, tlsConfig *tls.Config, s schema.ImmuServiceServer, l logger.Logger) (*http.Server, error) {
	proxyMux := runtime.NewServeMux()
	err := schema.RegisterImmuServiceHandlerServer(context.Background(), proxyMux, s)
	if err != nil {
		return nil, err
	}

	webMux := http.NewServeMux()
	webMux.Handle("/api/", http.StripPrefix("/api", proxyMux))

	err = webconsole.SetupWebconsole(webMux, l, addr)
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{Addr: addr, Handler: webMux}
	httpServer.TLSConfig = tlsConfig

	go func() {
		var err error
		if tlsConfig != nil && len(tlsConfig.Certificates) > 0 {
			l.Infof("Web API server enabled on %s/api (https)", addr)
			err = httpServer.ListenAndServeTLS("", "")
		} else {
			l.Infof("Web API server enabled on %s/api (http)", addr)
			err = httpServer.ListenAndServe()
		}

		if err == http.ErrServerClosed {
			l.Debugf("Web API/console server closed")
		} else {
			l.Errorf("Web API/console error: %s", err)
		}
	}()

	return httpServer, nil
}
