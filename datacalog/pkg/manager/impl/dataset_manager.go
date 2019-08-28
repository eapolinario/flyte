package impl

import (
	"context"
	"time"

	"github.com/lyft/datacatalog/pkg/manager/impl/validators"
	"github.com/lyft/datacatalog/pkg/manager/interfaces"
	"github.com/lyft/datacatalog/pkg/repositories"
	"github.com/lyft/datacatalog/pkg/repositories/transformers"
	datacatalog "github.com/lyft/datacatalog/protos/gen"

	"github.com/lyft/datacatalog/pkg/errors"
	"github.com/lyft/flytestdlib/logger"
	"github.com/lyft/flytestdlib/promutils"
	"github.com/lyft/flytestdlib/promutils/labeled"
	"github.com/lyft/flytestdlib/storage"
)

type datasetMetrics struct {
	scope                   promutils.Scope
	createSuccessCounter    labeled.Counter
	createErrorCounter      labeled.Counter
	getSuccessCounter       labeled.Counter
	getErrorCounter         labeled.Counter
	transformerErrorCounter labeled.Counter
	validationErrorCounter  labeled.Counter
	createResponseTime      labeled.StopWatch
	getResponseTime         labeled.StopWatch
	alreadyExistsCounter    labeled.Counter
	doesNotExistCounter     labeled.Counter
}

type datasetManager struct {
	repo          repositories.RepositoryInterface
	store         *storage.DataStore
	systemMetrics datasetMetrics
}

// Create a Dataset with optional metadata. If one already exists a grpc AlreadyExists err will be returned
func (dm *datasetManager) CreateDataset(ctx context.Context, request datacatalog.CreateDatasetRequest) (*datacatalog.CreateDatasetResponse, error) {
	t := dm.systemMetrics.createResponseTime.Start(ctx)
	defer t.Stop()

	err := validators.ValidateDatasetID(request.Dataset.Id)
	if err != nil {
		logger.Errorf(ctx, "Invalid create dataset request %+v err: %v", request, err)
		dm.systemMetrics.validationErrorCounter.Inc(ctx)
		return nil, err
	}

	datasetModel, err := transformers.CreateDatasetModel(request.Dataset)
	if err != nil {
		logger.Errorf(ctx, "Unable to transform create dataset request %+v err: %v", request, err)
		dm.systemMetrics.transformerErrorCounter.Inc(ctx)
		return nil, err
	}

	err = dm.repo.DatasetRepo().Create(ctx, *datasetModel)
	if err != nil {
		if errors.IsAlreadyExistsError(err) {
			logger.Warnf(ctx, "Dataset already exists key: %+v, err %v", request.Dataset, err)
			dm.systemMetrics.alreadyExistsCounter.Inc(ctx)
		} else {
			logger.Errorf(ctx, "Failed to create dataset model: %+v err: %v", datasetModel, err)
			dm.systemMetrics.createErrorCounter.Inc(ctx)
		}
		return nil, err
	}

	logger.Debugf(ctx, "Successfully created dataset %+v", request.Dataset)
	dm.systemMetrics.createSuccessCounter.Inc(ctx)
	return &datacatalog.CreateDatasetResponse{}, nil
}

// Get a Dataset with the given DatasetID if it exists. If none exist a grpc NotFound err will be returned
func (dm *datasetManager) GetDataset(ctx context.Context, request datacatalog.GetDatasetRequest) (*datacatalog.GetDatasetResponse, error) {
	t := dm.systemMetrics.getResponseTime.Start(ctx)
	defer t.Stop()

	err := validators.ValidateDatasetID(request.Dataset)
	if err != nil {
		logger.Errorf(ctx, "Invalid get dataset request %+v err: %v", request, err)
		dm.systemMetrics.validationErrorCounter.Inc(ctx)
		return nil, err
	}

	datasetKey := transformers.FromDatasetID(*request.Dataset)
	datasetModel, err := dm.repo.DatasetRepo().Get(ctx, datasetKey)

	if err != nil {
		if errors.IsDoesNotExistError(err) {
			logger.Warnf(ctx, "Dataset does not exist key: %+v, err %v", datasetKey, err)
			dm.systemMetrics.doesNotExistCounter.Inc(ctx)
		} else {
			logger.Errorf(ctx, "Unable to get dataset request %+v err: %v", request, err)
			dm.systemMetrics.getErrorCounter.Inc(ctx)
		}
		return nil, err
	}

	datasetResponse, err := transformers.FromDatasetModel(datasetModel)
	if err != nil {
		dm.systemMetrics.transformerErrorCounter.Inc(ctx)
		return nil, err
	}

	dm.systemMetrics.getSuccessCounter.Inc(ctx)
	return &datacatalog.GetDatasetResponse{
		Dataset: datasetResponse,
	}, nil
}

func NewDatasetManager(repo repositories.RepositoryInterface, store *storage.DataStore, datasetScope promutils.Scope) interfaces.DatasetManager {
	return &datasetManager{
		repo:  repo,
		store: store,
		systemMetrics: datasetMetrics{
			scope:                   datasetScope,
			createResponseTime:      labeled.NewStopWatch("create_duration", "The duration of the create dataset calls.", time.Millisecond, datasetScope, labeled.EmitUnlabeledMetric),
			getResponseTime:         labeled.NewStopWatch("get_duration", "The duration of the get dataset calls.", time.Millisecond, datasetScope, labeled.EmitUnlabeledMetric),
			createSuccessCounter:    labeled.NewCounter("create_success_count", "The number of times create dataset was called", datasetScope, labeled.EmitUnlabeledMetric),
			getSuccessCounter:       labeled.NewCounter("get_success_count", "The number of times get dataset was called", datasetScope, labeled.EmitUnlabeledMetric),
			createErrorCounter:      labeled.NewCounter("create_failed_count", "The number of times create dataset failed", datasetScope, labeled.EmitUnlabeledMetric),
			getErrorCounter:         labeled.NewCounter("get_failed_count", "The number of times get dataset failed", datasetScope, labeled.EmitUnlabeledMetric),
			transformerErrorCounter: labeled.NewCounter("transformer_failed_count", "The number of times transformations failed", datasetScope, labeled.EmitUnlabeledMetric),
			validationErrorCounter:  labeled.NewCounter("validation_failed_count", "The number of times validation failed", datasetScope, labeled.EmitUnlabeledMetric),
			alreadyExistsCounter:    labeled.NewCounter("already_exists_count", "The number of times a dataset already exists", datasetScope, labeled.EmitUnlabeledMetric),
			doesNotExistCounter:     labeled.NewCounter("does_not_exists_count", "The number of times a dataset was not found", datasetScope, labeled.EmitUnlabeledMetric),
		},
	}
}
