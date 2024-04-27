// Code generated by fastssz. DO NOT EDIT.
// Hash: a63387b6d1f288be86bb43febdcd67174ea6fd3cbf23b63c8656748b54289b53
// Version: 0.1.3
package electra

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/holiman/uint256"
)

// MarshalSSZ ssz marshals the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionPayloadHeader object to a target array
func (e *ExecutionPayloadHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(648)

	// Field (0) 'ParentHash'
	dst = append(dst, e.ParentHash[:]...)

	// Field (1) 'FeeRecipient'
	dst = append(dst, e.FeeRecipient[:]...)

	// Field (2) 'StateRoot'
	dst = append(dst, e.StateRoot[:]...)

	// Field (3) 'ReceiptsRoot'
	dst = append(dst, e.ReceiptsRoot[:]...)

	// Field (4) 'LogsBloom'
	dst = append(dst, e.LogsBloom[:]...)

	// Field (5) 'PrevRandao'
	dst = append(dst, e.PrevRandao[:]...)

	// Field (6) 'BlockNumber'
	dst = ssz.MarshalUint64(dst, e.BlockNumber)

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, e.GasLimit)

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, e.GasUsed)

	// Field (9) 'Timestamp'
	dst = ssz.MarshalUint64(dst, e.Timestamp)

	// Offset (10) 'ExtraData'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.ExtraData)

	// Field (11) 'BaseFeePerGas'
	baseFeePerGas := e.BaseFeePerGas.Bytes32()
	for i := 0; i < 32; i++ {
		dst = append(dst, baseFeePerGas[31-i])
	}

	// Field (12) 'BlockHash'
	dst = append(dst, e.BlockHash[:]...)

	// Field (13) 'TransactionsRoot'
	dst = append(dst, e.TransactionsRoot[:]...)

	// Field (14) 'WithdrawalsRoot'
	dst = append(dst, e.WithdrawalsRoot[:]...)

	// Field (15) 'BlobGasUsed'
	dst = ssz.MarshalUint64(dst, e.BlobGasUsed)

	// Field (16) 'ExcessBlobGas'
	dst = ssz.MarshalUint64(dst, e.ExcessBlobGas)

	// Field (17) 'DepositReceiptsRoot'
	dst = append(dst, e.DepositReceiptsRoot[:]...)

	// Field (18) 'WithdrawalRequestsRoot'
	dst = append(dst, e.WithdrawalRequestsRoot[:]...)

	// Field (10) 'ExtraData'
	if size := len(e.ExtraData); size > 32 {
		err = ssz.ErrBytesLengthFn("ExecutionPayloadHeader.ExtraData", size, 32)
		return
	}
	dst = append(dst, e.ExtraData...)

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 648 {
		return ssz.ErrSize
	}

	tail := buf
	var o10 uint64

	// Field (0) 'ParentHash'
	copy(e.ParentHash[:], buf[0:32])

	// Field (1) 'FeeRecipient'
	copy(e.FeeRecipient[:], buf[32:52])

	// Field (2) 'StateRoot'
	copy(e.StateRoot[:], buf[52:84])

	// Field (3) 'ReceiptsRoot'
	copy(e.ReceiptsRoot[:], buf[84:116])

	// Field (4) 'LogsBloom'
	copy(e.LogsBloom[:], buf[116:372])

	// Field (5) 'PrevRandao'
	copy(e.PrevRandao[:], buf[372:404])

	// Field (6) 'BlockNumber'
	e.BlockNumber = ssz.UnmarshallUint64(buf[404:412])

	// Field (7) 'GasLimit'
	e.GasLimit = ssz.UnmarshallUint64(buf[412:420])

	// Field (8) 'GasUsed'
	e.GasUsed = ssz.UnmarshallUint64(buf[420:428])

	// Field (9) 'Timestamp'
	e.Timestamp = ssz.UnmarshallUint64(buf[428:436])

	// Offset (10) 'ExtraData'
	if o10 = ssz.ReadOffset(buf[436:440]); o10 > size {
		return ssz.ErrOffset
	}

	if o10 < 648 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (11) 'BaseFeePerGas'
	baseFeePerGasBE := make([]byte, 32)
	for i := 0; i < 32; i++ {
		baseFeePerGasBE[i] = buf[471-i]
	}
	e.BaseFeePerGas = &uint256.Int{}
	e.BaseFeePerGas.SetBytes32(baseFeePerGasBE)

	// Field (12) 'BlockHash'
	copy(e.BlockHash[:], buf[472:504])

	// Field (13) 'TransactionsRoot'
	copy(e.TransactionsRoot[:], buf[504:536])

	// Field (14) 'WithdrawalsRoot'
	copy(e.WithdrawalsRoot[:], buf[536:568])

	// Field (15) 'BlobGasUsed'
	e.BlobGasUsed = ssz.UnmarshallUint64(buf[568:576])

	// Field (16) 'ExcessBlobGas'
	e.ExcessBlobGas = ssz.UnmarshallUint64(buf[576:584])

	// Field (17) 'DepositReceiptsRoot'
	copy(e.DepositReceiptsRoot[:], buf[584:616])

	// Field (18) 'WithdrawalRequestsRoot'
	copy(e.WithdrawalRequestsRoot[:], buf[616:648])

	// Field (10) 'ExtraData'
	{
		buf = tail[o10:]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(e.ExtraData) == 0 {
			e.ExtraData = make([]byte, 0, len(buf))
		}
		e.ExtraData = append(e.ExtraData, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) SizeSSZ() (size int) {
	size = 648

	// Field (10) 'ExtraData'
	size += len(e.ExtraData)

	return
}

// HashTreeRoot ssz hashes the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionPayloadHeader object with a hasher
func (e *ExecutionPayloadHeader) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	hh.PutBytes(e.ParentHash[:])

	// Field (1) 'FeeRecipient'
	hh.PutBytes(e.FeeRecipient[:])

	// Field (2) 'StateRoot'
	hh.PutBytes(e.StateRoot[:])

	// Field (3) 'ReceiptsRoot'
	hh.PutBytes(e.ReceiptsRoot[:])

	// Field (4) 'LogsBloom'
	hh.PutBytes(e.LogsBloom[:])

	// Field (5) 'PrevRandao'
	hh.PutBytes(e.PrevRandao[:])

	// Field (6) 'BlockNumber'
	hh.PutUint64(e.BlockNumber)

	// Field (7) 'GasLimit'
	hh.PutUint64(e.GasLimit)

	// Field (8) 'GasUsed'
	hh.PutUint64(e.GasUsed)

	// Field (9) 'Timestamp'
	hh.PutUint64(e.Timestamp)

	// Field (10) 'ExtraData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(e.ExtraData))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(e.ExtraData)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
	}

	// Field (11) 'BaseFeePerGas'
	baseFeePerGas := make([]byte, 32)
	baseFeePerGasBE := e.BaseFeePerGas.Bytes32()
	for i := 0; i < 32; i++ {
		baseFeePerGas[i] = baseFeePerGasBE[31-i]
	}
	hh.PutBytes(baseFeePerGas)

	// Field (12) 'BlockHash'
	hh.PutBytes(e.BlockHash[:])

	// Field (13) 'TransactionsRoot'
	hh.PutBytes(e.TransactionsRoot[:])

	// Field (14) 'WithdrawalsRoot'
	hh.PutBytes(e.WithdrawalsRoot[:])

	// Field (15) 'BlobGasUsed'
	hh.PutUint64(e.BlobGasUsed)

	// Field (16) 'ExcessBlobGas'
	hh.PutUint64(e.ExcessBlobGas)

	// Field (17) 'DepositReceiptsRoot'
	hh.PutBytes(e.DepositReceiptsRoot[:])

	// Field (18) 'WithdrawalRequestsRoot'
	hh.PutBytes(e.WithdrawalRequestsRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}
