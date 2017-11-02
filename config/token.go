package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func (c *Config) Token() string {
	if c.token == "" {
		buf := bytes.NewBuffer(nil)
		f, err := os.Open(c.tokenFilename())
		if err != nil {
			fmt.Println("Could not open token file")
			os.Exit(1)
		}
		_, err = io.Copy(buf, f)
		if err != nil {
			fmt.Println(err)
			f.Close()
			os.Exit(1)
		}
		f.Close()
		s := string(buf.Bytes())
		c.token = s
	}

	return c.token
}

func (c *Config) tokenFilename() string {
	if c.Debug && c.json.Debug.TokenPath != "" {
		return c.json.Debug.TokenPath
	}
	return c.json.Global.TokenPath
}
