package response

import "github.com/roskee/gotbot/entity"

type EditMessageTextResponse struct {
	Ok     bool
	Result entity.Message
}
