package configuration

// Load - Load the encrypted file into memory.
//
// The configuration file is loaded from disk and into memory.
// The data is still encrypted.  If no data is read or the file
// does not exist, this method will return an error.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (em *EncryptedMap) Load() (err error) {
	if err = em.loadData(); err != nil {
		return err
	}
	if err = em.loadIdentity(); err != nil {
		return err
	}
	return err
}
