package producer

import (
	"context"
	"post-management/pkg"
	"post-management/pkg/dto"
	"time"

	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PostProducer struct {
	ch *amqp.Channel
}

func NewPostProducer(ch *amqp.Channel) *PostProducer {
	return &PostProducer{
		ch: ch,
	}
}
func (p *PostProducer) PostCreated(data *dto.EventProducer[dto.EventPostCreatedProducer]) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = p.ch.PublishWithContext(
		ctx,
		pkg.BROKER_EXCHANGE_POST_MANAGEMENT,
		pkg.BROKER_ROUTE_POST_CREATED,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func (p *PostProducer) PostUpdated(data *dto.EventProducer[dto.EventPostUpdatedProducer]) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = p.ch.PublishWithContext(
		ctx,
		pkg.BROKER_EXCHANGE_POST_MANAGEMENT,
		pkg.BROKER_ROUTE_POST_UPDATED,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func (p *PostProducer) LikeTotal(data *dto.EventProducer[dto.EventLikeTotalProducer]) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = p.ch.PublishWithContext(
		ctx,
		pkg.BROKER_EXCHANGE_POST_MANAGEMENT,
		pkg.BROKER_ROUTE_LIKE_TOTAL,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func (p *PostProducer) CommentTotal(data *dto.EventProducer[dto.EventCommentTotalProducer]) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = p.ch.PublishWithContext(
		ctx,
		pkg.BROKER_EXCHANGE_POST_MANAGEMENT,
		pkg.BROKER_ROUTE_COMMENT_TOTAL,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
