/*
Copyright 2019 The HAProxy Ingress Controller Authors.

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

package types

import (
	"fmt"
	"sort"
)

// NewEndpoint ...
func (b *Backend) NewEndpoint(ip string, port int, target string) *Endpoint {
	endpoint := &Endpoint{
		IP:     ip,
		Port:   port,
		Target: target,
		Weight: 1,
	}
	b.Endpoints = append(b.Endpoints, endpoint)
	sort.Slice(b.Endpoints, func(i, j int) bool {
		return b.Endpoints[i].IP < b.Endpoints[j].IP
	})
	return endpoint
}

// HreqValidateUserlist ...
func (b *Backend) HreqValidateUserlist(userlist *Userlist) {
	// TODO implement
	b.HTTPRequests = append(b.HTTPRequests, &HTTPRequest{})
}

func (h *HTTPRequest) String() string {
	return fmt.Sprintf("%+v", *h)
}
