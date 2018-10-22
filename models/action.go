package models

// // CreateAction creates a db entry for an action object
// func CreateAction(action Action) error {
// 	client, destroyFunc, err := storage.CreateClient()
// 	if err != nil {
// 		return err
// 	}
// 	defer destroyFunc()
// 	timeOutContext, cancel := context.WithTimeout(
// 		context.Background(), 5*time.Second)
// 	defer cancel()

// 	data, err := action.MarshalJSON()
// 	if err != nil {
// 		return err
// 	}

// 	path := fmt.Sprintf("%s%s/%s", actionRegistryName, action.Module, action.Name)
// 	err = client.CreateObject(timeOutContext, path, string(data), 0)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // GetAction retrieves a db entry for an action object
// func GetAction(module, name string) (*Action, error) {
// 	client, destroyFunc, err := storage.CreateClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer destroyFunc()
// 	timeOutContext, cancel := context.WithTimeout(
// 		context.Background(), 5*time.Second)
// 	defer cancel()

// 	path := fmt.Sprintf("%s%s/%s", actionRegistryName, module, name)
// 	data, err := client.GetObject(timeOutContext, path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var a Action
// 	err = a.UnmarshalJSON(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, err
// }

// // GetActions retrieves all actions
// func GetActions() ([]*Action, error) {
// 	client, destroyFunc, err := storage.CreateClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer destroyFunc()
// 	timeOutContext, cancel := context.WithTimeout(
// 		context.Background(), 5*time.Second)
// 	defer cancel()

// 	data, err := client.GetObjects(timeOutContext, actionRegistryName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var actions []*Action

// 	for _, d := range data {
// 		var a Action
// 		err = json.Unmarshal(d, &a)
// 		if err != nil {
// 			log.Println(err)
// 			return nil, err
// 		}
// 		actions = append(actions, &a)
// 	}
// 	return actions, err
// }

// // DeleteAction deletes a db entry for an action object
// func DeleteAction(module, name string) error {
// 	client, destroyFunc, err := storage.CreateClient()
// 	if err != nil {
// 		return err
// 	}
// 	defer destroyFunc()
// 	timeOutContext, cancel := context.WithTimeout(
// 		context.Background(), 5*time.Second)
// 	defer cancel()

// 	path := fmt.Sprintf("%s%s/%s", actionRegistryName, module, name)
// 	count, err := client.DeleteObject(timeOutContext, path)
// 	if err != nil {
// 		return err
// 	}
// 	if count == 0 {
// 		return fmt.Errorf("key %s not exist, nothing deleted", path)
// 	}
// 	return nil
// }

// // UpdateAction updates a db entry for an action object
// func UpdateAction(action *Action) error {
// 	client, destroyFunc, err := storage.CreateClient()
// 	if err != nil {
// 		return err
// 	}
// 	defer destroyFunc()
// 	timeOutContext, cancel := context.WithTimeout(
// 		context.Background(), 5*time.Second)
// 	defer cancel()

// 	path := fmt.Sprintf("%s%s/%s", actionRegistryName, action.Module, action.Name)
// 	data, err := json.Marshal(action)
// 	if err != nil {
// 		return err
// 	}
// 	return client.UpdateObject(timeOutContext, path, string(data))
// }
