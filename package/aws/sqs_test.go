package aws_test

import (
	"riad/package/aws"
	"testing"
)

func TestSQS(t *testing.T) {
	aws.ReceiveMessage()
}
