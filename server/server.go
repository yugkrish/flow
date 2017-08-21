package server

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	protoactor "github.com/AsynkronIT/protoactor-go/actor"
	"github.com/fnproject/completer/actor"
	"github.com/fnproject/completer/model"
	"github.com/fnproject/completer/query"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/url"
	"github.com/fnproject/completer/setup"
)

const (
	headerDatumType    string = "FnProject-DatumType"
	headerResultStatus string = "FnProject-ResultStatus"
	headerErrorType    string = "FnProject-ErrorType"
	headerThreadID     string = "FnProject-ThreadID"
	headerStageID      string = "FnProject-StageID"
	headerHeaderPrefix string = "FnProject-Header-"
	headerMethod       string = "FnProject-Method"
	headerResultCode   string = "FnProject-ResultCode"
)

var log = logrus.WithField("logger", "server")

func (s *Server) completeExternally(graphID string, stageID string, body []byte, headers http.Header, method string, contentType string, b bool) (*model.CompleteStageExternallyResponse, error) {
	var hs []*model.HttpHeader
	for k, vs := range headers {
		for _, v := range vs {
			hs = append(hs, &model.HttpHeader{
				Key:   k,
				Value: v,
			})
		}
	}

	var m model.HttpMethod
	if methodValue, found := model.HttpMethod_value[method]; found {
		m = model.HttpMethod(methodValue)
	} else {
		m = model.HttpMethod_unknown_method
	}

	if contentType == "" {
		contentType = "application/octet-stream"
	}
	httpReqDatum := model.HttpReqDatum{
		Body:    model.NewBlob(contentType, body),
		Headers: hs,
		Method:  m,
	}

	request := model.CompleteStageExternallyRequest{
		GraphId: graphID,
		StageId: stageID,
		Result: &model.CompletionResult{
			Successful: b,
			Datum:      model.NewHttpReqDatum(&httpReqDatum),
		},
	}

	response, err := s.graphManager.CompleteStageExternally(&request, 5*time.Second)

	return response, err
}

