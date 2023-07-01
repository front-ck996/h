package csy

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileHandle struct {
}

func NewFile() *FileHandle {
	return &FileHandle{}
}

func (file *FileHandle) FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFile is_file()
func (file *FileHandle) IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDir is_dir()
func (file *FileHandle) IsDir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

// ISIMG
func (file *FileHandle) IsImg(filename string) (bool, error) {
	// Open File
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		return false, err
	}

	contentType := http.DetectContentType(buffer)
	log.Println(contentType)
	if !strings.HasPrefix(contentType, "im") {
		return false, nil
	}
	return true, nil
}

// Copy 文件
func (file *FileHandle) CopyFile(sourceFile, destinationFile string, markDir ...bool) error {
	if len(markDir) != 0 {
		os.MkdirAll(filepath.Dir(destinationFile), 0777)
	}
	src, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return nil
}

func (file *FileHandle) SplitFile(originalFile, targetFolder string, chunkSize uint64) (err error) {
	var splitFiles []string
	// 打开原始文件
	original, err := os.Open(originalFile)
	if err != nil {
		return fmt.Errorf("无法打开原始文件: %s", err)
	}
	defer original.Close()

	// 获取原始文件的信息
	_, err = original.Stat()
	if err != nil {
		return fmt.Errorf("无法获取原始文件信息: %s", err)
	}

	// 创建目标文件夹
	os.RemoveAll(targetFolder)
	err = os.MkdirAll(targetFolder, 0755)
	if err != nil {
		return fmt.Errorf("无法创建目标文件夹: %s", err)
	}

	// 缓冲区大小
	bufferSize := int(chunkSize)
	buffer := make([]byte, bufferSize)

	// 当前分片大小
	currentChunkSize := uint64(0)

	// 当前分片编号
	currentChunkNumber := 1

	// 获取原始文件的后缀名
	originalExt := filepath.Ext(originalFile)
	// 去掉后缀名的原始文件名
	originalName := strings.TrimSuffix(filepath.Base(originalFile), originalExt)

	// 创建新的分片文件
	chunkFileName := fmt.Sprintf("%s/%s_chunk%d%s", targetFolder, originalName, currentChunkNumber, originalExt)
	chunkFile, err := os.Create(chunkFileName)
	if err != nil {
		return fmt.Errorf("无法创建分片文件: %s", err)
	}
	defer chunkFile.Close()

	// 逐个字节进行分割
	for {
		// 从原始文件读取字节到缓冲区
		bytesRead, err := original.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("无法读取原始文件: %s", err)
		}

		// 如果没有更多字节可读，则结束循环
		if bytesRead == 0 {
			break
		}

		// 写入缓冲区的字节到当前分片文件
		_, err = chunkFile.Write(buffer[:bytesRead])
		if err != nil {
			return fmt.Errorf("无法写入分片文件: %s", err)
		}

		// 更新当前分片大小
		currentChunkSize += uint64(bytesRead)

		// 如果当前分片大小超过指定大小，则创建新的分片文件
		if currentChunkSize >= chunkSize {
			// 关闭当前分片文件
			err = chunkFile.Close()
			if err != nil {
				return fmt.Errorf("无法关闭分片文件: %s", err)
			}
			splitFiles = append(splitFiles, filepath.Base(chunkFileName))
			fmt.Printf("已创建分片文件: %s\n", chunkFileName)

			// 递增分片编号
			currentChunkNumber++

			// 重置当前分片大小
			currentChunkSize = 0

			// 创建新的分片文件
			chunkFileName = fmt.Sprintf("%s/%s_chunk%d%s", targetFolder, originalName, currentChunkNumber, originalExt)
			chunkFile, err = os.Create(chunkFileName)
			if err != nil {
				return fmt.Errorf("无法创建分片文件: %s", err)
			}
			defer chunkFile.Close()
		}
	}

	// 关闭最后一个分片文件
	err = chunkFile.Close()
	if err != nil {
		return fmt.Errorf("无法关闭分片文件: %s", err)
	}

	fmt.Printf("已创建分片文件: %s\n", chunkFileName)
	splitFiles = append(splitFiles, filepath.Base(chunkFileName))
	marshal, err := json.Marshal(splitFiles)
	if err == nil {
		os.WriteFile(filepath.Dir(chunkFileName)+"/files.txt", marshal, 0755)
	}
	return nil
}

