package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/servicebus/azservicebus"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cs := os.Getenv("SERVICEBUS_CONNECTION_STRING")
	queue := os.Getenv("QUEUE_NAME")

	if cs == "" || queue == "" {
		log.Fatalf("SERVICEBUS_CONNECTION_STRING and QUEUE_NAME must be defined in the environment for this sample")
	}

	serviceBusClient, err := azservicebus.NewServiceBusClient(azservicebus.ServiceBusWithConnectionString(cs))

	if err != nil {
		log.Fatalf("Failed to create service bus client: %s", err.Error())
	}

	defer serviceBusClient.Close(context.TODO())

	//
	// Send some sample messages
	//

	sender, err := serviceBusClient.NewSender(queue)

	if err != nil {
		log.Fatalf("Failed to create the sender: %s", err.Error())
	}

	err = sender.SendMessage(context.TODO(), &azservicebus.Message{
		Body: []byte("hello, from the processor sample!"),
	})

	if err != nil {
		log.Fatalf("Failed to send sample message to queue: %s", queue)
	}

	//
	// receive the messages we've sent using the processor
	//

	processor, err := serviceBusClient.NewProcessor(
		// Will auto-complete or auto-abandon messages, based on the result from you callback
		// (this is true, by default)
		azservicebus.ProcessorWithAutoComplete(true),
		azservicebus.ProcessorWithQueue(queue),
		// or for a subscription
		// azservicebus.ProcessorWithSubscription("topic", "subscription"),
		azservicebus.ProcessorWithReceiveMode(azservicebus.ReceiveModePeekLock),
	)

	if err != nil {
		log.Fatalf("Failed to create processor: %s", err.Error())
	}

	err = processor.Start(func(message *azservicebus.ReceivedMessage) error {
		log.Printf("Received message %s", string(message.Body))

		// with auto-complete on (which it is, by default):
		// - a nil will cause us to Complete the message on the user's behalf.
		// - a non-nil error will cause to Abandon the message. The error will also be forwarded
		//   to their error handler below.
		return nil
	}, func(err error) {
		// a customer can reasonably expect to see some errors here when the processor recovers
		// from connection errors, or if some automated operations failed (like autocomplete
		// settlement failure)
		if errors.Is(err, context.Canceled) {
			// filter out errors that we expect or are not concerned about.
			return
		}

		log.Printf("Error: %s", err.Error())
	})

	if err != nil {
		log.Fatalf("Failed to start processor : %s", err.Error())
	}

	log.Printf("Waiting for 30 seconds for any messages to arrive")

	select {
	// processor.Done() can be used for a simple way to block until the processor has closed (for instance,
	// if your application was shutting down)
	case <-processor.Done():
	case <-time.After(time.Second * 30):
		break
	}

	if err := processor.Close(context.TODO()); err != nil {
		log.Fatalf("Failed to close processor: %s", err.Error())
	}

	log.Printf("Finished listening")
}
