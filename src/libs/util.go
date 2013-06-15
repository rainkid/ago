package libs

import (
	"dogo"
	"fmt"
)

func GetConfig(filename string, name string) string {
	config, err := dogo.NewConfig(fmt.Sprintf("src/configs/%s.yaml", filename))
	if err != nil {
		return ""
	}
	str, err := config.String(ENV(), name)
	if err != nil {
		return ""
	}
	return str
}

func ENV() string {
	config, err := dogo.NewConfig("src/configs/app.yaml")
	if err != nil {
		return "product"
	}
	env, err := config.String("base", "env")
	if err != nil {
		return "product"
	}
	return env
}
