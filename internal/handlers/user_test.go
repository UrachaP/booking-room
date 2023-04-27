package handlers

//func Test_GetUser(t *testing.T) {
//	expected := http.StatusOK
//	//requestId := 2
//	//jsonRequest, _ := json.Marshal(requestId)
//	//request := httptest.NewRequest("GET", "/user/:id", bytes.NewBuffer(jsonRequest))
//	request := httptest.NewRequest("GET", "/user/2", nil)
//	writer := httptest.NewRecorder()
//
//	mockRepository := new(mock.MockRepository)
//	mockRepository.On("GetUserById", 2).
//		Return(models.Users{
//			Model: models.Model{
//				ID:        2,
//				UpdatedAt: time.Date(2022, 11, 9, 14, 2, 26, 0, time.Local),
//			},
//			FirstName: "wasawat",
//			LastName:  "bungkanjana",
//			SumGrade:  "C",
//			AGrade:    "C",
//			BGrade:    "D",
//			CGrade:    "B",
//			ImagePath: "/Users/uracha.p/GolandProjects/booking-rooms/assets/image/1600px-Pizigani_1367_Chart_1MB.jpeg",
//			ImageID:   10,
//		}, nil)
//
//	handler := controllers.Handlers{
//		Service: &services.Service{
//			Repository: mockRepository,
//		},
//		}
//	}
//
//	testRoute := echo.New()
//	testRoute.GET("/user/:id", handler.GetUser)
//	testRoute.ServeHTTP(writer, request)
//
//	response := writer.Result()
//
//	assert.Equal(t, expectedStatus, response.StatusCode)
//}
