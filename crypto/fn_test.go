package crypto

import "testing"

func TestCryptoSortKeys(t *testing.T) {
	type (
		student struct {
			Name string `json:"name"`
			Age  uint64 `json:"age"`
		}

		school struct {
			Num      uint64            `json:"num"`
			Address  string            `json:"address"`
			Students []student         `json:"students"`
			Teachers map[string]string `json:"-"`
		}
	)

	data := &school{
		Num:     377,
		Address: "安徽省合肥市蜀山区...",
		Students: []student{
			{
				Name: "A",
				Age:  19,
			},
			{
				Name: "hello",
				Age:  17,
			},
		},

		Teachers: map[string]string{
			`chinese`: `中文`,
			`english`: `英语`,
		},
	}

	//sign, err := SortKeys(data)
	t.Logf("sort: %s \r\n", SortData(data))
	//t.Logf("refresh: %s \r\n", sign)

}
