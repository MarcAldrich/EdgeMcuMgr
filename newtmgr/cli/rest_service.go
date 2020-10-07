/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"html"
	"log"
	"net/http"
)

func handle_GET_image_list(w http.ResponseWriter, r *http.Request) {
	// form request
	creq := &CommandRequested{
		EndPoint:              r.RequestURI,
		AsCobraCompatibleArgs: "image list",
	}

	// create result
	cres := NewCommandResult(creq)

	// "xmit" result{request}
	CmdFromRestChan <- *cres

	// wait on result
	output := <- cres.ReturnReady

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output.Bytes())
}

func restSvcStartCmd(cmd *cobra.Command, args []string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/images/list", handle_GET_image_list)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func restSvcCmd() *cobra.Command {
	restSvcCmd := &cobra.Command{
		Use:   "restsvc",
		Short: "Provides restful service to act as an edge management device",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Starts HTTP based restful service",
		Run:   restSvcStartCmd,
	}

	restSvcCmd.AddCommand(startCmd)

	return restSvcCmd
}