func (s *Server) handleStageOperation(c *gin.Context) {
	graphID := c.Param("graphId")
	stageID := c.Param("stageId")
	operation := c.Param("operation")
	body, err := c.GetRawData()
	if err != nil {
		message := "handleStageOperation can't get raw data from the request"
		log.WithError(err).Error(message)
		c.Status(http.StatusInternalServerError)
		return
	}

	// TODO: generic error handling (e.g. graph not found)
	switch operation {
	case "complete":
		response, err := s.completeExternally(graphID, stageID, body, c.Request.Header, c.Request.Method, c.ContentType(), true)
		if err != nil {
			message := "completeExternally returned an error trying to complete with success"
			log.WithError(err).Error(message)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Header(headerStageID, response.StageId)
		c.Status(http.StatusOK)
	case "fail":
		response, err := s.completeExternally(graphID, stageID, body, c.Request.Header, c.Request.Method, c.ContentType(), false)
		if err != nil {
			message := "completeExternally returned an error trying to complete with failure"
			log.WithError(err).Error(message)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Header(headerStageID, response.StageId)
		c.Status(http.StatusOK)
	default:
		other := c.Query("other")
		cids := []string{stageID}
		if other != "" {
			cids = append(cids, other)
		}

		completionOperation, found := model.CompletionOperation_value[operation]
		if !found {
			message := "invalid completion operation '" + operation + "'"
			c.Data(http.StatusBadRequest, "text/plain", []byte(message))
			return
		}

		// TODO: enforce valid content type
		// TODO: generic error handling
		request := &model.AddChainedStageRequest{
			GraphId:   graphID,
			Deps:      cids,
			Operation: model.CompletionOperation(completionOperation),
			Closure:   model.NewBlob(c.ContentType(), body),
		}
		response, err := s.addStage(request)

		if err != nil {
			renderError(err, c, "Graph Operation Failed ")
			return
		}

		c.Header(headerStageID, response.StageId)
		c.Status(http.StatusOK)
	}
}
func renderError(err error, c *gin.Context, msg string) {
	log.WithError(err).Error(msg)
	switch e := err.(type) {
	case model.BadRequestError:
		c.Data(400, "text/plain", []byte(e.UserMessage()))
	default:
		c.Status(http.StatusInternalServerError)
	}
}

func (s *Server) handleCreateGraph(c *gin.Context) {
	log.Info("Creating graph")
	functionID := c.Query("functionId")

	if functionID == "" {
		c.Data(http.StatusBadRequest, "text/plain", []byte("Function ID must be provided"))
		return
	}
	// TODO: Validate the format of the functionId

	graphID, err := uuid.NewRandom()
	if err != nil {
		message := "can't get a UUID for the new graph"
		log.WithError(err).Error(message)
		c.Status(http.StatusInternalServerError)
		return
	}

	req := &model.CreateGraphRequest{FunctionId: functionID, GraphId: graphID.String()}

	// TODO: sort out timeouts in a consistent way
	result, err := s.graphManager.CreateGraph(req, 5*time.Second)
	if err != nil {
		renderError(err, c, "failed to create new graph")
		return
	}
	c.Header(headerThreadID, result.GraphId)
	c.Status(http.StatusOK)
}

func (s *Server) handleGraphState(c *gin.Context) {

	graphID := c.Param("graphId")
	log.Info("Requested graph with Id " + graphID)

	request := &model.GetGraphStateRequest{GraphId: graphID}

	resp, err := s.graphManager.GetGraphState(request, 5*time.Second)

	if err != nil {
		renderError(err, c, "failed to read graph state")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func resultStatus(result *model.CompletionResult) string {
	if result.GetSuccessful() {
		return "success"
	}
	return "failure"
}

func (s *Server) handleGetGraphStage(c *gin.Context) {
	graphID := c.Param("graphId")
	stageID := c.Param("stageId")

	request := model.GetStageResultRequest{
		GraphId: graphID,
		StageId: stageID,
	}

	// TODO: return 408 on timeout
	response, err := s.graphManager.GetStageResult(&request, 5*time.Second)

	if err == protoactor.ErrTimeout {
		c.Data(http.StatusRequestTimeout, "text/plain", []byte("stage not completed"))
		return
	}

	if err != nil {
		renderError(err, c, "Error getting stage result")
		return
	}

	result := response.GetResult()
	datum := result.GetDatum()
	val := datum.GetVal()
	if val == nil {
		message := "GetStageResult produced a result but the datum value is null"
		log.Error(message)
		c.Data(http.StatusInternalServerError, "text/plain", []byte(message))
		return
	}

	switch v := val.(type) {

	// TODO: refactor this by adding a writer to a context in multipart_write.go
	case *model.Datum_Error:
		c.Header(headerDatumType, "error")
		c.Header(headerResultStatus, resultStatus(result))
		err := v.Error
		c.Header(headerErrorType, model.ErrorDatumType_name[int32(err.GetType())])
		c.String(http.StatusOK, err.GetMessage())
		return
	case *model.Datum_Empty:
		c.Header(headerDatumType, "empty")
		c.Header(headerResultStatus, resultStatus(result))
		c.Status(http.StatusOK)
		return
	case *model.Datum_Blob:
		c.Header(headerDatumType, "blob")
		c.Header(headerResultStatus, resultStatus(result))
		blob := v.Blob
		c.Data(http.StatusOK, blob.GetContentType(), blob.GetDataString())
		return
	case *model.Datum_StageRef:
		c.Header(headerDatumType, "stageref")
		c.Header(headerResultStatus, resultStatus(result))
		stageRef := v.StageRef
		c.Header(headerStageID, stageRef.StageRef)
		c.Status(http.StatusOK)
		return
	case *model.Datum_HttpReq:
		c.Header(headerDatumType, "httpreq")
		c.Header(headerResultStatus, resultStatus(result))
		httpReq := v.HttpReq
		for _, header := range httpReq.Headers {
			c.Header(headerHeaderPrefix+header.GetKey(), header.GetValue())
		}
		httpMethod := model.HttpMethod_name[int32(httpReq.GetMethod())]
		c.Header(headerMethod, httpMethod)
		c.Data(http.StatusOK, httpReq.Body.GetContentType(), httpReq.Body.GetDataString())
		return
	case *model.Datum_HttpResp:
		c.Header(headerDatumType, "httpresp")
		c.Header(headerResultStatus, resultStatus(result))
		httpResp := v.HttpResp
		for _, header := range httpResp.Headers {
			c.Header(headerHeaderPrefix+header.GetKey(), header.GetValue())
		}
		statusCode := strconv.FormatUint(uint64(httpResp.GetStatusCode()), 32)
		c.Header(headerResultCode, statusCode)
		c.Data(http.StatusOK, httpResp.Body.GetContentType(), httpResp.Body.GetDataString())
		return
	default:
		log.Error("unrecognized datum type when getting graph stage")
		c.Status(http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleExternalCompletion(c *gin.Context) {
	graphID := c.Param("graphId")

	request := &model.AddExternalCompletionStageRequest{GraphId: graphID}

	response, err := s.addStage(request)

	if err != nil {
		renderError(err, c, "Failed to add stage")
		return
	}
	c.Header(headerStageID, response.StageId)
	c.Status(http.StatusOK)
}

func (s *Server) allOrAnyOf(c *gin.Context, op model.CompletionOperation) {
	cidList := c.Query("cids")
	graphID := c.Param("graphId")
	cids := strings.Split(cidList, ",")

	request := &model.AddChainedStageRequest{
		GraphId:   graphID,
		Operation: op,
		Closure:   nil,
		Deps:      cids,
	}

	response, err := s.addStage(request)

	// TODO: Actually some errors should be user errors here (e.g. AnyOf with zero dependencies)
	if err != nil {
		renderError(err, c, "Failed to add stage")
		return
	}
	c.Header(headerStageID, response.StageId)
	c.Status(http.StatusOK)
}

func (s *Server) handleAllOf(c *gin.Context) {
	s.allOrAnyOf(c, model.CompletionOperation_allOf)
}

func (s *Server) handleAnyOf(c *gin.Context) {
	s.allOrAnyOf(c, model.CompletionOperation_anyOf)
}

func (s *Server) handleSupply(c *gin.Context) {
	graphID := c.Param("graphId")
	// TODO: Validate graphId: not empty

	body, err := c.GetRawData()
	if err != nil {
		renderError(err, c, "Failed to read body")
		return
	}

	request := &model.AddChainedStageRequest{
		GraphId:   graphID,
		Operation: model.CompletionOperation_supply,
		Closure:   model.NewBlob(c.ContentType(), body),
		Deps:      []string{},
	}

	// TODO: handle user errors properly (E.g. graph not found)
	response, err := s.addStage(request)
	if err != nil {
		renderError(err, c, "Failed to add stage")
		return
	}
	c.Header(headerStageID, response.StageId)
	c.Status(http.StatusOK)
}

func (s *Server) handleCompletedValue(c *gin.Context) {
	graphID := c.Param("graphId")
	// TODO: again, validate graphId

	body, err := c.GetRawData()
	if err != nil {
		message := "completedValue cannot get raw request data"
		log.WithError(err).Error(message)
		c.Status(http.StatusInternalServerError)
		return
	}

	var contentType = c.ContentType()
	if contentType == "" {
		// TODO: error with a bad request
	}
	bodyBlob := model.NewBlob(contentType, body)

	result := model.CompletionResult{
		Successful: true,
		Datum:      model.NewBlobDatum(bodyBlob),
	}

	// TODO: generic error handling
	request := &model.AddCompletedValueStageRequest{
		GraphId: graphID,
		Result:  &result,
	}

	response, err := s.addStage(request)
	if err != nil {
		renderError(err, c, "Failed to add stage")
		return
	}
	c.Header(headerStageID, response.StageId)
	c.Status(http.StatusOK)
}

func (s *Server) addStage(request model.AddStageCommand) (*model.AddStageResponse, error) {
	// TODO: we should have a synchronous api in the graph manager instead of this
	return s.graphManager.AddStage(request, 5*time.Second)
}

func (s *Server) handleCommit(c *gin.Context) {
	graphID := c.Param("graphId")
	request := model.CommitGraphRequest{GraphId: graphID}

	response, err := s.graphManager.Commit(&request, 5*time.Second)
	if err != nil {
		message := "handleCommit failed to commit"
		log.WithError(err).Error(message)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header(headerThreadID, response.GraphId)
	c.Status(http.StatusOK)
}

func (s *Server) handleDelay(c *gin.Context) {
	graphID := c.Param("graphId")
	delayMs := c.Query("delayMs")
	if delayMs == "" {
		log.Info("Empty or missing delay value supplied to add delay stage")
		c.Status(http.StatusBadRequest)
		return
	}

	delay, err := strconv.ParseInt(delayMs, 10, 64)
	if err != nil {
		log.Info("Invalid delay value supplied to add delay stage")
		c.Status(http.StatusBadRequest)
		return
	}

	// TODO: generic error handling (e.g. graph not found)
	request := &model.AddDelayStageRequest{GraphId: graphID, DelayMs: delay}

	response, err := s.addStage(request)

	if err != nil {
		renderError(err, c, "Failed to add stage")
		return
	}
	c.Header(headerStageID, response.StageId)
	c.Status(http.StatusOK)
}

func unwrapPrefixedHeaders(hs http.Header) []*model.HttpHeader {
	var headers []*model.HttpHeader
	for k, vs := range hs {
		canonicalKey := http.CanonicalHeaderKey(k)
		canonicalPrefix := http.CanonicalHeaderKey(headerHeaderPrefix)
		if strings.HasPrefix(canonicalKey, canonicalPrefix) {
			trimmedHeader := strings.TrimPrefix(canonicalKey, canonicalPrefix)
			for _, v := range vs {
				headers = append(headers, &model.HttpHeader{
					Key:   trimmedHeader,
					Value: v,
				})
			}
		}
	}
	return headers
}

func (s *Server) handleInvokeFunction(c *gin.Context) {
	graphID := c.Param("graphId")
	// TODO: validate graphId not empty

	functionID := c.Query("functionId")
	if functionID == "" {
		message := "Empty or missing functionId supplied to add invokeFunction stage"
		log.Info(message)
		c.Data(http.StatusBadRequest, "text/plain", []byte(message))
		return
	}
	// TODO: Validate the format of the functionId

	body, err := c.GetRawData()
	if err != nil {
		message := "Invalid request body supplied to add invokeFunction stage"
		log.Info(message)
		c.Data(http.StatusBadRequest, "text/plain", []byte(message))
		return
	}

	requestMethod := strings.ToLower(c.GetHeader(http.CanonicalHeaderKey(headerMethod)))
	var method model.HttpMethod
	if m, found := model.HttpMethod_value[requestMethod]; found {
		method = model.HttpMethod(m)
	} else {
		message := "Empty, missing, or invalid HTTP method supplied to add invokeFunction stage"
		log.Info(message)
		c.Data(http.StatusBadRequest, "text/plain", []byte(message))
		return
	}

	var contentType = c.ContentType()
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	var bodyBlob *model.BlobDatum = nil
	if len(body) != 0 {
		bodyBlob = model.NewBlob(contentType, body)
	}

	// TODO: test this with an empty request
	request := &model.AddInvokeFunctionStageRequest{
		GraphId:    graphID,
		FunctionId: functionID,
		Arg: &model.HttpReqDatum{
			Body:    bodyBlob,
			Headers: unwrapPrefixedHeaders(c.Request.Header),
			Method:  method,
		},
	}

	// TODO: generic error handling (e.g. graph does not exist)
	response, err := s.addStage(request)
	if err != nil {
		renderError(err, c, "Failed to add stage")
		return
	}
	c.Header(headerStageID, response.StageId)
	c.Status(http.StatusOK)
}

type Server struct {
	engine       *gin.Engine
	graphManager actor.GraphManager
	apiUrl       *url.URL
	listen       string
}

func NewFromEnv(manager actor.GraphManager) (*Server, error) {

	s := &Server{
		graphManager: manager,
		engine:       gin.Default(),
		listen:       setup.GetString(setup.EnvListen),
	}

	s.engine.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	s.engine.GET("/wss", func(c *gin.Context) {
		query.WSSHandler(manager, c.Writer, c.Request)
	})

	graph := s.engine.Group("/graph")
	{
		graph.POST("", s.handleCreateGraph)
		graph.GET("/:graphId", s.handleGraphState)

		graph.POST("/:graphId/supply", s.handleSupply)
		graph.POST("/:graphId/invokeFunction", s.handleInvokeFunction)
		graph.POST("/:graphId/completedValue", s.handleCompletedValue)
		graph.POST("/:graphId/delay", s.handleDelay)
		graph.POST("/:graphId/allOf", s.handleAllOf)
		graph.POST("/:graphId/anyOf", s.handleAnyOf)
		graph.POST("/:graphId/externalCompletion", s.handleExternalCompletion)
		graph.POST("/:graphId/commit", s.handleCommit)

		stage := graph.Group("/:graphId/stage")
		{
			stage.GET("/:stageId", s.handleGetGraphStage)
			stage.POST("/:stageId/:operation", s.handleStageOperation)
		}
	}

	return s, nil
}

func (s *Server) Run() {
	log.WithField("listen_url", s.listen).Info("Starting Completer server")

	s.engine.Run(s.listen)
}