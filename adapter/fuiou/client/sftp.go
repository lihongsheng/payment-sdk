package client

import (
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"net/http"
)

type Sftp struct {
	C      enum.FileHost
	client *sftp.Client
}

func NewSftp(c enum.FileHost) (*Sftp, error) {
	client, err := connectSFTP(c.Host, c.User, c.Pwd, c.Port)
	if err != nil {
		return nil, err
	}
	return &Sftp{C: c, client: client}, nil
}

// 创建 SFTP 客户端连接
func connectSFTP(host, username, password string, port int) (*sftp.Client, error) {
	// 配置 SSH 客户端
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password), // 密码认证
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥验证（生产环境需谨慎使用）
	}

	// 连接 SSH 服务器
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败: %v", err)
	}

	// 创建 SFTP 客户端
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		sshClient.Close() // 失败时关闭 SSH 连接
		return nil, fmt.Errorf("创建 SFTP 客户端失败: %v", err)
	}

	return sftpClient, nil
}

func (s *Sftp) Close() error {
	return s.client.Close()
}

func (s *Sftp) Upload(ossUrl string, remotePath string) error {
	// 1. 通过http，从oss下载文件
	// 2. 读取文件，通过文件流上传文件到sftp
	resp, err := http.Get(ossUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 检查响应状态码（200表示成功）
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，状态码: %d", resp.StatusCode)
	}
	// 创建远程文件（权限与本地保持一致）
	remoteFile, err := s.client.Create(remotePath)
	if err != nil {
		return fmt.Errorf("创建远程文件失败: %v", err)
	}
	defer remoteFile.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	_, err = remoteFile.Write(bodyBytes)
	return err
}
