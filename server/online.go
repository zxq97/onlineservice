package server

import "context"

func startUp(ctx context.Context, uid int64) error {
	return cacheStartUp(ctx, uid)
}

func shutdown(ctx context.Context, uid int64) error {
	return cacheShutdown(ctx, uid)
}

func getOnlineAll(ctx context.Context) ([]int64, int64, error) {
	return cacheGetOnlineAll(ctx)
}
