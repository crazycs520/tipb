package tipb

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
)

func TestIndexLookUpRootRoundTrip(t *testing.T) {
	exec := &Executor{
		Tp: ExecType_TypeIndexLookUpRoot,
		IndexLookupRoot: &IndexLookUp{
			IndexHandleOffsets: []uint32{1, 2},
			KeepOrder:          proto.Bool(true),
			Children: []*Executor{
				{
					Tp: ExecType_TypeIndexScan,
					IdxScan: &IndexScan{
						TableId: 1,
						IndexId: 2,
					},
				},
				{
					Tp: ExecType_TypeTableScan,
					TblScan: &TableScan{
						TableId: 3,
					},
				},
			},
		},
	}

	data, err := proto.Marshal(exec)
	if err != nil {
		t.Fatalf("marshal IndexLookUpRoot: %v", err)
	}

	var got Executor
	if err := proto.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal IndexLookUpRoot: %v", err)
	}

	if !proto.Equal(exec, &got) {
		t.Fatalf("IndexLookUpRoot round trip mismatch: got %s want %s", proto.CompactTextString(&got), proto.CompactTextString(exec))
	}
}
