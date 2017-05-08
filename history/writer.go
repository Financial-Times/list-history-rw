package history

import mgo "gopkg.in/mgo.v2"

// Writer writes a list body to a database
type Writer interface {
	Write(uuid string, list List, transactionID string) error
}

func newMongoDBWriter(dbURL string, database string, collection string) Writer {
	return &mongoDBWriter{
		dbURL:          dbURL,
		collectionName: collection,
		database:       database,
	}
}

type mongoDBWriter struct {
	dbURL          string
	collectionName string
	database       string
}

func (w *mongoDBWriter) Write(uuid string, list List, transactionID string) error {
	session, err := mgo.Dial(w.dbURL)
	if err != nil {
		return err
	}
	defer session.Close()
	collection := session.DB(w.database).C(w.collectionName)
	err = collection.Insert(list)
	if err != nil {
		return err
	}
	return nil
}
