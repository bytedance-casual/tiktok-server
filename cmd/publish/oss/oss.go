package oss

import "bytes"

// Download 从指定的 oss bucket uri 下载文件
func Download(filename string) (data []byte, err error) {
	objectName := baseURL + filename
	reader, err := bucket.GetObject(objectName)
	buf := &bytes.Buffer{}
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	data = buf.Bytes()
	return data, nil
}

// Upload 向指定的 oss bucket uri 上传文件
func Upload(filename string, data []byte) error {
	reader := bytes.NewReader(data)
	objectName := baseURL + filename
	err := bucket.PutObject(objectName, reader)
	if err != nil {
		return err
	}
	return nil
}

// UploadFromPath 向指定的 oss bucket uri 上传文件
func UploadFromPath(filename string, filePath string) error {
	objectName := baseURL + filename
	err := bucket.PutObjectFromFile(objectName, filePath)
	if err != nil {
		return err
	}
	return nil
}