// 合并分割之后的文件
func (file *FileHandle) MergeFiles(sourceFolder, targetFile string) error {
	// 创建目标文件
	target, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("无法创建目标文件: %s", err)
	}
	defer target.Close()

	// 获取源文件夹中的所有分片文件
	chunkFiles, err := filepath.Glob(fmt.Sprintf("%s/*", sourceFolder))
	if err != nil {
		return fmt.Errorf("无法读取分片文件: %s", err)
	}

	// 按照分片文件名排序
	_sortChunkFiles(chunkFiles)

	// 逐个分片文件进行合并
	for _, chunkFile := range chunkFiles {
		// 打开分片文件
		chunk, err := os.Open(chunkFile)
		if err != nil {
			return fmt.Errorf("无法打开分片文件: %s", err)
		}
		defer chunk.Close()

		// 从分片文件复制内容到目标文件
		_, err = io.Copy(target, chunk)
		if err != nil {
			return fmt.Errorf("无法复制分片文件内容: %s", err)
		}

		fmt.Printf("已合并分片文件: %s\n", chunkFile)
	}

	fmt.Printf("已还原文件: %s\n", targetFile)

	return nil
}

// 排序分片文件，按照分片编号排序
func _sortChunkFiles(chunkFiles []string) {
	sort.Slice(chunkFiles, func(i, j int) bool {
		chunkFile1 := chunkFiles[i]
		chunkFile2 := chunkFiles[j]

		// 提取分片编号
		chunkNumber1 := _getChunkNumber(chunkFile1)
		chunkNumber2 := _getChunkNumber(chunkFile2)

		return chunkNumber1 < chunkNumber2
	})
}

// 提取分片编号
func _getChunkNumber(chunkFile string) int {
	fileName := filepath.Base(chunkFile)
	fileExt := filepath.Ext(fileName)
	fileNameWithoutExt := strings.TrimSuffix(fileName, fileExt)

	var chunkNumber int
	_, err := fmt.Sscanf(fileNameWithoutExt, "chunk%d", &chunkNumber)
	if err != nil {
		return 0
	}

	return chunkNumber
}

// 解压 zip 文件
func UnZip(zipFilePath, targetDir string) error {
	// 打开 ZIP 文件
	zipFile, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return fmt.Errorf("无法打开 ZIP 文件: %s", err)
	}
	defer zipFile.Close()

	// 创建目标文件夹
	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		return fmt.Errorf("无法创建目标文件夹: %s", err)
	}

	// 遍历 ZIP 文件中的文件和文件夹
	for _, file := range zipFile.File {
		// 构建解压后的文件路径
		filePath := filepath.Join(targetDir, file.Name)

		if file.FileInfo().IsDir() {
			// 创建文件夹
			err = os.MkdirAll(filePath, file.Mode())
			if err != nil {
				return fmt.Errorf("无法创建文件夹: %s", err)
			}
			continue
		}

		// 创建解压后的文件
		outputFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return fmt.Errorf("无法创建文件: %s", err)
		}
		defer outputFile.Close()

		// 打开 ZIP 文件中的文件
		zipFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("无法打开 ZIP 文件中的文件: %s", err)
		}
		defer zipFile.Close()

		// 将 ZIP 文件中的内容复制到解压后的文件
		_, err = io.Copy(outputFile, zipFile)
		if err != nil {
			return fmt.Errorf("无法解压文件: %s", err)
		}
	}

	return nil
}
