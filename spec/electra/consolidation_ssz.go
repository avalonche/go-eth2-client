// Code generated by fastssz. DO NOT EDIT.
// Hash: 17d4c9180818d70e873edf284079b326d586a16686d17c7c974a8a2fd19ec3e9
// Version: 0.1.3
package electra

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Consolidation object
func (c *Consolidation) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Consolidation object to a target array
func (c *Consolidation) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SourceIndex'
	dst = ssz.MarshalUint64(dst, uint64(c.SourceIndex))

	// Field (1) 'TargetIndex'
	dst = ssz.MarshalUint64(dst, uint64(c.TargetIndex))

	// Field (2) 'Epoch'
	dst = ssz.MarshalUint64(dst, uint64(c.Epoch))

	return
}

// UnmarshalSSZ ssz unmarshals the Consolidation object
func (c *Consolidation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 24 {
		return ssz.ErrSize
	}

	// Field (0) 'SourceIndex'
	c.SourceIndex = phase0.ValidatorIndex(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'TargetIndex'
	c.TargetIndex = phase0.ValidatorIndex(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'Epoch'
	c.Epoch = phase0.Epoch(ssz.UnmarshallUint64(buf[16:24]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Consolidation object
func (c *Consolidation) SizeSSZ() (size int) {
	size = 24
	return
}

// HashTreeRoot ssz hashes the Consolidation object
func (c *Consolidation) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Consolidation object with a hasher
func (c *Consolidation) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SourceIndex'
	hh.PutUint64(uint64(c.SourceIndex))

	// Field (1) 'TargetIndex'
	hh.PutUint64(uint64(c.TargetIndex))

	// Field (2) 'Epoch'
	hh.PutUint64(uint64(c.Epoch))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Consolidation object
func (c *Consolidation) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
