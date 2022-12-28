package main

import (
	"log"
	"net"
	"os"
	"strconv"

	auction "github.com/Mlth/AuctionHouse/proto"
	"google.golang.org/grpc"
)

type AuctionHouseServer struct {
	auction.AuctionHouseServer
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000
	ownPortStr := strconv.Itoa(int(ownPort))

	//Listening on own port
	list, err := net.Listen("tpc", ":"+ownPortStr)
	if err != nil {
		log.Fatalf("Failed to listen on port")
	}
	grpcServer := grpc.NewServer()
	auction.RegisterAuctionHouseServer(grpcServer, &AuctionHouseServer{})
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}

func (s *AuctionHouseServer) ReceiveBid(*auction.BidMessage) (*auction.AckMessage, error) {

}
