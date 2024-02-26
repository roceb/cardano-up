// Copyright 2024 Blink Labs Software
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

package pkgmgr

var RegistryPackages = []Package{
	{
		Name:        "cardano-node",
		Version:     "8.7.3",
		Description: "Cardano node software by Input Output Global",
		InstallSteps: []PackageInstallStep{
			{
				Docker: &PackageInstallStepDocker{
					ContainerName: "cardano-node",
					Image:         "ghcr.io/blinklabs-io/cardano-node:8.7.3",
					Env: map[string]string{
						"NETWORK":                  "{{ .Context.Network }}",
						"CARDANO_NODE_SOCKET_PATH": "/ipc/node.socket",
					},
					Ports: []string{
						"3001",
					},
				},
			},
			{
				File: &PackageInstallStepFile{
					Filename: "cardano-cli",
					// TODO: figure out how to get network magic for named network
					Content: `#!/bin/bash
docker exec -ti {{ .Package.Name }}-cardano-node cardano-cli $@
`,
				},
			},
		},
	},
}
