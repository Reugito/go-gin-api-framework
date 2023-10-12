// init/initializer.go

package init

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gin-gonic/gin"
	"go-gin-mongo-apis/controllers"
	"go-gin-mongo-apis/services"
)

func InitApp() (*gin.Engine, *mongo.Client, *controllers.UserController, error) {
	ctx := context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		return nil, nil, nil, err
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, nil, err
	}

	fmt.Println("MongoDB connection established")

	userc := mongoclient.Database("userdb").Collection("users")
	us := services.NewUserService(userc, ctx)
	uc := controllers.New(us)
	server := gin.Default()

	return server, mongoclient, &uc, nil
}
