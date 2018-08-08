package fishyv3

import (
	"context"

	"github.com/coadler/fishyv3/internal/models"
	"github.com/coadler/fishyv3/pb"
)

// Blacklist is the GRPC route for blacklisting a user from fishy.
func (s *FishyServerImpl) Blacklist(ctx context.Context, req *pb.BlacklistRequest) (*pb.BlacklistResponse, error) {
	return &pb.BlacklistResponse{}, inTxn(ctx, s.db, func(txn models.XODB) error {
		return liftDB(
			(&models.Blacklist{
				User: req.User,
			}).Insert(txn),
			"failed to insert blacklist",
		)
	})
}

// Unblacklist is the GRPC route for unblacklisting a user from fishy.
func (s *FishyServerImpl) Unblacklist(ctx context.Context, req *pb.UnblacklistRequest) (res *pb.UnblacklistResponse, err error) {
	return res, inTxn(ctx, s.db, func(txn models.XODB) error {
		return liftDB(
			(&models.Blacklist{
				User: req.User,
			}).Delete(txn),
			"failed to delete blacklist",
		)
	})
}