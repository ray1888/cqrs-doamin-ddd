package tab

import (
	"testing"

	"github.com/cqrs-domain-ddd/v1/src/domain/tab/model"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestCanOpenANewTab(t *testing.T) {
	Convey("given Nothing", t, func() {
		Convey("When new tab ", func() {
			openTabUUID, _ := uuid.NewUUID()
			openTabItem := model.OpenTab{openTabUUID, 1, "1"}
			Convey("Then new a tab opened", func() {
				TabOpenItem := Handler{}.handle(openTabItem)
				assert.Equal(t, openTabItem.Guid, TabOpenItem.Guid)
			})
		})
	})
}
