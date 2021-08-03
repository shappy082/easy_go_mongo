// GET ALL

filter := bson.M{}
findOptions := options.Find()

returnValue := model.ResponseMessageStatusData{}
returnValue.Data = []string{}

result, err := repository.CollectionTodoList.Find(filter, findOptions)
log.Println(result, err)
if err != nil {
	returnValue.Status = "error"
	return returnValue
}

returnValue.Status = "success"
returnValue.Data = result
return returnValue

// ADDED

returnValue := model.ResponseMessageStatusData{}
returnValue.Data = []string{}

data := map[string]interface{}{
	"id":          "1",
	"subject":     "First Task",
	"description": "Description of first task",
	"time":        "10:00",
	"create_date": "2020-12-31 23:59:59",
	"update_date": "2020-12-31 23:59:59",
}
err := repository.CollectionTodoList.Add(data)
log.Println(err)
if err != nil {
	returnValue.Status = "error"
	return returnValue
}

returnValue.Status = "success"
return returnValue

// EDITED

returnValue := model.ResponseMessageStatusData{}
returnValue.Data = []string{}

filter := bson.M{"id": "2"}
data := map[string]interface{}{
	"id":          "2",
	"subject":     "2 Edited Task",
	"description": "Description of 2 task",
	"time":        "11:00",
	"create_date": "2020-12-31 23:59:59",
	"update_date": "2020-12-31 23:59:59",
}
err := repository.CollectionTodoList.Update(filter, data)
log.Println(err)
if err != nil {
	returnValue.Status = "error"
	return returnValue
}

returnValue.Status = "success"
return returnValue

// DELETE

returnValue := model.ResponseMessageStatusData{}
returnValue.Data = []string{}

filter := bson.M{"id": "2"}
err := repository.CollectionTodoList.Delete(filter)
log.Println(err)
if err != nil {
	returnValue.Status = "error"
	return returnValue
}

returnValue.Status = "success"
return returnValue