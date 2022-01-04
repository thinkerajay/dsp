package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerajay/dsp/pb"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"os"
)

type adCandidate struct{
	Id string `json:"id"`
	LineItemID string `json:"line_item_id"`
}

type incomingRequest struct {
	AuctionType string `json:"auction-type"`
	Domain      string `json:"domain"`
	Referrer    string `json:"referrer"`
	AdCandidates []*adCandidate `json:"ad-candidates"`
}

func MainHandler(ctx *gin.Context) {

	log.Println("received request", ctx.Request)
	protoRequest := pb.IncomingRequest{}

	reqBody := ctx.Request.Body
	if reqBody == nil {
		ctx.JSON(400, "bad request !")
		return
	}
	defer func(reqBody io.ReadCloser) {
		err := reqBody.Close()
		if err != nil {

		}
	}(reqBody)

	incomingReq := new(incomingRequest)
	reqBytes, err := io.ReadAll(reqBody)
	if err != nil {
		ctx.JSON(400, "bad request !")
		return
	}

	err = json.Unmarshal(reqBytes, incomingReq)
	if err != nil {
		ctx.JSON(400, "bad request !")
		return
	}
	protoRequest.AuctionType = incomingReq.AuctionType
	protoRequest.Referrer = incomingReq.Referrer
	protoRequest.Domain = incomingReq.Domain
	protoRequest.Adcandidates = make([]*pb.AdCandidate, 0 )

	for _, adCandidate := range incomingReq.AdCandidates{
		protoRequest.Adcandidates = append(protoRequest.Adcandidates, &pb.AdCandidate{Id: adCandidate.Id, LineItemId: adCandidate.LineItemID})
	}

	//TODO: remove this code
	pfile, err := os.Create("requestFile")
	if err != nil {
		ctx.JSON(400, "bad request !")
		return
	}
	wb, _ := proto.Marshal(&protoRequest)

	_, err = pfile.Write(wb)
	if err != nil {
		return
	}

	ctx.JSON(200, "Success !")
}
