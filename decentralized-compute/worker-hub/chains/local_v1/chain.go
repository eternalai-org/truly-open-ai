package local_v1

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"solo/internal/contracts/worker_hub"
	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg"
	"solo/pkg/eth"
	"solo/pkg/lighthouse"
	"solo/pkg/utils"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"solo/config"
	"solo/pkg/logger"

	"go.uber.org/zap"
)

type chain struct {
	common          port.ICommon
	workerHub       *worker_hub.WorkerHub
	seizeMinerRoles map[string]bool
	task            *model.Task
}

func NewChain(ctx context.Context, c port.ICommon) (port.IChain, error) {
	wkHub, err := worker_hub.NewWorkerHub(
		c.GetWorkerHubAddress(),
		c.GetClient(),
	)
	if err != nil {
		return nil, err
	}

	return &chain{common: c, seizeMinerRoles: make(map[string]bool), workerHub: wkHub}, nil
}

func (b *chain) SetTask(task *model.Task) {
	b.task = task
}

func (b *chain) getAssigmentInfo(ctx context.Context, assignmentId *big.Int, out chan model.AssimentChan) {
	var err error
	assignment := new(model.Assiment)
	defer func() {
		out <- model.AssimentChan{
			AssismentID: assignmentId,
			Err:         err,
			Data:        assignment,
		}
	}()

	_assignment, err := b.workerHub.Assignments(nil, assignmentId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("WorkerHub.Assignments",
			zap.String("task.InferenceID", assignment.InferenceId.String()),
			zap.Error(err))
		return
	}

	_b, err := json.Marshal(_assignment)
	if err != nil {
		return
	}

	err = json.Unmarshal(_b, assignment)
}

func (c *chain) GetPendingTasks(ctx context.Context, startBlock, endBlock uint64, out chan *model.Task) error {
	iter, err := c.workerHub.FilterRawSubmitted(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetPendingTasks#error", zap.Error(err))
		return err
	}

	for iter.Next() {
		event := &model.Event{}
		err := utils.Copy(event, iter.Event)
		if err != nil {
			return err
		}

		task, err := c.getPendingTask(ctx, event, startBlock, endBlock)
		if err != nil {
			return err
		}

		// push to channel
		out <- task
	}

	return nil
}

