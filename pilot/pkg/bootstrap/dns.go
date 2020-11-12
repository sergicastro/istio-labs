// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bootstrap

import (
	"istio.io/istio/pkg/config/dns"
	"istio.io/pkg/log"
)

// initDnsResolver initializes DNS resolver.
func (s *Server) initDnsResolver() {
	log.Info("initializing DNS resolver")

	if s.environment.Resolver == nil {
		log.Info("DNS resolver configuration not provided, falling back to the default")
		s.environment.Resolver = dns.NewResolver(dns.ResolverOpts{})
	}
}