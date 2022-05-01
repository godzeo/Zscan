package util

import (
	"Zscan/core/logger"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

// isRelative checks if a given path is a relative path
func IsRelative(filePath string) bool {
	if strings.HasPrefix(filePath, "/") || strings.Contains(filePath, ":\\") {
		return false
	}

	return true
}

// resolvePath gets the absolute path to the template by either
// looking in the current directory or checking the nuclei templates directory.
//
// Current directory is given preference over the nuclei-templates directory.
func ResolvePath(templateName string, TemplatesDirectory string) (string, error) {
	curDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	templatePath := path.Join(curDirectory, templateName)
	if _, err := os.Stat(templatePath); !os.IsNotExist(err) {
		logger.Debugf("Found template in current directory: %s\n", templatePath)

		return templatePath, nil
	}

	if TemplatesDirectory != "" {
		templatePath := path.Join(TemplatesDirectory, templateName)
		if _, err := os.Stat(templatePath); !os.IsNotExist(err) {
			logger.Debugf("Found template in nuclei-templates directory: %s\n", templatePath)

			return templatePath, nil
		}
	}

	return "", fmt.Errorf("no such path found: %s", templateName)
}

func ResolvePathIfRelative(f string) (string, error) {
	var absPath string
	var err error
	if IsRelative(f) {
		absPath, err = ResolvePath(f, "")
		if err != nil {
			return "", errors.New(fmt.Sprintf("Could not find template file '%s': %s\n", f, err))
		}
	} else {
		absPath = f
	}
	return absPath, nil
}
