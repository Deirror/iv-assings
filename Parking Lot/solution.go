package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Vehicle interface and its implementations

const (
	CAR_TYPE        = "Car"
	MOTORCYCLE_TYPE = "Motorcycle"
	VAN_TYPE        = "Van"
)

type VehicleID interface {
	GetLicensePlate() string
}

type VehicleType interface {
	GetType() string
}

type Vehicle interface {
	VehicleID
	VehicleType
}

type Motorcycle struct {
	LicensePlate
}

func (m Motorcycle) GetType() string {
	return MOTORCYCLE_TYPE
}

type LicensePlate struct {
	licensePlate string
}

func (l LicensePlate) GetLicensePlate() string {
	return l.licensePlate
}

type Car struct {
	LicensePlate
}

func (c Car) GetType() string {
	return CAR_TYPE
}

type Van struct {
	LicensePlate
}

func (v Van) GetType() string {
	return VAN_TYPE
}

// ParkingLot Server struct and its methods

type ParkingLotServer struct {
	mu *sync.RWMutex
	db map[string]Vehicle

	motorcycleCount uint16
	carCount        uint16
	vanCount        uint16

	motorcycleLimit uint16
	compactLimit    uint16
	regularLimit    uint16
}

type VehicleRequest struct {
	LicensePlate string `json:"licensePlate"`
}

func NewParkingLotServer(motorcycleLimit, compactLimit, regularLimit uint16) *ParkingLotServer {
	return &ParkingLotServer{
		mu:              &sync.RWMutex{},
		db:              make(map[string]Vehicle),
		motorcycleLimit: motorcycleLimit,
		compactLimit:    compactLimit,
		regularLimit:    regularLimit,
	}
}

func (p *ParkingLotServer) HandleMotorcycle(w http.ResponseWriter, r *http.Request) {
	var req VehicleRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.db[req.LicensePlate]; ok {
		http.Error(w, "Motorcycle already exists", http.StatusConflict)
		return
	}

	if p.IsMotorcycleFull() {
		http.Error(w, "Parking lot for motorcycle is full", http.StatusConflict)
		return
	}

	if p.IsFull() {
		http.Error(w, "Parking lot is full", http.StatusConflict)
		return
	}

	p.motorcycleCount++
	p.db[req.LicensePlate] = Motorcycle{LicensePlate: LicensePlate{licensePlate: req.LicensePlate}}
}

func (p *ParkingLotServer) HandleCar(w http.ResponseWriter, r *http.Request) {
	var req VehicleRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.db[req.LicensePlate]; ok {
		http.Error(w, "Car already exists", http.StatusConflict)
		return
	}

	if p.IsCarFull() {
		http.Error(w, "Parking lot for cars is full", http.StatusConflict)
		return
	}

	if p.IsFull() {
		http.Error(w, "Parking lot is full", http.StatusConflict)
		return
	}

	p.carCount++
	p.db[req.LicensePlate] = Car{LicensePlate: LicensePlate{licensePlate: req.LicensePlate}}
}

func (p *ParkingLotServer) HandleVan(w http.ResponseWriter, r *http.Request) {
	var req VehicleRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.db[req.LicensePlate]; ok {
		http.Error(w, "Van already exists", http.StatusConflict)
		return
	}

	if p.IsVanFull() {
		http.Error(w, "Parking lot for vans is full", http.StatusConflict)
		return
	}

	if p.IsFull() {
		http.Error(w, "Parking lot is full", http.StatusConflict)
		return
	}

	p.vanCount++
	p.db[req.LicensePlate] = Van{LicensePlate: LicensePlate{licensePlate: req.LicensePlate}}
}

func (p *ParkingLotServer) HandleRemoveCar(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("id")

	fmt.Println("Removing car with ID:", strID)
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.db[strID]; !ok {
		http.Error(w, "Car doesn't exist", http.StatusConflict)
		return
	}

	if p.IsEmpty() {
		http.Error(w, "Parking lot is empty", http.StatusConflict)
		return
	}

	p.carCount--
	delete(p.db, strID)
}

func (p *ParkingLotServer) HandleRemoveMotorcycle(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("id")

	p.mu.Unlock()
	defer p.mu.Lock()

	if _, ok := p.db[strID]; !ok {
		http.Error(w, "Motorcycle doesn't exist", http.StatusConflict)
		return
	}

	if p.IsEmpty() {
		http.Error(w, "Parking lot is empty", http.StatusConflict)
		return
	}

	p.motorcycleCount--
	delete(p.db, strID)
}

func (p *ParkingLotServer) HandleRemoveVan(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("id")

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.db[strID]; !ok {
		http.Error(w, "Van doesn't exist", http.StatusConflict)
		return
	}

	if p.IsEmpty() {
		http.Error(w, "Parking lot is empty", http.StatusConflict)
		return
	}

	p.vanCount--
	delete(p.db, strID)
}

func (p ParkingLotServer) HandleVanCount(w http.ResponseWriter, r *http.Request) {
	vans := p.HowManyVansTakingSpots()
	response := map[string]uint16{"vans": vans}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (p ParkingLotServer) HandleCap(w http.ResponseWriter, r *http.Request) {
	cap := p.Capacity()
	response := map[string]uint16{"capacity": cap}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (p ParkingLotServer) HandleAvailable(w http.ResponseWriter, r *http.Request) {
	available := p.Available()
	response := map[string]uint16{"available": available}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (p ParkingLotServer) Capacity() uint16 {
	return p.motorcycleLimit + p.compactLimit + p.regularLimit
}

func (p ParkingLotServer) Available() uint16 {
	return p.Capacity() - uint16(len(p.db))
}

func (p ParkingLotServer) IsFull() bool {
	return p.Available() == 0
}

func (p ParkingLotServer) IsEmpty() bool {
	return len(p.db) == 0
}

func (p ParkingLotServer) IsMotorcycleFull() bool {
	return p.motorcycleCount >= (p.motorcycleLimit + p.compactLimit + p.regularLimit - p.carCount - p.vanCount*3)
}

func (p ParkingLotServer) IsCarFull() bool {
	return p.carCount >= p.compactLimit
}

func (p ParkingLotServer) IsVanFull() bool {
	return p.vanCount >= (p.regularLimit / 3)
}

func (p ParkingLotServer) HowManyVansTakingSpots() uint16 {
	return p.vanCount * 3
}