func (b *chain) seizeMinerRole(ctx context.Context, task *model.Task, assignment *model.Assiment) (*types.Transaction, error) {
	var err error
	assignmentId := task.AssignmentID
	auth, err := eth.CreateBindTransactionOpts(ctx, b.common.GetClient(), b.common.GetPrivateKey(), 200_000)
	tx := new(types.Transaction)

	if err != nil {
		return nil, err
	}

	auth.GasLimit = b.common.GetGasLimit()
	auth.GasPrice = auth.GasPrice.Mul(auth.GasPrice, big.NewInt(2))

	id, _ := big.NewInt(0).SetString(assignmentId, 10)
	tx, err = b.workerHub.SeizeMinerRole(auth, id)
	if err != nil {
		return nil, err
	}

	err = eth.WaitForTx(b.common.GetClient(), tx.Hash())
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// Implement
func (b *chain) getPendingTask(ctx context.Context, event *model.Event, startBlock, endBlock uint64) (*model.Task, error) {
	var err error
	requestId := event.InferenceId
	requestIdStr := requestId.String()
	_ = requestIdStr
	requestInfo, err := b.workerHub.GetInferenceInfo(nil, requestId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetInferenceInfo", zap.Error(err))
		return nil, err
	}

	assignmentIds, err := b.workerHub.GetAssignmentsByInference(nil, requestId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetAssignmentsByInference", zap.Error(err))
		return nil, err
	}

	respChan := make(chan model.AssimentChan)
	for _, assignmentId := range assignmentIds {
		go b.getAssigmentInfo(ctx, assignmentId, respChan)
	}
	// here

	task := &model.Task{}
	for range assignmentIds {

		assignmentFChan := <-respChan
		if assignmentFChan.Err != nil {
			continue
		}

		assignment := assignmentFChan.Data
		assignmentId := assignmentFChan.AssismentID

		task = &model.Task{
			TaskID:         assignment.InferenceId.String(),
			AssignmentID:   assignmentId.String(),
			ModelContract:  strings.ToLower(event.Model.Hex()),
			Params:         string(requestInfo.Input), // here
			Requestor:      strings.ToLower(requestInfo.Creator.Hex()),
			ZKSync:         true,
			InferenceID:    event.InferenceId.String(),
			AssignmentRole: pkg.MODE_VALIDATOR,
		}

		s, ok := b.seizeMinerRoles[task.AssignmentID]
		if ok && s == true {
			logger.GetLoggerInstanceFromContext(ctx).Info("SeizeMinerRole.Added",
				zap.String("task.TaskID", task.TaskID),
				zap.String("task.InferenceID", task.InferenceID),
				zap.String("task.AssignmentRole", task.AssignmentRole),
				zap.String("task.AssignmentID", task.AssignmentID),
				zap.String("assignment.Worker", assignment.Worker.String()),
				zap.Uint64("start_block", startBlock),
				zap.Uint64("end_block", endBlock),
			)
			continue
		}

		if !strings.EqualFold(assignment.Worker.String(), b.common.GetWalletAddres().Hex()) {
			// fmt.Println("-----> ", assignment.Worker.String())
			continue
		}

		task.AssignmentRole = pkg.MODE_MINER
		// register as a miner
		transact, err := b.seizeMinerRole(ctx, task, assignment)
		if err != nil {
			if strings.Contains(err.Error(), "revert") {
				task.AssignmentRole = pkg.MODE_VALIDATOR
			} else {
				logger.GetLoggerInstanceFromContext(ctx).Error("SeizeMinerRole.SeizeMinerRole",
					zap.String("task.TaskID", task.TaskID),
					zap.String("task.InferenceID", task.InferenceID),
					zap.String("task.AssignmentRole", task.AssignmentRole),
					zap.String("task.AssignmentID", task.AssignmentID),
					zap.String("assignment.Worker", assignment.Worker.String()),
					zap.Uint64("start_block", startBlock),
					zap.Uint64("end_block", endBlock),
					zap.Error(err))
				continue
			}
		}

		logs := []zap.Field{
			zap.String("task.InferenceID", task.InferenceID),
			zap.String("task.AssignmentRole", task.AssignmentRole),
			zap.String("task.AssignmentID", task.AssignmentID),
			zap.String("assignment.Worker", assignment.Worker.String()),
			zap.Uint64("start_block", startBlock),
			zap.Uint64("end_block", endBlock),
			zap.String("task.mode", task.AssignmentRole),
		}

		if transact != nil {
			logs = append(logs, zap.String("tx", transact.Hash().Hex()))
		}

		logger.GetLoggerInstanceFromContext(ctx).Info("SeizeMinerRole.Done", logs...)

		var batchInfers []*model.BatchInferHistory
		var externalData *model.AgentInferExternalData

		/*
			TODO - chainConfig.AgentContractAddress ???
			if chainConfig.AgentContractAddress != "" {
				isAgentInfer, batchInfers, externalData, err = s.handleNewInferIsAgentInfer(ctx, modelInfo.ModelID.String(), chainConfig, ethClient, event.Raw.TxHash, aiZKClient)
				if err != nil {
					return err
				}
			}*/

		// Detect if  is batch
		isBatch := false
		if strings.HasPrefix(string(requestInfo.Input), config.IPFSPrefix) {
			// TODO - HERE
			inputBytes, _, err := lighthouse.DownloadDataSimpleWithRetry(string(requestInfo.Input))
			if err == nil {
				batchFullPrompts := []*model.BatchInferHistory{}
				err = json.Unmarshal(inputBytes, &batchFullPrompts)
				if err != nil {
					logger.GetLoggerInstanceFromContext(ctx).Error("DownloadDataSimpleWithRetry", zap.Error(err))
				} else if len(batchFullPrompts) > 0 {
					batchInfers = batchFullPrompts
					isBatch = true
				}
			} else {
				logger.GetLoggerInstanceFromContext(ctx).Error("DownloadDataSimpleWithRetry", zap.Error(err))
			}

		}

		task.IsBatch = isBatch
		task.BatchInfers = batchInfers // here
		task.ExternalData = externalData
		b.seizeMinerRoles[task.AssignmentID] = true
		return task, nil
	}

	return nil, nil
}

func (b *chain) SubmitTask(ctx context.Context, assigmentID *big.Int, result []byte) (*types.Transaction, error) {
	auth, err := eth.CreateBindTransactionOpts(ctx, b.common.GetClient(), b.common.GetPrivateKey(), int64(b.common.GetGasLimit()))
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#CreateBindTransactionOpts",
			zap.String("worker_address", b.common.GetWalletAddres().Hex()),
			zap.Any("assigment_id", assigmentID),
			zap.Any("result", string(result)),
			zap.Error(err))
		return nil, err
	}
	mode := pkg.MODE_MINER
	auth.GasLimit = b.common.GetGasLimit()
	auth.GasPrice = auth.GasPrice.Mul(auth.GasPrice, big.NewInt(2))
	inferID := "-1"
	if b.task != nil {
		mode = b.task.AssignmentRole
		inferID = b.task.TaskID
	}

	inferIDBig, _ := big.NewInt(0).SetString(inferID, 10)

	if mode == pkg.MODE_MINER {
		tx, err := b.workerHub.SubmitSolution(auth, assigmentID, result)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#SubmitSolution",
				zap.String("worker_address", b.common.GetWalletAddres().Hex()),
				zap.Any("assigment_id", assigmentID),
				zap.Any("result", string(result)),
				zap.Error(err),
			)
			return nil, err
		}

		err = eth.WaitForTx(b.common.GetClient(), tx.Hash())
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#WaitForTx",
				zap.String("worker_address", b.common.GetWalletAddres().Hex()),
				zap.Error(err),
				zap.Any("assigment_id", assigmentID))
			return nil, err
		}

		receipt, err := b.common.GetClient().TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#TransactionReceipt",
				zap.String("worker_address", b.common.GetWalletAddres().Hex()),
				zap.Error(err),
				zap.Any("assigment_id", assigmentID))
			return nil, err
		}

		//TODO - check this
		/*if receipt.Status != ethtypes.ReceiptStatusSuccessful {
			return errors.New("tx failed")
		}

		for _, txLog := range receipt.Logs {
			feeLog, err := workerHub.WorkerHubFilterer.ParseTransferFee(*txLog)
			if err != nil {
				continue
			} else {
				if strings.EqualFold(feeLog.Miner.Hex(), tskw.address) {
					tskw.status.processedTasks++
					tskw.status.currentEarning.Add(tskw.status.currentEarning, feeLog.MingingFee)
				}
			}
		}*/
		_ = receipt

		return tx, nil
	}

	inferIdBig, _ := big.NewInt(0).SetString(inferID, 10)
	assignmentIDs, err := b.workerHub.GetAssignmentsByInference(nil, inferIdBig)
	if err != nil {
		return nil, err
	}

	//wait for task is submitted
