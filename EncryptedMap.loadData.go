package configuration

import (
	"fmt"
	"github.com/sam-caldwell/errors"
	"github.com/sam-caldwell/file"
	"strings"
)

// loadData - load the configuration data
func (em *EncryptedMap) loadData() (err error) {
	em.lock.Lock()
	defer em.lock.Unlock()
	if strings.TrimSpace(em.fileName) == "" {
		return fmt.Errorf(errors.NotInitialized)
	}
	if !file.Existsp(&em.fileName) {
		return fmt.Errorf(errors.NotFound)
	}
	var configFile file.File
	defer configFile.Close()
	if err = configFile.Open(em.fileName, rflags, perms); err != nil {
		return err
	}
	if em.data, err = configFile.ReadAll(); err != nil {
		return err
	}
	if em.data == nil || strings.TrimSpace(string(em.data)) == "" {
		return fmt.Errorf(errors.EmptySet)
	}
	return err
}
