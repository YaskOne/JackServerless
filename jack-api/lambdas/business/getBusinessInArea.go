package main

/*
	 Fetch businesses in area
*/

//type fetchBusinessInAreaRequest struct {
//	NearLeftLatitude float64 `json:"near_left_latitude"`
//	NearLeftLongitude float64 `json:"near_left_longitude"`
//	FarRightLatitude float64 `json:"far_right_latitude"`
//	FarRightLongitude float64 `json:"far_right_longitude"`
//}
//
//func fetchBusinessInArea(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
//	businesses := []db.Business {}
//
//	nearLeftLat, nearLeftLatErr := strconv.ParseFloat(request.QueryStringParameters["near_left_latitude"], 64)
//	nearLeftLng, nearLeftLngErr := strconv.ParseFloat(request.QueryStringParameters["near_left_longitude"], 64)
//	farRightLat, farRightLatErr := strconv.ParseFloat(request.QueryStringParameters["far_right_latitude"], 64)
//	farRightLng, farRightLngErr := strconv.ParseFloat(request.QueryStringParameters["far_right_longitude"], 64)
//
//	if nearLeftLatErr != nil || nearLeftLngErr != nil || farRightLatErr != nil || farRightLngErr != nil {
//		return core.MakeHTTPError(400, "Error in request parameters")
//	}
//
//	println("YAYAYAYA")
//
//	db.DB().Where(
//			"latitude >= ? AND latitude <= ? AND longitude >= ? AND longitude <= ?",
//			nearLeftLat,
//			farRightLat,
//			nearLeftLng,
//			farRightLng,
//		).Find(&businesses)
//
//	return core.MakeHTTPResponse(200, db.BusinessesResponse{businesses})
//}
//
//func main() {
//	lambda.Start(fetchBusinessInArea)
//}
