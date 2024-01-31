// mongo.go
package db

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB representa la conexión a la base de datos MongoDB.
type MongoDB struct {
	Client *mongo.Client
}

var (
	instance *MongoDB
	once     sync.Once
)

// ObtenerInstanciaMongoDB devuelve la única instancia de MongoDB.
func ObtenerInstanciaMongoDB(cadenaConexion, nombreBaseDatos string) (*MongoDB, error) {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(cadenaConexion).
			SetConnectTimeout(10 * time.Second)

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal("Error al conetarce a la base de datos:", err)
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			// Manejar el error
			log.Fatal("Error al conetarce a la base de datos:", err)
		}

		instance = &MongoDB{
			Client: client,
		}
	})

	return instance, nil
}

// CerrarConexion cierra la conexión a la base de datos.
func (db *MongoDB) CerrarConexion() error {
	return db.Client.Disconnect(context.TODO())
}
