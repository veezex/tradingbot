package env

import zeroLog "github.com/rs/zerolog/log"

type reader map[string]string

func (r reader) getString(key string) string {
	result, ok := r[key]
	if !ok {
		zeroLog.Fatal().Msgf("Undefined <%v> env variable", key)
	}

	return result
}
