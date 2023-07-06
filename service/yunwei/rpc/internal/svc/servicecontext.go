package svc

import (
	"ywadmin-v3/common/interceptor/metaserver"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/monitor/rpc/monitor"
	"ywadmin-v3/service/yunwei/model"
	"ywadmin-v3/service/yunwei/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config                        config.Config
	RedisClient                   *redis.Redis
	AdminRpc                      admin.Admin
	MonitorRpc                    monitor.Monitor
	AssetModel                    model.AssetModel
	ServerAffiliationModel        model.ServerAffiliationModel
	FeatureServerInfoModel        model.FeatureServerInfoModel
	PlatformModel                 model.PlatformModel
	KeyManageModel                model.KeyManageModel
	MaintainPlanModel             model.MaintainPlanModel
	MergePlanModel                model.MergePlanModel
	OpenPlanModel                 model.OpenPlanModel
	TasksModel                    model.TasksModel
	TasksTidPidModel              model.TasksTidPidModel
	TaskLogHistroyModel           model.TaskLogHistroyModel
	HotLogHistoryModel            model.HotLogHistoryModel
	ConfigFileModel               model.ConfigFileModel
	ConfigMngLogModel             model.ConfigMngLogModel
	StatServerGameInfoModel       model.StatServerGameInfoModel
	AutoOpengameRuleModel         model.AutoOpengameRuleModel
	SwitchEntranceGameserverModel model.SwitchEntranceGameserverModel
	AlarmThresholdManageModel     model.AlarmThresholdManageModel
	TasksScheduleQueueModel       model.TasksScheduleQueueModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 1024))
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:     c,
		AdminRpc:   admin.NewAdmin(zrpc.MustNewClient(c.AdminRpcConf, zrpc.WithUnaryClientInterceptor(metaserver.NameToInterceptor), zrpc.WithDialOption(dialOption))),
		MonitorRpc: monitor.NewMonitor(zrpc.MustNewClient(c.MonitorRpcConf, zrpc.WithUnaryClientInterceptor(metaserver.NameToInterceptor), zrpc.WithDialOption(dialOption))),

		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		AssetModel:                    model.NewAssetModel(sqlConn),
		ServerAffiliationModel:        model.NewServerAffiliationModel(sqlConn),
		FeatureServerInfoModel:        model.NewFeatureServerInfoModel(sqlConn),
		PlatformModel:                 model.NewPlatformModel(sqlConn),
		KeyManageModel:                model.NewKeyManageModel(sqlConn),
		MaintainPlanModel:             model.NewMaintainPlanModel(sqlConn),
		MergePlanModel:                model.NewMergePlanModel(sqlConn),
		OpenPlanModel:                 model.NewOpenPlanModel(sqlConn),
		TasksModel:                    model.NewTasksModel(sqlConn),
		TasksTidPidModel:              model.NewTasksTidPidModel(sqlConn),
		TaskLogHistroyModel:           model.NewTaskLogHistroyModel(sqlConn),
		HotLogHistoryModel:            model.NewHotLogHistoryModel(sqlConn),
		ConfigFileModel:               model.NewConfigFileModel(sqlConn),
		ConfigMngLogModel:             model.NewConfigMngLogModel(sqlConn),
		StatServerGameInfoModel:       model.NewStatServerGameInfoModel(sqlConn),
		AutoOpengameRuleModel:         model.NewAutoOpengameRuleModel(sqlConn),
		SwitchEntranceGameserverModel: model.NewSwitchEntranceGameserverModel(sqlConn),
		AlarmThresholdManageModel:     model.NewAlarmThresholdManageModel(sqlConn),
		TasksScheduleQueueModel:       model.NewTasksScheduleQueueModel(sqlConn),
	}
}
