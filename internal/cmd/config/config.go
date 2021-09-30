/*
 * Copyright 2021 Grzegorz Grzybek.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"
	ccConfig "github.com/containers/common/pkg/config"
	ciConfig "github.com/containers/image/v5/pkg/sysregistriesv2"
	"github.com/containers/image/v5/types"
	csConfig "github.com/containers/storage/types"
	gggc "github.com/grgrzybek/go-containers/internal/app"
	"github.com/spf13/cobra"
)

func init() {
	var cmd = &cobra.Command{
		Use:   "config",
		Short: "Using containers/* libraries for configuration analysis",
		Run:   run,
	}
	gggc.RootCmd.AddCommand(cmd)
}

func run(_ *cobra.Command, _ []string) {
	// here's a selected list of configuration references made by podman
	// cmd/podman/root.go initializes rootCmd with PersistentPreRunE=persistentPreRunE() function
	// - cmd/podman/registry/config.PodmanConfig() is called
	// - vendor/github.com/containers/common/pkg/config/config.DefaultConfig() is called
	// - $CONTAINERS_STORAGE_CONF variable is checked

	// - vendor/github.com/containers/storage/types/utils.DefaultConfigFile is called
	location, _ := csConfig.DefaultConfigFile(true)
	fmt.Printf("Default containers storage configuration (user): %s\n", location)
	location, _ = csConfig.DefaultConfigFile(false)
	fmt.Printf("Default containers storage configuration (root): %s\n", location)
	// (hardcoded /var/lib/containers/storage if [storage]/graphroot is empty)

	// [engine]/static_dir is set to [storage]/graphroot + "/libpod"
	// [engine]/volume_path is set to [storage]/graphroot + "/volumes"
	// [engine]/hooks_dir is set to "/usr/share/containers/oci/hooks.d" - hardcoded (?)

	// [engine]/SignaturePolicyPath (no toml key) is set to "/etc/containers/policy.json" (root), or
	//                                                      "~/.config/containers/policy.json" (user)
	// [network]/network_config_dir is set to "/etc/cni/net.d/" (root), or
	//                                        "~/.config/cni/net.d" (user)
	// [network]/default_subnet is set to "10.88.0.0/16"

	// - vendor/github.com/containers/common/pkg/config/config.systemConfigs() is called
	// - $CONTAINERS_CONF env variable
	// - /usr/share/containers/containers.conf
	fmt.Printf("Default containers configuration: %s\n", ccConfig.DefaultContainersConfig)
	// - /etc/containers/containers.conf
	fmt.Printf("Override containers configuration: %s\n", ccConfig.OverrideContainersConfig)
	// - /etc/containers/containers.conf.d/*.conf
	// - ~/.config/containers/containers.conf
	// - ~/.config/containers/containers.conf.d/*.conf
	fmt.Printf("Rootless containers configuration: ~/%s\n", ccConfig.UserOverrideContainersConfig)
	fmt.Printf("config.Path(): %s\n", ccConfig.Path())

	// for rootless installation, [network]/network_config_dir is changed to "" (by podman, not by containers/config)
	config, _ := ccConfig.DefaultConfig()
	fmt.Printf("Network config dir: %s\n", config.Network.NetworkConfigDir)

	// - /etc/containers/registries.conf
	// - /etc/containers/registries.conf.d
	// - ~/.config/containers/registries.conf
	// - ~/.config/containers/registries.conf.d
	ctx := &types.SystemContext{}
	registries, _ := ciConfig.UnqualifiedSearchRegistries(ctx)
	fmt.Println("Registries:")
	for _, v := range registries {
		fmt.Printf(" - %s\n", v)
	}
}
