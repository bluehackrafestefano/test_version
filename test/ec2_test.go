package test

import (
	"net/http"
	"os/exec"
	"testing"
	"time"

	"runtime"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEC2Instance(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	t.Log("Testing Terraform output: instance_id")
	instanceID := terraform.Output(t, terraformOptions, "instance_id")
	assert.NotNil(t, instanceID)
	assert.NotEmpty(t, instanceID, "EC2 instance ID should not be empty")

	t.Log("Testing Terraform output: public_ip")
	publicIP := terraform.Output(t, terraformOptions, "public_ip")
	assert.NotEmpty(t, publicIP, "EC2 instance should have a public IP")

	t.Log("Testing Terraform output: instance_state")
	instanceState := terraform.Output(t, terraformOptions, "instance_state")
	assert.Equal(t, "running", instanceState, "EC2 instance should be in 'running' state")

	t.Log("Verifying EC2 instance state using AWS SDK")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("us-east-1")},
	}))
	svc := ec2.New(sess)

	describeInput := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{&instanceID},
	}

	result, err := svc.DescribeInstances(describeInput)
	assert.NoError(t, err)
	assert.Len(t, result.Reservations, 1)
	assert.Len(t, result.Reservations[0].Instances, 1)

	awsInstance := result.Reservations[0].Instances[0]
	t.Logf("Checking EC2 instance state: %s", *awsInstance.State.Name)
	assert.Equal(t, "running", *awsInstance.State.Name)

	t.Log("Checking EC2 instance tag: Name")
	var nameTag string
	for _, tag := range awsInstance.Tags {
		if *tag.Key == "Name" {
			nameTag = *tag.Value
			break
		}
	}
	expectedTagValue := "rafe-terraform-example"
	assert.Equal(t, expectedTagValue, nameTag, "EC2 instance 'Name' tag should match the expected value")

	t.Log("Testing NGINX HTTP response")
	url := "http://" + publicIP
	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}
	var resp *http.Response
	var httpErr error
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		resp, httpErr = httpClient.Get(url)
		if httpErr == nil && resp.StatusCode == 200 {
			break
		}
		time.Sleep(10 * time.Second)
	}
	assert.NoError(t, httpErr)
	assert.Equal(t, 200, resp.StatusCode, "NGINX should return HTTP 200")

	t.Logf("NGINX HTTP response test succeeded. Opening the page: %s", url)
	// Try to open the URL in the default browser (works on most OS)
	// This is best-effort and non-blocking for test success.
	go func() {
		var cmd string
		var args []string
		switch {
		case isWindows():
			cmd = "rundll32"
			args = []string{"url.dll,FileProtocolHandler", url}
		case isMac():
			cmd = "open"
			args = []string{url}
		default: // Linux
			cmd = "xdg-open"
			args = []string{url}
		}
		_ = exec.Command(cmd, args...).Start()
	}()
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func isMac() bool {
	return runtime.GOOS == "darwin"
}
