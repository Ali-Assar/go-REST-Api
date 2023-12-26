package routes

func getEvents(context *gin.Context) {
	events,err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events. try again later"})
	return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context){
	eventId,err:= strconv.ParseInt(context.Param("id"),10,64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id"})
	return
	}

	event,err:= models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events. try again later"})
	return
	}
	context.JSON(http.StatusOK, event)

}


func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not create events. try again later"})
	return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
