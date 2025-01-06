package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg"
	"solo/pkg/lighthouse"

	"solo/config"
	"solo/pkg/logger"

	"go.uber.org/zap"
)

var TaskChecker = make(map[string]bool)

type taskWatcher struct {
	runnerLock   sync.RWMutex
	cnf          *config.Config
	IsStaked     *bool
	currentBlock uint64
	tasksQueue   chan *model.Task

	chain   port.IChain
	staking port.IStaking
	common  port.ICommon
}

func NewTasksWatcher(chain port.IChain, staking port.IStaking, common port.ICommon, cnf *config.Config) port.ITaskWatcher {
	return &taskWatcher{
		staking:    staking,
		chain:      chain,
		common:     common,
		cnf:        cnf,
		tasksQueue: make(chan *model.Task, 10),
	}
}

func (t *taskWatcher) GetPendingTasks(ctx context.Context) {
	for {
		// logger.AtLog.Info("Waiting task...")

		fBlock := t.common.FromBlock(t.currentBlock)
		tBlock := t.common.ToBlock()

		err := t.chain.GetPendingTasks(ctx, fBlock, tBlock, t.tasksQueue)
		if err != nil {
			if t.cnf.DebugMode {
				logger.GetLoggerInstanceFromContext(ctx).Error("GetPendingTasks",
					zap.Uint64("from_block", fBlock),
					zap.Uint64("to_block", tBlock),
					zap.Error(err),
				)
			}
		}

		t.currentBlock = tBlock
		time.Sleep(time.Second * pkg.TimeToWating)

	}
}

func (t *taskWatcher) ExecueteTasks(ctx context.Context) {
	for {
		task := <-t.tasksQueue
		if task == nil {
			continue
		}

		s, ok := TaskChecker[task.AssignmentID]
		if ok && s == true {
			logger.GetLoggerInstanceFromContext(ctx).Info("executeTasks.done",
				zap.Any("worker_address", t.common.GetWalletAddres()),
				zap.Any("assigment_id", task.AssignmentID),
				zap.String("inference_id", task.AssignmentID),
			)
			continue
		}

		assigmentID, ok := big.NewInt(0).SetString(task.AssignmentID, 10)
		if !ok {
			continue
		}

		taskResult, err := t.executeTasks(ctx, task)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("executeTasks",
				zap.Any("worker_address", t.common.GetWalletAddres()),
				zap.Any("assigment_id", task.AssignmentID),
				zap.String("inference_id", task.AssignmentID),
				zap.String("task_input", task.Params),
				zap.Error(err),
			)
			continue
		}

		resultData, err := json.Marshal(taskResult)
		if err != nil {
			if t.cnf.DebugMode {
				logger.GetLoggerInstanceFromContext(ctx).Error("executeTasks",
					zap.Any("worker_address", t.common.GetWalletAddres()),
					zap.Any("assigment_id", task.AssignmentID),
					zap.String("inference_id", task.AssignmentID),
					zap.Any("result_data", resultData),
					zap.String("inference_id", task.AssignmentID),
					zap.Error(err))
			}
			continue
		}

		tx, err := t.chain.SubmitTask(ctx, assigmentID, resultData)
		if err != nil {
			if t.cnf.DebugMode {
				logger.GetLoggerInstanceFromContext(ctx).Error("executeTasks",
					zap.Any("worker_address", t.common.GetWalletAddres()),
					zap.Any("assigment_id", task.AssignmentID),
					zap.String("inference_id", task.AssignmentID),
					zap.Any("result_data", resultData),
					zap.String("inference_id", task.AssignmentID),
					zap.Error(err))
			} else {
				logger.GetLoggerInstanceFromContext(ctx).Warn("executeTasks",
					zap.Any("worker_address", t.common.GetWalletAddres()),
					zap.Any("assigment_id", task.AssignmentID),
					zap.String("inference_id", task.AssignmentID),
					zap.Any("result_data", resultData),
					zap.String("inference_id", task.AssignmentID),
					zap.String("err", err.Error()))
			}
			continue
		}

		logger.GetLoggerInstanceFromContext(ctx).Info("executeTasks",
			zap.Any("worker_address", t.common.GetWalletAddres()),
			zap.Any("assigment_id", task.AssignmentID),
			zap.String("inference_id", task.InferenceID),
			zap.Any("result_data", resultData),
			zap.String("result_tx", tx.Hash().Hex()),
		)

		TaskChecker[task.AssignmentID] = true
	}
}

