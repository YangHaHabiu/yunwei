package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TasksScheduleQueueModel = (*customTasksScheduleQueueModel)(nil)

type (
	// TasksScheduleQueueModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTasksScheduleQueueModel.
	TasksScheduleQueueModel interface {
		tasksScheduleQueueModel
	}

	customTasksScheduleQueueModel struct {
		*defaultTasksScheduleQueueModel
	}
)

// NewTasksScheduleQueueModel returns a model for the database table.
func NewTasksScheduleQueueModel(conn sqlx.SqlConn) TasksScheduleQueueModel {
	return &customTasksScheduleQueueModel{
		defaultTasksScheduleQueueModel: newTasksScheduleQueueModel(conn),
	}
}
