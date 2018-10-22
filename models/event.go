package models

const (
	eventRegistryName = "/events/"
)

// // CreateIFEvent creates an IF event from a trigger
// func CreateIFEvent(source, message string, timestamp int64) (*Event, error) {
// 	return createEvent(IF, source, message, timestamp)
// }

// // CreateTHENEvent creates an THEN event from a trigger
// func CreateTHENEvent(source, message string, timestamp int64) (*Event, error) {
// 	return createEvent(THEN, source, message, timestamp)
// }

// // CreateFINALLYEvent creates an FINALLY event from a trigger
// func CreateFINALLYEvent(source, message string, timestamp int64) (*Event, error) {
// 	return createEvent(FINALLY, source, message, timestamp)
// }

// func createEvent(evtType EventType, source, message string, timestamp int64) (*Event, error) {
// 	event := Event{Type: evtType, Source: source, Message: message, Timestamp: timestamp}

// 	client, _, err := storage.CreateClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	//defer destroyFunc()
// 	// timeOutContext, cancel := context.WithTimeout(
// 	// 	context.Background(), 5*time.Second)
// 	// defer cancel()

// 	log.Println(client)

// 	data, err := json.Marshal(event)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = client.CreateObject(
// 		context.Background(),
// 		fmt.Sprintf("%s%s/%s/%s", eventRegistryName, event.Type, event.Source, strconv.FormatInt(event.Timestamp, 10)),
// 		string(data),
// 		30)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &event, nil
// }

// func WatchEvent() <-chan string {
// 	client, destroyFunc, err := storage.CreateClient()
// 	if err != nil {
// 		return nil
// 	}
// 	defer destroyFunc()
// 	timeOutContext, cancel := context.WithTimeout(
// 		context.Background(), 5*time.Second)
// 	defer cancel()

// 	out := make(chan string)
// 	go func() {
// 		in := client.Watch(timeOutContext, eventRegistryName)
// 		for msg := range in {
// 			out <- msg
// 		}
// 		close(out)
// 	}()

// 	return out
// }