func (t *taskWatcher) executeTasks(ctx context.Context, task *model.Task) (*model.TaskResult, error) {
	res := &model.TaskResult{}
	result := []byte{}
	if len(task.BatchInfers) > 0 && task.IsBatch {
		for _, b := range task.BatchInfers {
			seed := pkg.CreateSeed(b.PromptInput, task.TaskID)
			obj, err := t.inferChatCompletions(ctx, b.PromptInput, "", seed)
			if err != nil {
				return nil, err
			}
			_b, err := json.Marshal(obj)
			if err != nil {
				return nil, err
			}
			b.PromptOutput = string(_b)
		}

		objJson, err := json.Marshal(task.BatchInfers)
		if err != nil {
			return nil, err
		}

		result = objJson

	} else {
		seed := pkg.CreateSeed(task.Params, task.TaskID)
		obj, err := t.inferChatCompletions(ctx, task.Params, "", seed)
		if err != nil {
			return nil, err
		}

		objJson, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}

		result = objJson

	}

	res.Storage = model.LightHouseStorageType
	res.Data = result
	ext := "txt"
	url, err := lighthouse.UploadData(t.cnf.LighthouseKey, fmt.Sprintf("%v_result.%v", task.TaskID, ext), res.Data)
	if err != nil {
		return nil, err
	}
	res.ResultURI = "ipfs://" + url
	// logger.GetLoggerInstanceFromContext(ctx).Info("executeTasks", zap.Any("res", res))
	return res, nil
}

func (t *taskWatcher) inferChatCompletions(ctx context.Context, prompt string, modelName string, seed uint64) (*model.LLMInferResponse, error) {
	var err error
	key := "InferChatCompletions"
	logs := new([]zap.Field)
	*logs = []zap.Field{
		zap.String("model", modelName),
		zap.String("seed", modelName),
		zap.String("prompt", prompt),
	}
	defer func() {
		if t.cnf.DebugMode {
			if err != nil {
				*logs = append(*logs, zap.Error(err))
				logger.GetLoggerInstanceFromContext(ctx).Error(key, *logs...)
			} else {
				logger.GetLoggerInstanceFromContext(ctx).Info(key, *logs...)
			}
		}
	}()

	_b := []byte(prompt)

	res := &model.LLMInferResponse{}
	infer := &model.LLMInferRequest{}

	err = json.Unmarshal(_b, &infer)
	if err != nil {
		return nil, err
	}

	infer.MaxToken = 512
	infer.Temperature = 0.001
	oldModel := infer.Model
	if t.common.GetConfig().ModelName == "" {
		infer.Model = "hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q8_0"
	} else {
		infer.Model = t.common.GetConfig().ModelName
	}

	url := t.cnf.ApiUrl
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", t.cnf.ApiKey)
	*logs = append(*logs, zap.Any("headers", headers))
	*logs = append(*logs, zap.Any("inferJSON", infer))
	_b, respH, st, err := pkg.HttpRequest(url, "POST", headers, infer)
	if err != nil {
		return nil, err
	}

	*logs = append(*logs, zap.Any("resp_status_code", st))
	*logs = append(*logs, zap.Any("resp_headers", respH))
	*logs = append(*logs, zap.String("resp_body", string(_b)))
	if err = json.Unmarshal(_b, res); err != nil {
		return nil, err
	}

	res.Model = oldModel
	return res, nil
}

func (t *taskWatcher) Verify() bool {
	if t.IsStaked != nil && *t.IsStaked {
		return true
	}

	isStake, err := t.staking.IsStaked()
	if err != nil {
		isStake = false
		t.IsStaked = &isStake
		logger.AtLog.Error(err)
	}
	t.IsStaked = &isStake

	return *t.IsStaked
}

func (t *taskWatcher) MakeVerify() error {
	err := t.staking.StakeForWorker()
	if err != nil {
		return err
	}

	err = t.staking.JoinForMinting()
	if err != nil {
		return err
	}
	return nil
}

func (t *taskWatcher) rejoinForMinting(ctx context.Context) error {
	err := t.staking.JoinForMinting()
	if err != nil {
		// re-join for minting
		logger.GetLoggerInstanceFromContext(ctx).Error("reJoinForMinting",
			zap.String("worker_address", t.common.GetWalletAddres().Hex()),
			zap.Error(err),
		)

		return err
	}

	logger.GetLoggerInstanceFromContext(ctx).Info("reJoinForMinting",
		zap.String("worker_address", t.common.GetWalletAddres().Hex()),
		zap.String("msg", "SUCCESS!!!"),
	)
	return nil
}
