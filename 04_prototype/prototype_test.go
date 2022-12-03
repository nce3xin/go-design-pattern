package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeepClone(t *testing.T) {
	e := &Employee{
		Name: "Bob",
		Addr: &Address{
			Province: "Beijing",
			City:     "Beijing",
			Street:   "Haidian",
		},
		Company: &Company{
			Name: "Microsoft",
			Id:   "a65f88f9-52d6-40d5-86d8-a72447462da2",
		},
		Gender: "Female",
	}

	t.Run("case1", func(t *testing.T) {
		newE := e.DeepClone()
		// same(), NotSame() 比较对象的地址是否一致
		assert.NotSame(t, newE, e)
		assert.NotSame(t, newE.Addr, e.Addr)
		assert.NotSame(t, newE.Company, e.Company)
	})
}
