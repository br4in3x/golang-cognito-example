package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// OTP handles phone verification step.
func (c *CognitoExample) OTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	otp := r.Form.Get("otp")

	user := &cognito.ConfirmSignUpInput{
		ConfirmationCode: aws.String(otp),
		Username:         aws.String(c.RegFlow.Username),
		ClientId:         aws.String(c.AppClientID),
	}

	_, err := c.CognitoClient.ConfirmSignUp(user)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, fmt.Sprintf("/otp?error=%s", err.Error()), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/username", http.StatusFound)
}
