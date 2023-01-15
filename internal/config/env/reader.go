package env

import "github.com/rs/zerolog/log"

type reader map[string]string

func (r reader) getString(key string) string {
	result, ok := r[key]
	if !ok {
		log.Fatal().Msgf("Undefined <%v> env variable", key)
	}

	return result
}
