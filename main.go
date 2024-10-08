package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson" // Updated import path
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance struct:

// Client: A reference to the MongoDB client connection.
// Db: A reference to the MongoDB database.

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// MongoInstance is a data type in this case like int

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName // Added '/' before dbName

// A struct for a Reminder in Go would depend on what attributes or data you want to store for the reminde

// because MongoDB dont know json we need

// json:"id,omitempty"
// This tag indicates how the field ID should be serialized to JSON
// omitempty: This option tells the JSON encoder to omit this field from the output if its value is the zero value for its type (in this case, an empty string). This is useful for avoiding the inclusion of empty fields in the JSON output.
//The absence of a BSON tag for Name, Salary, and Age means that these fields will be stored in the MongoDB document using their original names in the struct. For instance, the Name field will be stored as "Name" in BSON format.

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Salary float64 `json:"salary,omitempty"`
	Age    float64 `json:"age,omitempty"`
}

func Connect() error {

	// mongo.NewClient(...):

	//This function is part of the MongoDB Go driver and is used to create a new client instance that can connect to a MongoDB database.

	// client will hold the newly created MongoDB client instance.

	// err will hold any error that occurs during the creation of the client.

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal(err) // Handle error if the client creation fails
	}

	// context.WithTimeout(...):

	//This function is part of the context package in Go and is used to create a new context that will automatically be canceled after a specified duration.

	//context.Background():

	// This function returns an empty context that is never canceled. It is often used as the base context for deriving other contexts.

	// defer statement is used to schedule a function call to be executed after the function in which it was called has complete

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx) //  This method call is used to establish a connection to the MongoDB server with the client created earlier.
	// ctx: This is the context that was created with a timeout. It is passed to the Connect method to ensure that the connection attempt respects the timeout.
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	// The code is creating a new instance of MongoInstance using a composite literal. This initializes the Client and Db fields of the struct with the values provided.

	mg = MongoInstance{
		Client: client, // This assigns the value of the variable client to the Client field of the MongoInstance struct
		Db:     db,     // This assigns the value of the variable db to the Db field of the MongoInstance struct.
	}

	return nil

}

// Fiber is a web framework written in Go, inspired by Express.js. It is designed to be fast, minimal, and expressive for building web applications and APIs.

