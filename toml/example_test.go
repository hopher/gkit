package toml

import (
	"fmt"
	"log"
)

func ExampleDecodeFile() {

	type Config struct {
		Debug    bool `toml:"Debug"`
		HttpPort int  `toml:"HttpPort"`
		GRPCPort int  `toml:"GRPCPort"`
	}

	var cfg Config
	if err := DecodeFile(fmt.Sprintf("config.%s.toml", "test"), &cfg); err != nil {
		log.Fatalln(err)
	}
}
