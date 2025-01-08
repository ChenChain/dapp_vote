package envfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileName = "./.env"

func UpdateOrAddEnvVariable(key, value string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	lines := []string{}
	keyExists := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			// 如果找到相应的key，替换value
			line = fmt.Sprintf("%s=%s", key, value)
			keyExists = true
		}
		lines = append(lines, line)
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// 如果key不存在，添加新的key-value对
	if !keyExists {
		newLine := fmt.Sprintf("%s=%s", key, value)
		lines = append(lines, newLine)
	}

	// 打开文件以进行写入操作，清空原内容
	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return writer.Flush()
}
