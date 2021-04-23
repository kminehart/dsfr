package example

import "context"

func (e *ExampleDataSource) CheckTableHealth(ctx context.Context) error {
	// TODO: GET /api/table ...
	// Failed? return err
	// Passed? return nil
	return nil
}

func (e *ExampleDataSource) CheckStatsHealth(ctx context.Context) error {
	// TODO: GET /api/stats ...
	// Failed? return err
	// Passed? return nil
	return nil
}
