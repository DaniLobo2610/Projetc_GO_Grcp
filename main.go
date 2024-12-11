package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	pb "Kombat_grpc/proto"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var db *sql.DB

type server struct {
	pb.UnimplementedKombatServiceServer
}

func (s *server) GetKombatInfo(ctx context.Context, req *pb.KombatRequest) (*pb.KombatResponse, error) {

	var name, ptype, Skills string

	query := "select * from Kombat.fighter where Name like  @name"

	row := db.QueryRowContext(ctx, query, sql.Named("name", "%"+req.Name+"%"))
	err := row.Scan(&name, &ptype, &Skills)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.KombatResponse{Name: "Not Found", Type: "Not Found", Skills: "Not Found"}, nil
		}
		return nil, err
	}
	return &pb.KombatResponse{
		Name:   name,
		Type:   ptype,
		Skills: Skills,
	}, nil
}

func (s *server) GetKombatList(req *pb.Empty, stream pb.KombatService_GetKombatListServer) error {
	query := "select * from Kombat.fighter"

	rows, err := db.Query(query)

	if err != nil {
		log.Panic(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name, ptype, Skills string

		if err := rows.Scan(&name, &ptype, &Skills); err != nil {
			log.Panic(err)
			return err
		}
		if err := stream.Send(&pb.KombatResponse{
			Name:   name,
			Type:   ptype,
			Skills: Skills,
		}); err != nil {
			log.Panic(err)
			return err
		}
	}
	return nil
}

func (s *server) AddKombats(stream pb.KombatService_AddKombatsServer) error {
	var count int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AddKombatResponse{Count: count})
		}
		if err != nil {
			log.Printf("Error receiving stream: %v", err)
			return err
		}

		//existe?
		var exists bool
		checkQuery := "SELECT 1 FROM Kombat.fighter WHERE Name = @name"
		err = db.QueryRow(checkQuery, sql.Named("name", req.Name)).Scan(&exists)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("Error checking existing record: %v", err)
			return err
		}

		if exists {
			log.Printf("Record already exists: %s", req.Name)
			return stream.SendAndClose(&pb.AddKombatResponse{
				Count: count,
				Error: fmt.Sprintf(" %s already exists", req.Name),
			})
		}

		query := "INSERT INTO Kombat.fighter (Name, Type, Skills) VALUES (@name, @type, @skills)"
		_, err = db.Exec(query,
			sql.Named("name", req.Name),
			sql.Named("type", req.Type),
			sql.Named("skills", req.Skills))
		if err != nil {
			log.Printf("Error inserting record: %v", err)
			return err
		}

		count++
		log.Printf("Inserted: %s", req.Name)
	}
}

func (s *server) GetKombatByType(stream pb.KombatService_GetKombatByTypeServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of Stream")
			return nil
		}
		if err != nil {
			log.Panic(err)
			return err
		}
		query := "select * from Kombat.fighter where lower(Type) =  lower(@type)"
		rows, err := db.Query(query, sql.Named("type", req.Type))

		if err != nil {
			log.Panic(err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var name, ptype, Skills string

			if err := rows.Scan(&name, &ptype, &Skills); err != nil {
				log.Panic(err)
				return err
			}

			if err := stream.Send(&pb.KombatResponse{
				Name:   name,
				Type:   ptype,
				Skills: Skills,
			}); err != nil {
				log.Panic(err)
				return err
			}

		}
	}

}

func (s *server) GetKombatBySkills(stream pb.KombatService_GetKombatBySkillsServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of Stream")
			return nil
		}
		if err != nil {
			log.Panic(err)
			return err
		}
		query := "select * from Kombat.fighter where lower(Skills) =  lower(@skills)"
		rows, err := db.Query(query, sql.Named("Skills", req.Skills))

		if err != nil {
			log.Panic(err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var name, ptype, Skills string

			if err := rows.Scan(&name, &ptype, &Skills); err != nil {
				log.Panic(err)
				return err
			}

			if err := stream.Send(&pb.KombatResponse{
				Name:   name,
				Type:   ptype,
				Skills: Skills,
			}); err != nil {
				log.Panic(err)
				return err
			}

		}
	}

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, pass, s, port, database)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Conexión establecida")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKombatServiceServer(grpcServer, &server{})

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
		log.Println("Servidor iniciado en el puerto :8081")
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	log.Println("Servidor iniciado en el puerto :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
