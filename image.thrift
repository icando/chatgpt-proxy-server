namespace go api

struct ImageRequest {
	string      Prompt
	i32         N
	string      Size
	string      ResponseFormat
	string      User
}

struct ImageResponse {
	i64                             Created
	list<ImageResponseDataInner>    Data
}

struct ImageResponseDataInner {
    string URL
    string B64JSON
}

service GPTService {
    ImageResponse CreateImage(ImageRequest request)
}