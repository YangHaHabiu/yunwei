package webssh

import (
	"golang.org/x/net/proxy"
	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
)

type AuthModel int8

const (
	PASSWORD AuthModel = iota + 1
	PUBLICKEY
)

type SSHClientConfig struct {
	AuthModel AuthModel
	HostAddr  string
	User      string
	Password  string
	KeyPath   string
	JumpAddr  string
	Timeout   time.Duration
	SocksName string
	SocksPwd  string
}

func SSHClientConfigPassword(hostAddr, user, Password, jumpAddr, SocksName, SocksPwd string) *SSHClientConfig {
	return &SSHClientConfig{
		Timeout:   time.Second * 5,
		AuthModel: PASSWORD,
		HostAddr:  hostAddr,
		User:      user,
		Password:  Password,
		JumpAddr:  jumpAddr,
		SocksName: SocksName,
		SocksPwd:  SocksPwd,
	}
}

func SSHClientConfigPulicKey(hostAddr, user, keyPath, jumpAddr, SocksName, SocksPwd string) *SSHClientConfig {
	return &SSHClientConfig{
		Timeout:   time.Second * 5,
		AuthModel: PUBLICKEY,
		HostAddr:  hostAddr,
		User:      user,
		KeyPath:   keyPath,
		JumpAddr:  jumpAddr,
		SocksName: SocksName,
		SocksPwd:  SocksPwd,
	}
}

func NewSSHClient(conf *SSHClientConfig) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         conf.Timeout,
		User:            conf.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //忽略know_hosts检查
	}

	switch conf.AuthModel {
	case PASSWORD:
		config.Auth = []ssh.AuthMethod{ssh.Password(conf.Password)}
	case PUBLICKEY:
		signer, err := getKey(conf.KeyPath)
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}

	//设置代理机
	dialer, err := proxy.SOCKS5("tcp", conf.JumpAddr, &proxy.Auth{
		User: conf.SocksName, Password: conf.SocksPwd,
	}, proxy.Direct)
	if err != nil {
		return nil, err
	}

	//读取远程主机的地址
	conn, err := dialer.Dial("tcp", conf.HostAddr)
	if err != nil {
		return nil, err
	}
	c, chans, reqs, err := ssh.NewClientConn(conn, conf.HostAddr, config)

	if err != nil {
		return nil, err
	}

	return ssh.NewClient(c, chans, reqs), err
}

func getKey(keyPath string) (ssh.Signer, error) {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	return ssh.ParsePrivateKey(key)
}
