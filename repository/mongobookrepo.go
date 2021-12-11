package repository

import (
	"books/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoBookRepository struct {
	collection mongo.Collection;
}

func NewMongoRepository(client *mongo.Client) BookRepository{
	return &MongoBookRepository{collection: *client.Database("ma").Collection("books")}
}

func (r *MongoBookRepository) GetAll() []domain.Book{
	cursor, err := r.collection.Find(context.TODO(), bson.D{})

	entities := make([]domain.Book, 0)

	if err != nil {
		return entities
	}

	for cursor.Next(context.TODO()) {
		e := domain.Book{}

		cursor.Decode(&e)

		entities = append(entities, e)
	}

	return entities;
}


func (r *MongoBookRepository) Save(b domain.Book) (domain.Book, error) {
	_, err := r.collection.InsertOne(context.TODO(),b)

	if err != nil {
		return domain.Book{}, fmt.Errorf("could not save book: %s", err.Error())
	}

	return b, nil

}
func (r *MongoBookRepository) Update(b domain.Book) (domain.Book, error) {
	fmt.Println(b)
	update := bson.D{{"$set", b}}
	_, err := r.collection.UpdateOne(context.TODO(), bson.D{{"isbn", b.Isbn}}, update)

	if err != nil {
		return b, fmt.Errorf("could not update book: %s", err.Error())
	}

	
	return b, nil
}
func (r *MongoBookRepository) GetByISBN(isbn string) (domain.Book, error) { 
	result := r.collection.FindOne(context.TODO(), bson.D{{"isbn", isbn}})

	if result.Err() != nil {
		return domain.Book{}, fmt.Errorf("could not get book %s", result.Err().Error())
	}
	b := domain.Book{}
	result.Decode(&b)
	return b, nil
}

func (r *MongoBookRepository) GetById(id interface{}) (domain.Book, error){
	result := r.collection.FindOne(context.TODO(), bson.D{{"_id", id}})

	if result.Err() != nil {
		return domain.Book{}, fmt.Errorf("could not get book %s", result.Err().Error())
	}
	b := domain.Book{}
	result.Decode(&b)
	return b, nil
}