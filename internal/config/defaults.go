/*
 * Copyright © 2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	// Config defaults.
	defaultConfigPath string = "configs/main"

	// Server defaults.
	defaultServerHost string = "post.service.durudex.local"
	defaultServerPort string = "8005"

	// TLS server defaults.
	defaultTLSEnable bool   = true
	defaultTLSCACert string = "./certs/rootCA.pem"
	defaultTLSCert   string = "./certs/post.service.durudex.local-cert.pem"
	defaultTLSKey    string = "./certs/post.service.durudex.local-key.pem"

	// Postgres database defaults.
	defaultDatabasePostgresMaxConns int32 = 20
	defaultDatabasePostgresMinConns int32 = 5
)

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables...")

	// Server defaults.
	viper.SetDefault("server.host", defaultServerHost)
	viper.SetDefault("server.port", defaultServerPort)

	// TLS server defaults.
	viper.SetDefault("server.tls.enable", defaultTLSEnable)
	viper.SetDefault("server.tls.ca-cert", defaultTLSCACert)
	viper.SetDefault("server.tls.cert", defaultTLSCert)
	viper.SetDefault("server.tls.key", defaultTLSKey)

	// Postgres database defaults.
	viper.SetDefault("database.postgres.max-conns", defaultDatabasePostgresMaxConns)
	viper.SetDefault("database.postgres.min-conns", defaultDatabasePostgresMinConns)
}
