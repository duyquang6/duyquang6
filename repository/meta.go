package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
)

type MetaData struct {
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Collection string `json:"title"`
}

type Meta interface{

}

type meta struct {
	firebaseCli *firestore.Client
}

func(s *meta) AddBucketMetadata(data MetaData) {

}

func(s *meta) RemoveBucketMetadata(ctx context.Context, collection string) error {


	// Remove bucket
	ref := s.firebaseCli.Collection(collection)
	for {
		// Get a batch of documents
		iter := ref.Limit(10).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := s.firebaseCli.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}