func main() {
	// fiber.New() initializes a new Fiber app. This creates a new instance of the Fiber web server,
	app := fiber.New()

	// The Connect() function is expected to return an error value (of type error) if something goes wrong.

	// In Go, if you omit the err != nil check after calling a function that returns an error (like Connect()), the program will not handle the error correctly

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// c is a context object (fiber.Ctx) passed to the route handler function in Fiber. It represents the context of the current HTTP request and response

	//You use c to:

	//Access request data (e.g., query parameters, request body).
	// Send responses back to the client (e.g., JSON, HTML).
	// Handle request-related tasks (e.g., set status codes, cookies, etc.).

	// fiber.Ctx is the context object type provided by the Fiber framework. It encapsulates all the details of an HTTP request and response

	// A context object type (in this case, fiber.Ctx) is a structure that encapsulates all the information about an HTTP request and response within a web framework or API

	// Here’s what fiber.Ctx typically includes:
	//
	// Request Data:
	//     URL path parameters (e.g., /user/:id).
	//     Query parameters (e.g., ?name=John).
	//     Request headers (e.g., Authorization, Content-Type).
	//     Body content (e.g., JSON, form data).
	//
	// Response Control:
	//     Send data back to the client (e.g., JSON, text).
	//     Set status codes (e.g., 200 OK, 404 Not Found).
	//     Set headers, cookies, or redirect users.

	app.Get("/employee", func(c *fiber.Ctx) error {

		//When you initialize query with bson.D{}, you are creating an empty BSON document
		query := bson.D{}

		// mg.dB: This represents the database instance from your MongoInstance struct that you defined earlier. It holds a reference to the MongoDB database you are working with.
		// .Collection("employees"): This method accesses the specified collection within the database—in this case, the "employees" collection. It allows you to perform operations on that specific collection, such as querying, inserting, updating, or deleting documents.
		// cursor: This variable will hold the result of the Find operation, which is a cursor that can be used to iterate through the returned documents.

		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// This line declares a slice variable named employees that will hold Employee structs.
		var employees []Employee // Declare a slice of Employee

		// You are telling Go to:

		// Create a slice of type Employee.
		// Set the initial length to 0, meaning that the slice starts off empty.
		// The capacity is also set to 0, meaning theres no pre-allocated space.
		employees = make([]Employee, 0) // Initialize an empty slice of Employee with length 0

		// cursor.All(...):

		//The All method is called on the cursor object. This method reads all remaining documents from the cursor and decodes them into the provided slice.
		//The first parameter is c.Context(), which provides the context for the operation. This context is important for managing timeouts and cancellations.

		//&employees:

		//This is a pointer to the employees slice where the retrieved documents will be stored. The & operator is used to pass the address of the employees slice so that it can be modified directly by the All method

		// // Read all documents from the cursor into the employees slice

		// converting data from a serialized format (like JSON or BSON) into a native data structure (like a struct or object in your programming language

		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())

		}

		// JSON Method:

		// The JSON method is a built-in function of the fiber.Ctx context that takes a Go data structure (in this case, a slice of Employee structs) and serializes it into JSON format. It then writes that JSON to the HTTP response body or Marshalling

		return c.JSON(employees)

	})

	// POSTTTTTTTTTTTTT

	app.Post("/employee", func(c *fiber.Ctx) error {

		collection := mg.Db.Collection("employees")
		//  a common way to create a new instance of a struct in Go.
		employee := new(Employee)
		//  is like opening the box and putting all the toys (data) into your special toy (the Employee struct) so you can use them later.
		// BodyParser is like a translator: It takes the JSON data from the request and fills in your Employee struct with that data, making it easy to use in your application.
		//When you send JSON data in a POST request, the server parses this data into a predefined struct (like Employee) that defines what the data looks like.
		//This allows the server to understand the data structure, making it easier to work with and ensuring that the data is in the correct format.
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		employee.ID = ""

		// collection.InsertOne(...): This function inserts a single document into the specified collection. The first argument is the context (c.Context()) for managing request deadlines and cancellations, and the second argument is the document you want to insert (employee).

		insertionResult, err := collection.InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// You're creating a filter to query a MongoDB collection based on the _id of the document you just inserted.

		// Example:- filter := bson.D{{Key: "_id", Value: id}}

		// Ordered List: When you put toys in the box, you can say, First is the teddy bear, then the car, and last is the robot. In programming, bson.D helps you keep things in that order so the computer knows exactly what you want.
		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

		// FindOne:

		// This method is part of the MongoDB Go driver. It is used to retrieve a single document from the specified collection that matches the provided filter criteria
		// The line createdRecord := collection.FindOne(c.Context(), filter) is a MongoDB query that attempts to find a single document in the specified collection that matches the provided filter criteria, using the context for operation management
		createdRecord := collection.FindOne(c.Context(), filter)

		//In summary, the line createdEmployee := &Employee{} initializes a pointer to a new Employee struct with default values, preparing it for subsequent operations, such as parsing data or inserting it into a database.
		//The curly braces {} create a new Employee struct with default values for its fields. Since the fields are not explicitly initialized, they will have zero values:
		//The & operator is used to get a pointer to the newly created Employee instance. This means that createdEmployee will point to the memory location where the Employee struct is stored, allowing you to modify its fields directly
		createdEmployee := &Employee{}

		//Decode is a magic tool that helps you organize data from a box (the database) into your own toy box (your Go struct).
		//When you get data from the database, it comes in a box, and you need to use Decode to put that data into a format you can easily use.
		//It matches the data from the box to the correct spots in your toy box based on special labels (struct tags).
		//This way, you can play with your data (use it in your application) just like you play with your toys
		createdRecord.Decode(createdEmployee)

		return c.Status(201).JSON(createdEmployee)

	})

	// PUTTTT
	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")

		employeeid, err := primitive.ObjectIDFromHex(idParam)

		if err != nil {
			return c.SendStatus(400)
		}

		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: employeeid}}

		// $set is a special MongoDB operator that updates the specified fields in a document. If a field does not exist, it will create it.

		// The inner bson.D contains the fields you want to update: name, age, and salary. Each field's new value is taken from the employee struct.

		update := bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{Key: "name", Value: employee.Name},
					{Key: "age", Value: employee.Age},
					{Key: "salary", Value: employee.Salary},
				},
			},
		}

		// mg.Db.Collection("employees"):
		//This part accesses the employees collection in your MongoDB database (mg.Db is an instance of your database). Youre specifying which collection you want to work with.

		//FindOneAndUpdate(...):
		//This method is a MongoDB operation that finds a single document based on a query and updates it. It combines the actions of finding a document and updating it into one operation, which is more efficient than doing them separately.

		// query: This is the filter used to find the document you want to update. It's usually in the form of a bson.D structure that defines the criteria for finding the document.
		//update: This specifies the update operation you want to perform (in this case, the fields to update, as defined in your previous update variable).

		//How It Works

		//Finding the Document:
		//The FindOneAndUpdate function first searches for a document in the employees collection that matches the query.

		//Updating the Document:
		//If a matching document is found, it updates that document according to the update variable.

		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(400)
			}

			return c.SendStatus(500)
		}

		employee.ID = idParam

		return c.Status(200).JSON(employee)

	})

	// Deletee
	app.Delete("/employee/:id", func(c *fiber.Ctx) error {

		// primitive.ObjectIDFromHex is a method from the go.mongodb.org/mongo-driver/bson/primitive package that converts a hexadecimal string (like MongoDB's _id) into an ObjectID

		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

		if err != nil {
			return c.SendStatus(400)
		}

		// Value: employeeID:

		// employeeID is the ObjectID that was previously converted from the string using primitive.ObjectIDFromHex(). This is the value you're using to search for a document with a matching _id.

		query := bson.D{{Key: "_id", Value: employeeID}}

		// &query: This is a pointer to the query document that specifies which document to delete. The query is typically in BSON format, and you're passing it by reference (hence the &)
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), &query)

		if err != nil {
			return c.SendStatus(500)
		}

		// If DeletedCount is less than 1, it means the deletion did not find a matching document (i.e., no document was deleted).

		if result.DeletedCount < 1 {
			return c.SendStatus(404)
		}

		// Return success response
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Employee successfully deleted",
		})

	})

	log.Fatal(app.Listen(":3000"))

}
