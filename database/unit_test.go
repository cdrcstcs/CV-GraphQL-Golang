package database
import (
	"context"
	"testing"
	"time"
	"log"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/akhil/gql/graph/model"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func setup() (*DB, *mongo.Collection, context.Context, context.CancelFunc) {
	mongoURI := "mongodb://localhost:27017"
	if mongoURI == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	db := &DB{client: client}
	coll := client.Database("graphql").Collection("jobs")
	return db, coll, ctx, cancel
}
func TestCreateJobListing(t *testing.T) {
	db, coll, ctx, cancel := setup()
	defer cancel()
	jobInfo := model.CreateJobListingInput{
		Title:       "Software Engineer",
		Description: "Develop and maintain software applications.",
		Company:     "Tech Corp",
		URL:         "http://techcorp.com/jobs/123",
	}
	job := db.CreateJobListing(jobInfo)
	assert.NotNil(t, job)
	assert.Equal(t, jobInfo.Title, job.Title)
	assert.Equal(t, jobInfo.Description, job.Description)
	assert.Equal(t, jobInfo.Company, job.Company)
	assert.Equal(t, jobInfo.URL, job.URL)
	_, err := coll.DeleteOne(ctx, bson.M{"_id": job.ID})
	assert.NoError(t, err)
}
func TestGetJob(t *testing.T) {
	db, coll, ctx, cancel := setup()
	defer cancel()
	jobInfo := model.CreateJobListingInput{
		Title:       "Test Job",
		Description: "Test Description",
		Company:     "Test Company",
		URL:         "http://test.com",
	}
	job := db.CreateJobListing(jobInfo)
	retrievedJob := db.GetJob(job.ID)
	assert.NotNil(t, retrievedJob)
	assert.Equal(t, job.Title, retrievedJob.Title)
	assert.Equal(t, job.Description, retrievedJob.Description)
	assert.Equal(t, job.Company, retrievedJob.Company)
	assert.Equal(t, job.URL, retrievedJob.URL)
	_, err := coll.DeleteOne(ctx, bson.M{"_id": job.ID})
	assert.NoError(t, err)
}
func pointerToString(s string) *string {
	return &s
}
func TestUpdateJobListing(t *testing.T) {
	db, coll, ctx, cancel := setup()
	defer cancel()
	jobInfo := model.CreateJobListingInput{
		Title:       "Old Title",
		Description: "Old Description",
		Company:     "Old Company",
		URL:         "http://old.com",
	}
	job := db.CreateJobListing(jobInfo)
	updateInfo := model.UpdateJobListingInput{
		Title:       pointerToString("New Title"),
		Description: pointerToString("New Description"),
		URL:         pointerToString("http://new.com"),
	}
	updatedJob := db.UpdateJobListing(job.ID, updateInfo)
	assert.NotNil(t, updatedJob)
	if updateInfo.Title != nil {
		assert.Equal(t, *updateInfo.Title, updatedJob.Title)
	}
	if updateInfo.Description != nil {
		assert.Equal(t, *updateInfo.Description, updatedJob.Description)
	}
	if updateInfo.URL != nil {
		assert.Equal(t, *updateInfo.URL, updatedJob.URL)
	}
	_, err := coll.DeleteOne(ctx, bson.M{"_id": job.ID})
	assert.NoError(t, err)
}
func TestDeleteJobListing(t *testing.T) {
	db, coll, ctx, cancel := setup()
	defer cancel()
	jobInfo := model.CreateJobListingInput{
		Title:       "Job to Delete",
		Description: "Description",
		Company:     "Company",
		URL:         "http://delete.com",
	}
	job := db.CreateJobListing(jobInfo)
	deleteResponse := db.DeleteJobListing(job.ID)
	assert.NotNil(t, deleteResponse)
	assert.Equal(t, job.ID, deleteResponse.DeletedJobID)
	var result model.JobListing
	err := coll.FindOne(ctx, bson.M{"_id": job.ID}).Decode(&result)
	assert.Error(t, err)
}