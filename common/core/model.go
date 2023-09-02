/**
 * @author:       wangxuebing
 * @fileName:     model.go
 * @date:         2023/4/29 12:12
 * @description:
 */

package core

type (
	ComId struct {
		ID string `json:"id"`
	}

	ComKv struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)

func StructIdToArrayId(items []*ComId) []string {
	if items == nil {
		return nil
	}
	ids := make([]string, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}
