package tunnel

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
)

type SSHTunnelConfig struct {
	User string
	Pass string
	Addr string
	Port int
}

func Open(cfg SSHTunnelConfig) (con *ssh.Client) {
	var (
		agentClient agent.Agent
		conn        net.Conn
		err         error
	)

	// Establish a connection to the local ssh-agent
	if conn, err = net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err != nil {
		panic(fmt.Sprintf("System not support SSH_AUTH_SOCK, %v", err))
	}

	defer conn.Close()
	// Create a new instance of the ssh agent
	agentClient = agent.NewClient(conn)

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User:            cfg.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{},
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}

	//When there's a non empty password add the password AuthMethod
	if cfg.Pass != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return cfg.Pass, nil
		}))
	}

	con, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port), sshConfig)
	if err != nil {
		panic(fmt.Sprintf("无法连接到SSH代理服务器: %v\n", err))
	}

	return con
}
