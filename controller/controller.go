package controller

//func StartServer() {
//
//	//data.ConnectionToElasticSearchServer()
//
//	l := log.New(os.Stdout, "Rules-API", log.LstdFlags)
//
//	rh := handlers.NewRules(l)
//	//rc := handlers.NewRuleChecker(l)
//
//	sm := mux.NewRouter()
//	getRouter := sm.Methods(http.MethodGet).Subrouter()
//	getRouter.HandleFunc("/", rh.GetRules)
//
//	postRouter := sm.Methods(http.MethodPost).Subrouter()
//	postRouter.HandleFunc("/add", rh.AddRule)
//	//postRouter.HandleFunc("/cacheRule", rc.GetRuleCheck)
//
//	putRouter := sm.Methods(http.MethodPut).Subrouter()
//	putRouter.HandleFunc("/{id:[0-9]+}", rh.UpdateRule)
//
//	s := &http.Server{
//		Addr:         ":8088",
//		Handler:      sm,
//		ReadTimeout:  10 * time.Second,
//		WriteTimeout: 10 * time.Second,
//		IdleTimeout:  120 * time.Second,
//	}
//	s.ListenAndServe()
//
//}

//func rules(rw http.ResponseWriter, r *http.Request) {
//	lr := data.GetRules()
//	byte,err := json.Marshal(&lr)
//
//	if err !=nil {
//		http.Error(rw, "Unable to Marshal", http.StatusInternalServerError)
//	}
//	rw.Write(byte)
//}

//
//func handler(rw http.ResponseWriter, r *http.Request) {
//
//}
//
//func checksFact(rw http.ResponseWriter, r *http.Request) {
//
//	body,err := ioutil.ReadAll(r.Body)
//	if err != nil{
//	fmt.Println(body)
//	}
