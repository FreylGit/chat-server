package chat

import "context"

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.chatRepo.Delete(ctx, id)
	return err
}
