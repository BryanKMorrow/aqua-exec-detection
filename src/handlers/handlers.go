package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/BryanKMorrow/aqua-exec-detection/src/aqua"
	"log"
	"net/http"
)

// IndexHandler - Home route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to aqua-exec-detection"))
}

// WebhookHandler handles all incoming messages
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var audit aqua.Runtime
	err := json.NewDecoder(r.Body).Decode(&audit)
	if err != nil {
		log.Println("Failed to decode audit event from Aqua ", err)
	}
	if audit.Action == "exec" && audit.Category == "container" {
		log.Println("Creating runtime policy with the following criteria")
		log.Printf("Pod: %s  Namespace: %s\n", audit.Podname, audit.Podnamespace)
		// Prepare the runtime policy
		namespace := aqua.Variable{
			Attribute: "kubernetes.namespace",
			Value:     audit.Podnamespace,
		}
		pod := aqua.Variable{
			Attribute: "kubernetes.pod",
			Value:     audit.Podname,
		}
		var variables []aqua.Variable
		variables = append(variables, namespace, pod)
		scope := aqua.Scope{
			Expression: "v1 && v2",
			Variables:  variables,
		}
		auditing := aqua.Auditing{
			Enabled:             true,
			AuditAllProcesses:   true,
			AuditProcessCmdline: true,
			AuditAllNetwork:     true,
			AuditOsUserActivity: true,
		}
		var scopes []string
		scopes = append(scopes, "Global")
		policy := aqua.Policy{
			Name:              fmt.Sprintf("exec-detection-%s", audit.Podname),
			Enabled:           true,
			Type:              "runtime.policy",
			RuntimeType:       "container",
			Enforce:           false,
			Scope:             scope,
			Auditing:          auditing,
			ApplicationScopes: scopes,
		}
		// Create Runtime Policy
		cli := aqua.NewClient()
		connected := cli.GetAuthToken()
		if !connected {
			log.Fatal("failed connection to the Aqua API, please verify your environment variables")
		}
		response := cli.CreateRuntimePolicy(policy)
		w.Write([]byte(response))
	}
}

/*
  "scope": {
    "expression": "v1 && v2",
    "variables": [
      {
        "attribute": "kubernetes.namespace",
        "value": "blog"
      },
      {
        "attribute": "kubernetes.pod",
        "value": "wp-db-6f5fccb649-ws5x7"
      }
    ]
  },
 */

/*
"scope": {
    "expression": "v1 && v2",
    "variables": [
      {
        "attribute": "kubernetes.namespace",
        "value": "blog"
      },
      {
        "attribute": "kubernetes.pod",
        "value": "wp-db-6f5fccb649-ws5x7"
      }
    ]
  },
 */
