package main

import (
	"github.com/Benbentwo/utils/util"
	"github.com/go-errors/errors"
	"log"
)

func main() {
	var myStr string

	err := util.PromptForMissingString(&myStr, "Jira Username", "What is the jira username for the robot?", false)
	if err != nil {
		_ = errors.Errorf("Error Getting %s: %s", "JiraUser", err)
	}
	log.Printf("%s:\t%s", util.ColorInfo("MyStr"), util.ColorDebug(myStr))

}
