namespace cpp didirgeo.service

struct LocationInfoRequest {
    1: required string user_name;
    2: required string token;
    3: required string productid;
    4: required double lng;
    5: required double lat;
    6: optional i32 radius;
    7: optional i32 max_num;
    8: optional string traceid;
}

enum ResultType {
    FENCE = 1;
    LINES = 2;
}

struct Result {
    1: required string name;  //名称，可以商圈或者街道名称
    2: required string id;    //商圈或者街道的唯一标识
    3: required ResultType kind;  //表示是“街道” 或者“商圈”
    4: optional double dist;  //查询点到街道的最近距离，单位是米(对商圈，该字段无意义)
}

struct LocationInfoResponse {
    1: required i32 status;
    2: required string message;
    3: required i32 time;
    4: required i32 area;
    5: required string country;
    6: required string provice;
    7: required string city;
    8: required string district;
    9: required list<Result> result;
}

service DidiRgeoService {
    LocationInfoResponse location_info(1:LocationInfoRequest request);
}
