package pubsub

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var projectID string = "my-pubsub-404318"
var topicID string = "MyTopic"
var subID string = "MySub"
var writer io.Writer = os.Stdout

func init() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/workspace/go-pub-sub/.service_account.json")
}

func Test_publish(t *testing.T) {

	msg := fmt.Sprintf("Hello World %d", rand.Intn(100))
	// writer := os.Stdout

	e := publish(writer, projectID, topicID, msg)
	// if e != nil {
	// 	fmt.Print(e)
	// }

	assert.Nil(t, e, e)

}

func Test_pullMsgs(t *testing.T) {

	e := pullMsgs(writer, projectID, subID)

	assert.Nil(t, e, e)

}
