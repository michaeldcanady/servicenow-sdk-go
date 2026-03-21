//go:build integration

package tests

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/gorilla/mux"
)

const (
	tableRecordPath     = "/api/now/v1/table/{tableName}/{sys_id}"
	tableCollectionPath = "/api/now/v1/table/{tableName}"
)

func getRecord(w http.ResponseWriter, r *http.Request, table, sysID string) {
	store.mu.RLock()
	t, ok := store.Tables[table]
	store.mu.RUnlock()

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	t.mu.RLock()
	rec, ok := t.records[sysID]
	t.mu.RUnlock()

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	writeResult(w, http.StatusOK, rec)
}

func updateRecord(w http.ResponseWriter, r *http.Request, table, sysID string) {
	store.mu.RLock()
	t, ok := store.Tables[table]
	store.mu.RUnlock()

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	rec, ok := t.records[sysID]
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	var patch map[string]any
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	for k, v := range patch {
		rec[k] = v
	}

	writeResult(w, http.StatusOK, rec)
}

func deleteRecord(w http.ResponseWriter, r *http.Request, table, sysID string) {
	store.mu.RLock()
	t, ok := store.Tables[table]
	store.mu.RUnlock()

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	if _, exists := t.records[sysID]; !exists {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	delete(t.records, sysID)

	// ServiceNow returns 204 No Content with {"result": null}
	writeResult(w, http.StatusNoContent, nil)
}

func handleTableRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table := vars["tableName"]
	sysID := vars["sys_id"]

	switch r.Method {
	case http.MethodGet:
		getRecord(w, r, table, sysID)
	case http.MethodPatch:
		updateRecord(w, r, table, sysID)
	case http.MethodDelete:
		deleteRecord(w, r, table, sysID)
	}
}

func writeResult(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"result": data,
	})
}

type TableStore struct {
	mu      sync.RWMutex
	records map[string]map[string]any // sys_id → record
}

func NewTableStore() *TableStore {
	return &TableStore{
		records: make(map[string]map[string]any),
	}
}

type Storage struct {
	Tables      map[string]*TableStore
	Attachments map[string][]byte
	mu          sync.RWMutex
}

var store = &Storage{
	Tables:      make(map[string]*TableStore),
	Attachments: make(map[string][]byte),
}

func newSysID() string {
	const hex = "0123456789abcdef"
	b := make([]byte, 32)
	for i := range b {
		b[i] = hex[rand.Intn(len(hex))]
	}
	return string(b)
}

func createRecord(w http.ResponseWriter, r *http.Request, table string) {
	var rec map[string]any
	json.NewDecoder(r.Body).Decode(&rec)

	store.mu.Lock()
	defer store.mu.Unlock()

	t, ok := store.Tables[table]
	if !ok {
		t = NewTableStore()
		store.Tables[table] = t
	}

	sysID := newSysID()
	rec["sys_id"] = sysID

	t.mu.Lock()
	t.records[sysID] = rec
	t.mu.Unlock()

	writeResult(w, http.StatusCreated, rec)
}

func listRecords(w http.ResponseWriter, r *http.Request, table string) {
	store.mu.RLock()
	t, ok := store.Tables[table]
	store.mu.RUnlock()

	if !ok {
		writeResult(w, http.StatusOK, []map[string]any{})
		return
	}

	t.mu.RLock()
	defer t.mu.RUnlock()

	out := make([]map[string]any, 0, len(t.records))
	for _, rec := range t.records {
		out = append(out, rec)
	}

	writeResult(w, http.StatusOK, out)
}

func handleTableCollection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table := vars["tableName"]

	switch r.Method {
	case http.MethodGet:
		listRecords(w, r, table)
	case http.MethodPost:
		createRecord(w, r, table)
	}
}

func newServer() *httptest.Server {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()

	nowRouter := apiRouter.PathPrefix("/now").Subrouter()

	tableRouter := nowRouter.PathPrefix("/v1/table/{tableName}").Subrouter()

	tableRouter.HandleFunc(
		"/{sys_id}",
		handleTableRecord,
	).Methods("GET", "PATCH", "DELETE")

	tableRouter.HandleFunc(
		"/",
		handleTableCollection,
	).Methods("GET", "POST")

	return httptest.NewServer(router)
}
