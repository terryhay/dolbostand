package structs

import (
	"dolbostand/structs/strutcspb"
	"fmt"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
	"unsafe"
)

func TestAlignment(t *testing.T) {
	fmt.Printf("Демонстрация выравнивания структур\n\n")

	fmt.Printf("Какие размеры будут у структур, если у них два поля по байту, а одно по 8 байт?\n")

	fmt.Printf("int64, bool, bool: %d\n", unsafe.Sizeof(strutcspb.Struct1{}))
	fmt.Printf("bool, int64, bool: %d\n", unsafe.Sizeof(strutcspb.Struct2{}))
	fmt.Printf("bool, bool, int64: %d\n", unsafe.Sizeof(strutcspb.Struct3{}))

	fmt.Printf("\nОтлитчно. Повлияет ли это на сериализацию?\n")

	var (
		bytes []byte
		err   error
	)

	{
		struct1 := &strutcspb.Struct1{
			Field1: 1,
			Field2: true,
			Field3: true,
		}
		bytes, err = protojson.Marshal(struct1)
		require.NoError(t, err)

		protojson.Unmarshal(bytes, struct1)
		require.NoError(t, err)

		fmt.Printf("int64, bool, bool: %d\tData: %v\n", len(bytes), struct1)
	}
	{
		struct2 := &strutcspb.Struct2{
			Field1: true,
			Field2: 1,
			Field3: true,
		}
		bytes, err = protojson.Marshal(struct2)
		require.NoError(t, err)

		err = protojson.Unmarshal(bytes, struct2)
		require.NoError(t, err)

		fmt.Printf("bool, int64, bool: %d\tData: %v\n", len(bytes), struct2)
	}
	{
		struct3 := &strutcspb.Struct3{
			Field1: true,
			Field2: true,
			Field3: 1,
		}
		bytes, err = protojson.Marshal(struct3)
		require.NoError(t, err)

		err = protojson.Unmarshal(bytes, struct3)
		require.NoError(t, err)

		fmt.Printf("bool, bool, int64: %d\tData: %v\n", len(bytes), struct3)
	}
	fmt.Printf("\nЗанятно. А что будет, если задать число не 1, а 1000000?\n")

	struct3 := &strutcspb.Struct3{
		Field1: true,
		Field2: true,
		Field3: 1000000,
	}
	bytes, err = protojson.Marshal(struct3)
	require.NoError(t, err)

	err = protojson.Unmarshal(bytes, struct3)
	require.NoError(t, err)

	fmt.Printf("bool, bool, int64: %d\tData: %v\n", len(bytes), struct3)

	fmt.Printf("\nДлина байтового массива увеличилась. Протобаф не прост = )\n")
}
