package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func (c *Config) Token() string {
	if c.token == "" {
		buf := bytes.NewBuffer(nil)
		f, err := os.Open(c.tokenFilename())
		if err != nil {
			fmt.Printf("Could not open token file from %s\n", c.tokenFilename())
			os.Exit(1)
		}
		_, err = io.Copy(buf, f)
		if err != nil {
			fmt.Println(err.Error())
			f.Close()
			os.Exit(1)
		}
		f.Close()
		s := string(buf.Bytes())
		s = strings.TrimSpace(s)
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
