package postgresql

import (
	"github.com/gozelle/_mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsert(t *testing.T) {
	var err error
	mapper := new(MakeMapper)
	for i := 0; i < 10; i++ {
		
		err = mapper.InsertParameterBigintInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterInt8Int8(_mock.Int8())
		require.NoError(t, err)
		
		err = mapper.InsertParameterBigserialInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterSerial8Int64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterBitByte(_mock.Byte())
		require.NoError(t, err)
		
		err = mapper.InsertParameterBitVaryingInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterBooleanInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterBoolInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterBoxInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterByteaInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterCharacterInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterCharInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterCharacterVaryingInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterVarcharInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterCidrInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterCircleInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterDateInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterDoublePrecisionInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterFloat8Int64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterInetInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterIntegerInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterIntInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterInt4Int64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterIntervalInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterJsonInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterJsonbInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterLineInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterLsegInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterMacaddrInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterMacaddr8Int64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterMoneyInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterNumericInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterDecimalInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterPathInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterPgLsnInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterPgSnapshotInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterPointInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterPolygonInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterRealInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterFloat4Int64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterSmallintInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterInt2Int64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterSmallserialInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterSerial2Int64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterSerialInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterSerial4Int64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTextInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTimeInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTimeWithTimezoneInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTimezInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTimestampInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTimestampWithTimezoneInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTimestampzInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTsqueryInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTsvectorInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterTxidSnapshotInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterUuidInt64(_mock.Int64())
		require.NoError(t, err)
		
		err = mapper.InsertParameterXmlString(_mock.String())
		require.NoError(t, err)
		
	}
}
