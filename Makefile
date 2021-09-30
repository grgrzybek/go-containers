#
# Copyright 2021 Grzegorz Grzybek.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

FILES = \
	internal/app/gggc.go \
	internal/cmd/config/config.go \
	cmd/gggc/main.go

all: bin/hello bin/gggc

bin/hello: cmd/hello/main.go
	go build -o bin/hello github.com/grgrzybek/go-containers/cmd/hello

bin/gggc: $(FILES)
	go build -o bin/gggc github.com/grgrzybek/go-containers/cmd/gggc

clean:
	rm -rf bin

.PHONY: clean
