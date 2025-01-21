package usecase

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"os"
	"solo/internal/port"
	"solo/pkg"
	"strconv"
	"strings"
)

type CMD_CLUSTER struct {
	taskWatcher    port.IMiner
	configFileName string
}

type clusterGroupData struct {
	groupName  string
	clusterIDs []*big.Int
}

func NewCmdCluster() (*CMD_CLUSTER, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cnfName := fmt.Sprintf(pkg.ENV_FILE, currentDir)
	m := &CMD_CLUSTER{
		configFileName: cnfName,
	}
	return m, nil
}

func (c *CMD_CLUSTER) SetWatcher(w port.IMiner) error {
	if w == nil {
		errStr := fmt.Sprintf("%s not found. \n%s", c.configFileName, pkg.ErrorFillOut)
		err := errors.New(errStr)
		return err
	}
	c.taskWatcher = w
	return nil
}

func (c *CMD_CLUSTER) verifyMiner() error {
	errStr := fmt.Sprintf("%s not found. \n%s", c.configFileName, pkg.ErrorFillOut)
	var err error
	fName := c.configFileName
	_, err = os.Stat(fName)
	if err != nil {
		return errors.New(errStr)
	}

	_b, err := os.ReadFile(fName)
	if err != nil {
		return errors.New(errStr)
	}

	if len(_b) == 0 {
		return errors.New(errStr)
	}

	return nil
}

func (c *CMD_CLUSTER) CreateCluster(input map[string]string) (*types.Transaction, *big.Int, error) {
	version, ok := input[pkg.COMMAND_CLUSTER_CREATE_VERSION]
	if !ok {
		version = "1"
	}

	ctype, ok := input[pkg.COMMAND_CLUSTER_CREATE_TYPE]
	if !ok {
		ctype = "text"
	}

	minHardware, ok := input[pkg.COMMAND_CLUSTER_CREATE_MIN_HARDWARE]
	if !ok {
		minHardware = "1"
	}

	modelName, ok := input[pkg.COMMAND_CLUSTER_CREATE_MODEL_NAME]
	if !ok {
		err := errors.New("model name is required")
		return nil, nil, err
	}

	versionInt, err := strconv.Atoi(version)
	if err != nil {
		versionInt = 1
	}

	minHardwareInt, err := strconv.Atoi(minHardware)
	if err != nil {
		minHardwareInt = 1
	}

	cnf := c.taskWatcher.GetConfig()
	if cnf.ModelLoadBalancerAddress == "" {
		err = errors.New("`MODEL_LOAD_BALANCER_ADDRESS` is empty. \n" + pkg.ErrorFillOut)
		return nil, nil, err
	}

	if cnf.ModelCollectionAddress == "" {
		err = errors.New("`COLLECTION_ADDRESS` is empty. \n" + pkg.ErrorFillOut)
		return nil, nil, err
	}

	return c.taskWatcher.GetCluster().CreateCluster(versionInt, minHardwareInt, modelName, ctype)

}

func (c *CMD_CLUSTER) clusterGroupData(input map[string]string) (*clusterGroupData, error) {
	groupName, ok := input[pkg.COMMAND_CREATE_GROUP_NAME]
	if !ok {
		err := errors.New("`group name` is required")
		return nil, err
	}

	cids, ok := input[pkg.COMMAND_CREATE_GROUP_CLUSTER_IDS]
	if !ok {
		err := errors.New("`clusterIDs` is required")
		return nil, err
	}

	_cids := strings.Split(cids, ",")
	clusterIDs := []*big.Int{}
	for _, id := range _cids {
		clusterID, ok := big.NewInt(0).SetString(id, 10)
		if !ok {
			continue
		}
		clusterIDs = append(clusterIDs, clusterID)

	}

	return &clusterGroupData{
		groupName:  groupName,
		clusterIDs: clusterIDs,
	}, nil
}

func (c *CMD_CLUSTER) CreateAGroupOfCluster(input map[string]string) (*types.Transaction, error) {
	cData, err := c.clusterGroupData(input)
	if err != nil {
		return nil, err
	}

	return c.taskWatcher.GetCluster().CreateAGroupOfCluster(cData.groupName, cData.clusterIDs)
}

func (c *CMD_CLUSTER) AddClustersToGroup(input map[string]string) (*types.Transaction, error) {
	cData, err := c.clusterGroupData(input)
	if err != nil {
		return nil, err
	}

	return c.taskWatcher.GetCluster().AddClustersToGroup(cData.groupName, cData.clusterIDs)
}

func (c *CMD_CLUSTER) JoinCluster() (*types.Transaction, *types.Transaction, error) {
	err := c.verifyMiner()
	if err != nil {
		return nil, nil, err
	}

	cnf := c.taskWatcher.GetConfig()
	if cnf.ClusterID == "" {
		return nil, nil, errors.New("CLUSTER_ID is empty, use `config` to set it")
	}

	registerTx, joinForMintingTx, err := c.taskWatcher.MakeVerify()
	if err != nil {
		return nil, nil, err
	}

	return registerTx, joinForMintingTx, nil
}
