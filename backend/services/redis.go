package services

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"spotify-widget-v2/models"
	"time"
)

type SpotifySession struct {
	Session models.Token
}

type RedisService struct {
	redisClient *redis.Client
}

func NewRedisService(addr string, password string) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // make configurable
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return &RedisService{
		redisClient: rdb,
	}
}

func (s *RedisService) CreateSession(token *models.Token, id string) error {
	ctx := context.Background()
	t, err := json.Marshal(&token)
	if err != nil {
		return err
	}

	err = s.redisClient.Set(ctx, "session:"+id, t, time.Hour).Err()
	s.redisClient.Expire(ctx, id, time.Hour)
	if err != nil {
		return err
	}

	return nil
}

func (s *RedisService) GetSession(id string) (*models.Token, error) {
	ctx := context.Background()
	t, err := s.redisClient.Get(ctx, "session:"+id).Result()
	if err != nil {
		return nil, err
	}

	var session *models.Token
	err = json.Unmarshal([]byte(t), &session)
	if err != nil {
		return nil, err
	}

	return session, nil
}
