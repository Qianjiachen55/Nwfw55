package main

import "github.com/Qianjiachen55/Nwfw55/framework"

func UserLoginController(c *framework.Context) error {
	c.Json(200,"ok, UserLoginController")
	return nil
}