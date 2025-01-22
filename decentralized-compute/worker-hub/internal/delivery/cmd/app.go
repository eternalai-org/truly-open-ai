package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"solo/config"
	"solo/internal/factory"
	"solo/internal/port"
	"solo/internal/usecase"
	"solo/pkg"
	"strings"

	"github.com/spf13/cobra"
)

var separator = " > "

type CMD struct {
	rootCmd       *cobra.Command
	clusterCMD    *usecase.CMD_CLUSTER
	localChainCMD port.ICMDLocalChain
	cnf           *config.Config
	taskWatcher   port.IMiner
	globalCmd     []*pkg.Command
	rootNodeCmd   *pkg.Command
}

func NewCMD() (*CMD, error) {
	clusterCMD, _ := usecase.NewCmdCluster()
	//localChainCMD, _ := usecase.NewCMDLocalChain() New, but we are back to the v1 one.
	//localChainCMD, _ := usecase.NewCMDLocalChainV1()
	localChainCMD := factory.NewLocalChain("v1")
	c := &CMD{
		clusterCMD: clusterCMD,
		rootCmd: &cobra.Command{
			Use:   "mycli",
			Short: "My interactive CLI application",
			Long:  `An example of an interactive CLI built with Cobra.`,
			Run:   func(cmd *cobra.Command, args []string) {},
		},
		globalCmd: []*pkg.Command{
			/*{
				Key:  pkg.COMMAND_HELP,
				Help: "Show help information",
			},*/
		},
		rootNodeCmd: &pkg.Command{
			Key:      "",
			Help:     "",
			Children: []*pkg.Command{},
		},
		localChainCMD: localChainCMD,
	}

	c.loadWatcher()
	return c, nil
}

func (c *CMD) Run() {
	commands := c.cliCommand()
	c.rootNodeCmd.Children = commands
	c.rootCmd.Run = func(cmd *cobra.Command, args []string) {

		fmt.Printf("%sWelcome to Neurons/Solo CLI!\n", pkg.Line)
		c.interactiveMode(commands, nil)
	}

	// Execute the root command
	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *CMD) buildBreadcrumb(parentNode *pkg.Command) string {
	if parentNode == nil {
		return pkg.RootNodeTxt
	}

	parentTxt := c.buildBreadcrumb(parentNode.Parent)
	if parentNode.Name == "" {
		txt := parentTxt
		return txt
	}

	txt := parentTxt + separator + parentNode.Name
	return txt
}

func (c *CMD) interactiveMode(commands []*pkg.Command, parentNode *pkg.Command) {
	reader := bufio.NewReader(os.Stdin)
	ac := c.buildMenu(commands)

	if parentNode == nil {
		ac += pkg.PrintCommandText(pkg.COMMAND_EXIT, " - Exit the application")
	} else {
		ac += pkg.PrintCommandText(pkg.COMMAND_BACK, " - Back to the "+parentNode.Name)
	}

	txt := c.buildBreadcrumb(parentNode)
	additionalTxt := ":"
	prefix := fmt.Sprintf("%s%s\n%s", pkg.Line, txt, pkg.Line)

	h := fmt.Sprintf("%sAvailable commands%s", prefix, additionalTxt)
	fmt.Println(h)
	fmt.Println(ac)

	for {
		txt1 := c.buildBreadcrumb(parentNode)
		fmt.Print(txt1, separator)
		command, err := c.ReadString(reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		node := c.findCommand(command, commands, reader)
		if node == nil {
			switch command {
			case pkg.COMMAND_EXIT:
				if parentNode == nil {
					os.Exit(1)
				} else {
					c.interactiveMode(c.rootNodeCmd.Children, c.rootNodeCmd.Parent)
				}

			default:
				fmt.Println(ac)
			}

		} else {
			if node.Function != nil {
				node.Function(reader, node)
				fmt.Println(h)
				fmt.Println(ac)
			} else if len(node.Children) != 0 {
				c.interactiveMode(node.Children, node)
			}
		}
	}
}

func (c *CMD) buildMenu(commands []*pkg.Command) string {
	ac := ""
	for _, info := range commands {
		cmd := info.Key
		help := ""
		if info.Help != "" {
			help = fmt.Sprintf(" - %s", info.Help)
		}
		///ac += fmt.Sprintf("   %-25s %s \n", cmd, help)
		ac += pkg.PrintCommandText(cmd, help)
	}

	for _, info := range c.globalCmd {
		cmd := info.Key
		help := ""
		if info.Help != "" {
			help = fmt.Sprintf(" - %s", info.Help)
		}
		///ac += fmt.Sprintf("   %-25s %s \n", cmd, help)
		ac += pkg.PrintCommandText(cmd, help)
	}

	return ac
}

func (c *CMD) loadWatcher() {
	if c.cnf == nil {
		return
	}

	if c.cnf.Rpc == "" {
		return
	}

	if c.cnf.ChainID == "" {
		return
	}

	if c.cnf.Account == "" {
		return
	}

	taskWatcher, err := factory.NewMiner(c.cnf)
	//if err != nil {
	//logger.AtLog.Fatal(err)
	//}
	if err == nil && taskWatcher != nil {
		c.taskWatcher = taskWatcher
	}
}

func (c *CMD) buildTree(commands []*pkg.Command, parent *pkg.Command) {
	if len(commands) == 0 {
		return
	}

	for _, node := range commands {
		if parent != nil {
			node.Parent = parent
		} else {
			node.Parent = c.rootNodeCmd
		}

		if len(node.Children) > 0 {
			c.buildTree(node.Children, node)
		}
	}
}

// processCommand handles user input commands
func (c *CMD) findCommand(command string, commands []*pkg.Command, reader *bufio.Reader) *pkg.Command {
	for _, val := range commands {
		if strings.EqualFold(val.Key, command) {
			return val
		}
	}
	return nil
}

func (c *CMD) ReadString(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	command := input[:len(input)-1]
	return command, nil
}

func (c *CMD) buildCommand(conf *pkg.Command, reader *bufio.Reader, verifyIn []string) (string, error) {
	var err error
	intro := os.Getenv(conf.Key)
	if intro == "" {
		intro = conf.Default
	}

	required := ""
	suffix := ""
	if conf.Required {
		required = "Required"
		suffix = " ,"
	}

	defaultText := ""
	if intro != "" {
		defaultText = fmt.Sprintf("%sDefault: %s", suffix, intro)
	}

	if required == "" && defaultText == "" {
		fmt.Print(fmt.Sprintf("> %s: ", conf.Help))
	} else {
		fmt.Print(fmt.Sprintf("> %s (%s%s): ", conf.Help, required, defaultText))
	}

	str := ""
	for {
		str, err = c.ReadString(reader)
		if err != nil {
			return "", err
		}

		if str == "" {
			if intro != "" {
				str = intro
			}
		}

		if len(verifyIn) > 0 {
			if !pkg.IsInArray(verifyIn, str) {
				continue
			}
		}

		if conf.Required && str != "" {
			break
		}

		if !conf.Required {
			break
		}

	}

	return str, nil
}

func (c *CMD) buildInputData(reader *bufio.Reader, node *pkg.Command) map[string]string {
	data := make(map[string]string)
	for _, val := range node.Children {
		str, err := c.buildCommand(val, reader, val.VerifyInArray)
		if err != nil {
			panic(err)
		}

		data[val.Key] = str
	}

	return data
}

func (c *CMD) verify() error {
	if c.taskWatcher == nil {
		err := errors.New(pkg.ErrorFillOut)
		return err
	}
	return nil
}
