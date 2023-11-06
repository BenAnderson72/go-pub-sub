package pubsub

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var projectID string = "my-pubsub-404318"
var topicID string = "MyTopic"

func Test_publish(t *testing.T) {

	msg := "Hello World"
	writer := os.Stdout

	e := publish(writer, projectID, topicID, msg)
	// if e != nil {
	// 	fmt.Print(e)
	// }

	assert.Nil(t, e, e)

}
