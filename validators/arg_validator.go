package validators

import (
	"errors"
	"fmt"
	"os"

	"github.com/wata1355/go-split-command/splitter"
)

func ValidateArgs(fileArgs []string, options splitter.Options) error {
	if err := validateFileArgs(fileArgs); err != nil {
		return err
	}
	if err := validateOptions(options); err != nil {
		return err
	}
	return nil
}

func validateFileArgs(fileArgs []string) error {
	if len(fileArgs) == 0 {
		return errors.New("Please specify file as argument")
	}
	if len(fileArgs) > 1 {
		return errors.New("Please specify only one file")
	}
	filePath := fileArgs[0]
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("File %s does not exist", filePath)
		} else {
			return fmt.Errorf("Failed to access %s: %s", filePath, err)
		}
	}
	if info.Size() == 0 {
		return fmt.Errorf("File %s is empty", filePath)
	}
	return nil
}

func validateOptions(options splitter.Options) error {
	count := []bool{options.Lines > 0, options.NumFiles > 0, options.Bytes > 0}
	trueCount := 0
	for _, cond := range count {
		if cond {
			trueCount++
		}
	}
	if trueCount == 0 {
		return errors.New("Please specify a split option (-l, -n, or -b)")
	}
	if trueCount > 1 {
		return errors.New("Please specify only one split option (-l, -n, or -b)")
	}
	return nil
}
