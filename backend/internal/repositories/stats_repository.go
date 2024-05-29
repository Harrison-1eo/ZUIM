package repositories

import (
	"backend/internal/models"
	"database/sql"
	"math/rand"
)

type StatsRepository struct {
}

func NewStatsRepository() *StatsRepository {
	return &StatsRepository{}
}

// 查询用户在多少个房间中
func (repo *StatsRepository) GetUserRoomCount(userID uint) (int, error) {
	var count int64
	err := db.Model(models.UserRoom{}).Where("user_id = ?", userID).Count(&count).Error
	return int(count), err
}

// 查询用户有多少个好友（测试用，返回users的总数）
func (repo *StatsRepository) GetUserFriendCount(userID uint) (int, error) {
	var count int64
	err := db.Model(models.User{}).Count(&count).Error
	return int(count), err
}

// 查询最近一周的该用户发送的消息数量分别是多少（流量）
// 返回map，key是日期，value是当天发送的消息数量
// 数据库中created_at是datetime类型，date()函数可以提取日期
func (repo *StatsRepository) GetUserMessageCount(userID uint) (map[string]int, error) {
	rows, err := db.Model(models.Message{}).
		Select("date(created_at) as date, count(*) as count").
		Where("sender_id = ?", userID).
		Group("date(created_at)").
		Rows()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			println(err.Error())
		}
	}(rows)

	result := make(map[string]int)
	for rows.Next() {
		var date string
		var count int
		err := rows.Scan(&date, &count)
		if err != nil {
			return nil, err
		}
		// count 乘一个500-1000的随机数，模拟真实数据
		result[date] = count * (500 + rand.Intn(500))
	}
	return result, nil
}

// 查询在线用户（测试用，返回users的总数*0.4）
func (repo *StatsRepository) GetOnlineUserCount() (int, error) {
	var count int64
	err := db.Model(models.User{}).Count(&count).Error
	return int(count) / 3, err
}

// 查询未读消息数量（测试用，返回messages的总数*0.1）
func (repo *StatsRepository) GetUnreadMessageCount(userID uint) (int, error) {
	var count int64
	err := db.Model(models.Message{}).Count(&count).Error
	return int(count) / 10, err
}

// 查询该用户在各个房间中的消息数量分别是多少
// 返回map，key是房间ID，value是该房间中发送的消息数量
func (repo *StatsRepository) GetUserMessageCountInRooms(userID uint) (map[uint]int, error) {
	rows, err := db.Model(models.Message{}).
		Select("room_id, count(*) as count").
		Where("user_id = ?", userID).
		Group("room_id").
		Rows()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			println(err.Error())
		}
	}(rows)

	result := make(map[uint]int)
	for rows.Next() {
		var roomID uint
		var count int
		err := rows.Scan(&roomID, &count)
		if err != nil {
			return nil, err
		}
		result[roomID] = count
	}
	return result, nil
}

// 查询该用户在的各个房间中的所有消息数量分别是多少
// 返回map，key是房间ID，value是该房间中的消息数量
func (repo *StatsRepository) GetUserAllMessageCountInRooms(userID uint) (map[uint]int, error) {
	rows, err := db.Model(models.Message{}).
		Select("room_id, count(*) as count").
		Where("sender_id = ?", userID).
		Group("room_id").
		Rows()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			println(err.Error())
		}
	}(rows)

	result := make(map[uint]int)
	for rows.Next() {
		var roomID uint
		var count int
		err := rows.Scan(&roomID, &count)
		if err != nil {
			return nil, err
		}
		// count 乘一个500-1000的随机数，模拟真实数据
		result[roomID] = count
	}
	return result, nil
}
