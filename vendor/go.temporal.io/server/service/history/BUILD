load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "history",
    srcs = [
        "commandChecker.go",
        "fx.go",
        "handler.go",
        "historyEngine.go",
        "historyEngineFactory.go",
        "historyEngineInterfaces.go",
        "historyEngineInterfaces_mock.go",
        "nDCStandbyTaskUtil.go",
        "nDCTaskUtil.go",
        "queueAckMgr.go",
        "queueFactoryBase.go",
        "queueProcessorBase.go",
        "service.go",
        "taskAllocator.go",
        "timerQueueAckMgr.go",
        "timerQueueActiveProcessor.go",
        "timerQueueActiveTaskExecutor.go",
        "timerQueueFactory.go",
        "timerQueueProcessor.go",
        "timerQueueProcessorBase.go",
        "timerQueueStandbyProcessor.go",
        "timerQueueStandbyTaskExecutor.go",
        "timerQueueTaskExecutorBase.go",
        "transferQueueActiveProcessor.go",
        "transferQueueActiveTaskExecutor.go",
        "transferQueueFactory.go",
        "transferQueueProcessor.go",
        "transferQueueProcessorBase.go",
        "transferQueueStandbyProcessor.go",
        "transferQueueStandbyTaskExecutor.go",
        "transferQueueTaskExecutorBase.go",
        "visibilityQueueFactory.go",
        "visibilityQueueProcessor.go",
        "visibilityQueueTaskExecutor.go",
        "workflowRebuilder.go",
        "workflowResetter_mock.go",
        "workflowTaskHandler.go",
        "workflowTaskHandlerCallbacks.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history",
    importpath = "go.temporal.io/server/service/history",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/pborman/uuid",
        "//vendor/go.opentelemetry.io/otel/trace",
        "//vendor/go.temporal.io/api/command/v1:command",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/failure/v1:failure",
        "//vendor/go.temporal.io/api/history/v1:history",
        "//vendor/go.temporal.io/api/query/v1:query",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/taskqueue/v1:taskqueue",
        "//vendor/go.temporal.io/api/workflow/v1:workflow",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/server/api/adminservice/v1:adminservice",
        "//vendor/go.temporal.io/server/api/clock/v1:clock",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/api/matchingservice/v1:matchingservice",
        "//vendor/go.temporal.io/server/api/namespace/v1:namespace",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/api/replication/v1:replication",
        "//vendor/go.temporal.io/server/api/token/v1:token",
        "//vendor/go.temporal.io/server/api/workflow/v1:workflow",
        "//vendor/go.temporal.io/server/client",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/archiver",
        "//vendor/go.temporal.io/server/common/archiver/provider",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/clock",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/convert",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/enums",
        "//vendor/go.temporal.io/server/common/failure",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/membership",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/payload",
        "//vendor/go.temporal.io/server/common/payloads",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/client",
        "//vendor/go.temporal.io/server/common/persistence/serialization",
        "//vendor/go.temporal.io/server/common/persistence/versionhistory",
        "//vendor/go.temporal.io/server/common/persistence/visibility",
        "//vendor/go.temporal.io/server/common/persistence/visibility/manager",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/elasticsearch",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/elasticsearch/client",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/standard/cassandra",
        "//vendor/go.temporal.io/server/common/primitives",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/common/quotas",
        "//vendor/go.temporal.io/server/common/resolver",
        "//vendor/go.temporal.io/server/common/resource",
        "//vendor/go.temporal.io/server/common/rpc/interceptor",
        "//vendor/go.temporal.io/server/common/sdk",
        "//vendor/go.temporal.io/server/common/searchattribute",
        "//vendor/go.temporal.io/server/common/serviceerror",
        "//vendor/go.temporal.io/server/common/tasks",
        "//vendor/go.temporal.io/server/common/timer",
        "//vendor/go.temporal.io/server/common/util",
        "//vendor/go.temporal.io/server/common/xdc",
        "//vendor/go.temporal.io/server/service",
        "//vendor/go.temporal.io/server/service/history/api",
        "//vendor/go.temporal.io/server/service/history/api/deleteworkflow",
        "//vendor/go.temporal.io/server/service/history/api/describemutablestate",
        "//vendor/go.temporal.io/server/service/history/api/describeworkflow",
        "//vendor/go.temporal.io/server/service/history/api/queryworkflow",
        "//vendor/go.temporal.io/server/service/history/api/reapplyevents",
        "//vendor/go.temporal.io/server/service/history/api/recordactivitytaskheartbeat",
        "//vendor/go.temporal.io/server/service/history/api/recordactivitytaskstarted",
        "//vendor/go.temporal.io/server/service/history/api/recordchildworkflowcompleted",
        "//vendor/go.temporal.io/server/service/history/api/refreshworkflow",
        "//vendor/go.temporal.io/server/service/history/api/removesignalmutablestate",
        "//vendor/go.temporal.io/server/service/history/api/replication",
        "//vendor/go.temporal.io/server/service/history/api/replicationadmin",
        "//vendor/go.temporal.io/server/service/history/api/requestcancelworkflow",
        "//vendor/go.temporal.io/server/service/history/api/resetstickytaskqueue",
        "//vendor/go.temporal.io/server/service/history/api/resetworkflow",
        "//vendor/go.temporal.io/server/service/history/api/respondactivitytaskcanceled",
        "//vendor/go.temporal.io/server/service/history/api/respondactivitytaskcompleted",
        "//vendor/go.temporal.io/server/service/history/api/respondactivitytaskfailed",
        "//vendor/go.temporal.io/server/service/history/api/signalwithstartworkflow",
        "//vendor/go.temporal.io/server/service/history/api/signalworkflow",
        "//vendor/go.temporal.io/server/service/history/api/startworkflow",
        "//vendor/go.temporal.io/server/service/history/api/terminateworkflow",
        "//vendor/go.temporal.io/server/service/history/api/verifychildworkflowcompletionrecorded",
        "//vendor/go.temporal.io/server/service/history/configs",
        "//vendor/go.temporal.io/server/service/history/consts",
        "//vendor/go.temporal.io/server/service/history/events",
        "//vendor/go.temporal.io/server/service/history/ndc",
        "//vendor/go.temporal.io/server/service/history/queues",
        "//vendor/go.temporal.io/server/service/history/replication",
        "//vendor/go.temporal.io/server/service/history/shard",
        "//vendor/go.temporal.io/server/service/history/tasks",
        "//vendor/go.temporal.io/server/service/history/vclock",
        "//vendor/go.temporal.io/server/service/history/workflow",
        "//vendor/go.temporal.io/server/service/worker/archiver",
        "//vendor/go.temporal.io/server/service/worker/parentclosepolicy",
        "//vendor/go.uber.org/fx",
        "//vendor/golang.org/x/exp/maps",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//health",
        "@org_golang_google_grpc//health/grpc_health_v1",
    ],
)
