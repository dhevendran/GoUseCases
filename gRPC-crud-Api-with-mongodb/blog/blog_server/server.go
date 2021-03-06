package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/vamsikrishnasiddu/gRPC-go-Projects/GRPC-GO-COURSE/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection
var bucket *gridfs.Bucket
var fsFiles *mongo.Collection

type server struct {
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	fmt.Println("Reading the blog Request")

	blogID := req.GetBlogId()
	oid, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}
	//create an empty struct
	data := &blogItem{}

	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID:%v", err),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: dataToBlogPb(data),
	}, nil

}

func dataToBlogPb(data *blogItem) *blogpb.Blog {
	return &blogpb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}
}
func (*server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {

	fmt.Println("List Blog request")

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error:%v", err),
		)

	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &blogItem{}
		if err := cur.Decode(data); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding the data from MongoDb:%v", err),
			)
		}
		stream.Send(&blogpb.ListBlogResponse{Blog: dataToBlogPb(data)})
		time.Sleep(1 * time.Second)
	}

	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error:%v", err),
		)
	}
	return nil
}
func (*server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {

	fmt.Println("DeleteBlog method is called...")

	blogID := req.GetBlogId()

	//Extract the id

	oid, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Unable to parse Id"),
		)
	}

	filter := bson.M{"_id": oid}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB:%v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog in MongoDB:%v", err),
		)
	}

	return &blogpb.DeleteBlogResponse{
		BlogId: req.GetBlogId(),
	}, nil
}

func (*server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Println("Updating the blog request....")

	//Extract the id from the req
	blog := req.GetBlog()
	oid, err := primitive.ObjectIDFromHex(blog.GetId())

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	data := &blogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID:%v", err),
		)

	}
	//we update our internal struct
	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()

	//we update the blog in the database

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)

	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in Mongodb:%v\n", updateErr),
		)
	}

	return &blogpb.UpdateBlogResponse{
		Blog: dataToBlogPb(data),
	}, nil

}
func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("Create blog request...")
	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Content:  blog.GetContent(),
		Title:    blog.GetTitle(),
	}

	res, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("InternalError: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OId"),
		)

	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil

}

func (*server) FileUpload(stream blogpb.BlogService_FileUploadServer) error {
	fmt.Println("FileUpload Request")
	var data []byte
	i := true
	var fileName string
	for {
		req, err := stream.Recv()
		if i {
			fmt.Printf("\n%s", req.GetFileName())
			fileName = req.GetFileName()
		}

		if err == io.EOF {
			fmt.Printf("\n%s", data)

			uploadStream, err := bucket.OpenUploadStream(
				fileName,
			)
			if err != nil {
				log.Fatalf("Error while opening file streaming : %v", err)
			}

			defer uploadStream.Close()

			fileSize, err := uploadStream.Write(data)
			if err != nil {
				log.Fatalf("Error while writing file : %v", err)

			}
			return stream.SendAndClose(&blogpb.FileUploadResponse{
				Message: fmt.Sprintf("Write file to DB was successful. File size: %d M\n", fileSize),
			})
		}
		newline := "\n"
		data = append(data, req.Content...)
		data = append(data, newline...)

		if err != nil {
			log.Fatalf("Error while reading file streaming : %v", err)
		}
		i = false
	}
}

func (*server) FileDownload(ctx context.Context, req *blogpb.FileDownloadRequest) (*blogpb.FileDownloadResponse, error) {

	fmt.Println("FileDownload Request")

	var results bson.M
	err := fsFiles.FindOne(context.Background(), bson.M{}).Decode(&results)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find file: %v", err),
		)
	}
	// you can print out the results
	fmt.Println(results)

	fileName := req.FileName

	var buf bytes.Buffer
	_, err = bucket.DownloadToStreamByName(fileName, &buf)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot download file from DB: %v", err),
		)
	}

	return &blogpb.FileDownloadResponse{
		Content: []byte(buf.Bytes()),
	}, nil

}

func main() {
	//if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog service Started")
	fmt.Println("Connecting to MongoDB")
	//connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("mydb").Collection("blog")

	//Creating bucket for file upload
	db := client.Database("myfiles")
	fsFiles = db.Collection("fs.files")
	bucket, err = gridfs.NewBucket(
		db,
	)
	if err != nil {
		log.Fatalf("Failed to Create Bucket: %v", err)
	}
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to Listen %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server....")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}()

	//Wait for Control C to exit
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	//Block until a signal is recieved
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing the mongodb connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of the program.")

}
