package sendgrid_client

import (
    "encoding/json"
    "log"
)

// DEBUG function
func prettyPrint(i interface{}) int{
    s, _ := json.MarshalIndent(i, "", "\t")
    log.Printf("[DEBUG] prettyPrint: %s", string(s))
    return 0
}


