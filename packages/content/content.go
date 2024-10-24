package content

import (
	"time"

	"github.com/vinibgoulart/gitbook-rag/packages/space"
)

type Content struct {
	ID        string `bun:"id,pk"`
	SpaceId   string
	Space     *space.Space `bun:"rel:has-one,join:space_id=id"`
	CreatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
}
