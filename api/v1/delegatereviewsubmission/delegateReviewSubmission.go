package delegatereviewsubmission

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/netsepio/netsepio-gateway/util/pkg/httphelper"
	"github.com/netsepio/netsepio-gateway/util/pkg/logwrapper"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateReviewSubmission")
	{
		g.POST("", deletegateReviewSubmission)
	}
}

func deletegateReviewSubmission(c *gin.Context) {
	var request DelegateReviewSubmissionRequest
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	//aptos move run --function-id '0xc333017dfac8488f3339d39c56b0b33191af5985d1335eeefa0b3bf59fc9f4a9::netsepio::delegate_submit_review' --args 'address:0xe6ff9904344bad934dd708233fcfb79aa385cd93fc424994ae0ed97fb71160c6' --args 'string:ipfs://ipfshash2' --args 'string:category' --args 'string:mydomain.com' --args 'string:full/url/link' --args 'string:type' --args 'string:tag1, tag2, tag3, tag4' --args 'string:safe' --url "https://fullnode.devnet.aptoslabs.com" --private-key "0x5d062b4e6fa5e47e114d076801dfd37e57ebbab81de4d9a5c00dda57b7c8395b" --assume-yes
	aptos_function := "0xc333017dfac8488f3339d39c56b0b33191af5985d1335eeefa0b3bf59fc9f4a9::netsepio::delegate_submit_review"
	arg1 := "address:" + request.Reviewer
	arg2 := "string:" + request.Metadata
	arg3 := "string:" + request.Category
	arg4 := "string:" + request.Domain
	arg5 := "string:" + request.SiteUrl
	arg6 := "string:" + request.SiteType
	arg7 := "string:" + request.SiteTag
	arg8 := "string:" + request.SiteSafety
	rest_url := "https://fullnode.devnet.aptoslabs.com"
	privKey := "0x5d062b4e6fa5e47e114d076801dfd37e57ebbab81de4d9a5c00dda57b7c8395b"

	cmd := exec.Command("aptos", "move", "run", "--function-id", aptos_function, "--args", arg1, "--args", arg2, "--args", arg3, "--args", arg4, "--args", arg5, "--args", arg6, "--args", arg7, "--args", arg8, "--url", rest_url, "--private-key", privKey, "--assume-yes")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Review Submission failed with %s\n", err)
	}
	outStr, errStr := stdout.String(), stderr.String()
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	var txResult DelegateReviewSubmissionPayload
	json.Unmarshal(stdout.Bytes(), &txResult)

	logwrapper.Infof("Transaction Result: %v", txResult)
	httphelper.SuccessResponse(c, "Transaction Sent, Review shall be published soon", txResult)
}

func String(out []byte) {
	panic("unimplemented")
}
