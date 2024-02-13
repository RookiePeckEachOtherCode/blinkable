package pkg

import (
	"blinkable/server/service/api/config"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
)

type MinioCen struct {
	m *minio.Client
}

var MinioCli MinioCen

func NewMinioClient() *MinioCen {
	MinioCli.m = &config.GlobalMinioClient
	return &MinioCli
}

// UploadFile 上传文件
func (q MinioCen) UploadFile(bucketName, objectName, filePath, contentType string) (url string, err error) {
	// 打开本地文件
	file, err := os.Open(filePath)
	if err != nil {
		klog.Errorf("open filePath: %s fail: %s", filePath, err)
		return "", err
	}
	defer file.Close()
	// 上传文件到存储桶
	_, err = q.m.PutObject(context.Background(), bucketName, objectName, file, -1, minio.PutObjectOptions{ContentType: contentType})
	url = fmt.Sprintf("http://127.0.0.1:9000/%s/%s", bucketName, objectName)
	return url, err
}

// DownloadFile 下载文件
func (q MinioCen) DownloadFile(bucketName, objectName string) (*os.File, error) {
	// 确定下载文件的路径
	downloadPath := filepath.Join("../../assets", objectName)

	// 创建本地文件
	file, err := os.Create(downloadPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 下载存储桶中的文件到本地
	err = q.m.FGetObject(context.Background(), bucketName, objectName, downloadPath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully downloaded", objectName)
	return file, nil
}

// DeleteFile 删除文件
func (q MinioCen) DeleteFile(bucketName, objectName string) (bool, error) {
	// 删除存储桶中的文件
	err := q.m.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})
	if err != nil {
		klog.Errorf("remove object fail: %s", err)
		return false, err
	}

	fmt.Println("Successfully deleted", objectName)
	return true, err
}

/*
// ListObjects 列出文件
func (m *MinioClient) ListObjects(bucketName, prefix string) ([]string, error) {
	var objectNames []string

	for object := range m.Client.ListObjects(bucketName, prefix, true, nil) {
		if object.Err != nil {
			return nil, object.Err
		}

		objectNames = append(objectNames, object.Key)
	}

	return objectNames, nil
}

// GetPresignedGetObject 返回对象的url地址，有效期时间为expires
func (m *MinioClient) GetPresignedGetObject(bucketName string, objectName string, expires time.Duration) (string, error) {
	object, err := m.Client.PresignedGetObject(bucketName, objectName, expires, nil)
	if err != nil {
		log.Println("get object fail: ", err)
		return "", err
	}

	return object.String(), nil
}
*/
