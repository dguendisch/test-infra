# Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#############      builder       #############
FROM golang:1.11.1 AS builder

WORKDIR /go/src/github.com/gardener/test-infra
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install \
#   -ldflags "-w -X github.com/gardener/gardener/pkg/version.Version=$(cat VERSION)" \
  ./cmd/...

############# tm-controller #############
FROM alpine:3.8 AS tm-controller

RUN apk add --update bash curl

COPY --from=builder /go/bin/testmachinery-controller /testmachinery-controller
COPY ./.env /

WORKDIR /

ENTRYPOINT ["/testmachinery-controller"]

############# tm-run #############
FROM eu.gcr.io/gardener-project/cc/job-image:1.101.0 AS tm-run

COPY --from=builder /go/bin/testmachinery-run /testmachinery-run
COPY ./.env /

WORKDIR /

ENTRYPOINT ["/testmachinery-run"]