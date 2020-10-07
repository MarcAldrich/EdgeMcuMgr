package cli

import (
	"bytes"
	"fmt"
	"mynewt.apache.org/newtmgr/nmxact/nmp"
)

type ApiJobAuth struct {
	AuthMethod string `json:"auth_method"`
}
func (aja *ApiJobAuth) ToString() string {
	return aja.AuthMethod
}

type CommandRequested struct {
	EndPoint              string `json:"command_endpoint_used"`
	AsCobraCompatibleArgs string `json:"as_cobra_compat_args"`
}

func (creq *CommandRequested) ToString() string {
	return fmt.Sprintf(
		"Endpoint used: %s\nAs Cobra Args: %s\n",
		creq.EndPoint,
		creq.AsCobraCompatibleArgs,
	)
}

type CommandResult struct {
	Request	*CommandRequested	`json:"command_requested"`
	ReturnReady	chan bytes.Buffer	`json:"return_data"` // returns success and provides simple wait
}

func NewCommandResult(cres *CommandRequested) *CommandResult {
	return &CommandResult{
		Request: cres,
		ReturnReady:  make(chan bytes.Buffer),
	}
}

type RestApiJob struct {
	authInfo	ApiJobAuth					// auth data
	commandRequested	CommandRequested	// command requested
	commandResult	CommandResult	// result as json if available
}


type ImageState struct {
	Image    int    `json:"image"`
	Slot     int    `json:"slot"`
	Version  string `json:"version"`
	Bootable bool   `json:"bootable"`
	Flags    string `json:"flags"`
	Hash     string `json:"hash"`
}

type cmdRspListImages struct {
	Images []nmp.ImageStateEntry `json:"images"`
}

func (is *cmdRspListImages) ToString() string {
	returnStr := "Images:\n"
	for image_idx := range is.Images {
		returnStr = fmt.Sprintf("%s\n%s\n", returnStr, is.Images[image_idx].ToString())
	}
	return returnStr
}
