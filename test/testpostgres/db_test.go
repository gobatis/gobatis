package testpostgres

import (
	"fmt"
	"testing"

	batis "github.com/gobatis/gobatis"
	"github.com/stretchr/testify/require"
)

var db *batis.DB

func init() {
	// launch postgres docker container
	fmt.Println("begin")
	defer func() {
		// close postgres docker container
		fmt.Println("end")
	}()
}

func prepareDatabase() {

}

func TestDB(t *testing.T) {
	require.NoError(t, fmt.Errorf("some error"))
}

func TestInsert(t *testing.T) {

}

func TestInsertBatch(t *testing.T) {

}

func TestQuery(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestExec(t *testing.T) {

}

func TestPaging(t *testing.T) {

}

func TestParallelQuery(t *testing.T) {

}

func TestFetchQuery(t *testing.T) {

}
