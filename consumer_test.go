package squeue_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/toretto460/squeue"
	"github.com/toretto460/squeue/driver"
)

type ConsumerTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	driver *MockDriver
}

func (suite *ConsumerTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.driver = NewMockDriver(suite.ctrl)
}

func (suite *ConsumerTestSuite) TearDownTest() {
	suite.ctrl.Finish()
	suite.ctrl = nil
	suite.driver = nil
}

func (suite *ConsumerTestSuite) TestNewConsumer() {
	squeue.NewConsumer[*TestMessage](suite.driver)
}

func (suite *ConsumerTestSuite) TestConsumeMessages_DriverError() {
	consumer := squeue.NewConsumer[*TestMessage](suite.driver)
	ctx := context.Background()
	queue := "test-queue"

	suite.driver.
		EXPECT().
		Consume(ctx, queue).
		Return(nil, errors.New("consume error"))

	messages, err := consumer.Consume(ctx, queue)
	suite.Nil(messages)
	suite.Error(err)
}

func (suite *ConsumerTestSuite) TestConsumeMessages_OneMessageWithError() {
	consumer := squeue.NewConsumer[*TestMessage](suite.driver)
	ctx := context.Background()
	queue := "test-queue"

	dMessages := make(chan driver.Message)
	go func() {
		dMessages <- driver.Message{Error: errors.New("error in message")}
		close(dMessages)
	}()

	suite.driver.
		EXPECT().
		Consume(ctx, queue).
		Return(dMessages, nil)

	messages, err := consumer.Consume(ctx, queue)

	suite.NotNil(messages)
	suite.Nil(err)

	messageCount := 0
	for m := range messages {
		messageCount++
		suite.Error(m.Error)
	}

	suite.Equal(1, messageCount)
}

func (suite *ConsumerTestSuite) TestConsumeMessages_OneMessageUnmarshallError() {
	consumer := squeue.NewConsumer[*TestMessage](suite.driver)
	ctx := context.Background()
	queue := "test-queue"

	dMessages := make(chan driver.Message)
	go func() {
		dMessages <- driver.Message{
			Body:  []byte("invalid json"),
			ID:    "1111",
			Error: nil,
		}
		close(dMessages)
	}()

	suite.driver.
		EXPECT().
		Consume(ctx, queue).
		Return(dMessages, nil)

	messages, err := consumer.Consume(ctx, queue)

	suite.NotNil(messages)
	suite.Nil(err)

	m := <-messages

	suite.Error(m.Error)
	suite.Contains(m.Error.Error(), "invalid character")
}

func (suite *ConsumerTestSuite) TestConsumeMessages_RealWorldScenarioWithErrors() {
	consumer := squeue.NewConsumer[*TestMessage](suite.driver)
	ctx := context.Background()
	queue := "test-queue"

	dMessages := make(chan driver.Message)
	go func() {
		dMessages <- driver.Message{
			Body:  []byte(`{"name":"test message"}`),
			ID:    "1111",
			Error: nil,
		}

		dMessages <- driver.Message{
			Error: errors.New("driver error"),
		}

		dMessages <- driver.Message{
			Body:  []byte(`{"name":"test another message"}`),
			ID:    "1111",
			Error: nil,
		}

		close(dMessages)
	}()

	suite.driver.
		EXPECT().
		Consume(ctx, queue).
		Return(dMessages, nil)

	messages, err := consumer.Consume(ctx, queue)

	suite.NotNil(messages)
	suite.Nil(err)

	m := <-messages
	suite.Nil(m.Error)
	suite.NotNil(m.Content)
	suite.Equal(m.Content.Name, "test message")

	m = <-messages
	suite.Error(m.Error)
	suite.Contains(m.Error.Error(), "driver error")

	m = <-messages
	suite.Nil(m.Error)
	suite.NotNil(m.Content)
	suite.Contains(m.Content.Name, "test another message")
}

func TestConsumerTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerTestSuite))
}