break_here:
	for i := 1; i < 1000; i++ {
		for _, assignmentID := range assignmentIDs {
			asInfo, err := b.workerHub.GetAssignmentInfo(nil, assignmentID)
			if err != nil {
				return nil, err
			}
			if len(asInfo.Output) > 0 {
				break break_here
			}
		}

		time.Sleep(time.Second * 2)
	}

	//commit task
	auth, err = eth.CreateBindTransactionOpts(ctx, b.common.GetClient(), b.common.GetPrivateKey(), int64(b.common.GetGasLimit()))
	if err != nil {
		return nil, err
	}

	//check task is submitted or not?
	infer, err := b.workerHub.GetInferenceInfo(nil, inferIDBig)
	if err != nil {
		return nil, err
	}

	randomNonce := pkg.RandomInRange(1, 1000000000)
	commitment := b.createCommitHash(uint64(randomNonce), b.common.GetWalletAddres(), []byte(result))

	txCommit, err := b.workerHub.Commit(auth, assigmentID, commitment)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("[ERROR] executeTasks#Commit",
			zap.Any("worker_address", b.common.GetWalletAddres()),
			zap.Any("assigment_id", b.task.AssignmentID),
			zap.String("inference_id", b.task.InferenceID),
			zap.String("inference_id_1", assigmentID.String()),
			zap.Any("commitment", commitment),
			zap.Any("infer_status", infer.Status),
			zap.String("err", err.Error()))
		return nil, err
	}

	err = eth.WaitForTx(b.common.GetClient(), txCommit.Hash())
	if err != nil {
		return nil, err
	}

	auth, err = eth.CreateBindTransactionOpts(ctx, b.common.GetClient(), b.common.GetPrivateKey(), int64(b.common.GetGasLimit()))
	if err != nil {
		return nil, err
	}

	asID, _ := big.NewInt(1).SetString(b.task.AssignmentID, 10)

	randomNonceBig := big.NewInt(int64(randomNonce))
	txReveal, err := b.workerHub.Reveal(auth, asID, randomNonceBig, result)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("[ERROR] executeTasks#Reveal",
			zap.Any("worker_address", b.common.GetWalletAddres()),
			zap.Any("assigment_id", b.task.AssignmentID),
			zap.Any("asID", asID.String()),
			zap.Any("commitment", commitment),
			zap.Any("nonce", randomNonceBig.String()),
			zap.Any("result", string(result)),
			zap.String("inference_id", b.task.InferenceID),
			zap.Any("infer_status", infer.Status),
			zap.String("err", err.Error()))
		return nil, err
	}

	err = eth.WaitForTx(b.common.GetClient(), txReveal.Hash())
	if err != nil {
		return nil, err
	}

	receipt, err := b.common.GetClient().TransactionReceipt(ctx, txReveal.Hash())
	if err != nil {
		return nil, err
	}

	_ = receipt
	return txReveal, nil

}

func (b *chain) GetInferenceByMiner() ([]*big.Int, error) {
	return []*big.Int{}, nil
}

func (b *chain) GetInferenceInfo(opt *bind.CallOpts, inferID uint64) (*model.InferInfo, error) {
	ID := big.NewInt(1).SetUint64(inferID)
	t, err := b.workerHub.GetInferenceInfo(opt, ID)
	if err != nil {
		return nil, err
	}

	_ = t
	return &model.InferInfo{}, nil
}

func (b *chain) createCommitHash(nonce uint64, sender common.Address, data []byte) [32]byte {
	// uint40
	packedData := make([]byte, 5+common.AddressLength+len(data))
	packedData[0] = byte(nonce >> 32)
	packedData[1] = byte(nonce >> 24)
	packedData[2] = byte(nonce >> 16)
	packedData[3] = byte(nonce >> 8)
	packedData[4] = byte(nonce)
	//seder
	copy(packedData[5:], sender.Bytes())
	//data
	copy(packedData[5+common.AddressLength:], data)
	return crypto.Keccak256Hash(packedData[:])
}
