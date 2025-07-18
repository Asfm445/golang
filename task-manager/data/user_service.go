package data

import (
	"context"

	"errors"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection
var JwtSecret = []byte("your_jwt_secret")

func PromoteUser(email string) error {
	_, err := userCollection.UpdateOne(context.TODO(), bson.M{"email": email}, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}

func UserRegistration(user models.User) error {
	// Check if user with same ID or email already exists
	filter := bson.M{
		"$or": []bson.M{
			{"_id": user.ID},
			{"email": user.Email},
		},
	}

	err := userCollection.FindOne(context.TODO(), filter).Err()
	if err == nil {
		return errors.New("user with this ID or email already exists")
	}
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	if count == 0 {
		user.Role = "admin" // Set role to admin if no users exist
	} else {
		user.Role = "user" // Set role to user if users
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert user
	_, err = userCollection.InsertOne(context.TODO(), user)
	return err
}

func UserLogin(email string, password string) (string, error) {
	// Find user by email
	user, err := FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role, // optional
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
		"iat":     time.Now().Unix(),
	})

	jwtToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", errors.New("internal server error")
	}
	return jwtToken, nil
}

func FindUserByEmail(email string) (models.User, error) {
	user := models.User{}
	filter := bson.M{"email": email}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}
