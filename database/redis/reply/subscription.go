package reply

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/moira-alert/moira-alert"
	"github.com/moira-alert/moira-alert/database"
)

func Subscription(rep interface{}, err error) (moira.SubscriptionData, error) {
	subscription := moira.SubscriptionData{
		// TODO not sure if this is still necessary, maybe we should just convert database and forget about it
		ThrottlingEnabled: true,
	}
	bytes, err := redis.Bytes(rep, err)
	if err != nil {
		if err == redis.ErrNil {
			return subscription, database.ErrNil
		}
		return subscription, err
	}
	err = json.Unmarshal(bytes, &subscription)
	if err != nil {
		return subscription, fmt.Errorf("Failed to parse subscription json %s: %s", string(bytes), err.Error())
	}
	return subscription, nil
}

func Subscriptions(rep interface{}, err error) ([]*moira.SubscriptionData, error) {
	values, err := redis.Values(rep, err)
	if err != nil {
		if err == redis.ErrNil {
			return make([]*moira.SubscriptionData, 0), nil
		}
		return nil, err
	}
	subscriptions := make([]*moira.SubscriptionData, len(values))
	for i, value := range values {
		subscription, err2 := Subscription(value, err)
		if err2 != nil && err2 != database.ErrNil {
			return nil, err2
		} else if err2 == database.ErrNil {
			subscriptions[i] = nil
		} else {
			subscriptions[i] = &subscription
		}
	}
	return subscriptions, nil
}