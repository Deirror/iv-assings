package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestParkingLot(t *testing.T) {
	mux := http.NewServeMux()
	s := NewParkingLotServer(10, 30, 10)
	fmt.Println("s.Capacity ", s.Capacity())
	mux.HandleFunc("/car", s.HandleCar)
	mux.HandleFunc("/motorcycle", s.HandleMotorcycle)
	mux.HandleFunc("/van", s.HandleVan)

	mux.HandleFunc("/car/{id}", s.HandleRemoveCar)
	mux.HandleFunc("/motorcycle/{id}", s.HandleRemoveMotorcycle)
	mux.HandleFunc("/van/{id}", s.HandleRemoveVan)

	mux.HandleFunc("/van_count", s.HandleVanCount)
	mux.HandleFunc("/cap", s.HandleCap)
	mux.HandleFunc("/available", s.HandleAvailable)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Testing starts here
	wg := &sync.WaitGroup{}

	// Test adding cars
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			j := i % 9
			license := fmt.Sprintf("car%d", j)
			body := map[string]string{"licensePlate": license}
			jsonBody, _ := json.Marshal(body)

			resp, err := http.Post(ts.URL+"/car", "application/json", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			defer resp.Body.Close()

			msg, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Read body error: %v", err)
				return
			}

			// Catches duplicate cars
			if resp.StatusCode != http.StatusOK {
				fmt.Println(string(msg))
			}
		}(i)
		time.Sleep(1 * time.Millisecond)
	}
	wg.Wait()

	// Test adding motorcycles
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			j := i % 9
			license := fmt.Sprintf("motorcycle%d", j)
			body := map[string]string{"licensePlate": license}
			jsonBody, _ := json.Marshal(body)

			resp, err := http.Post(ts.URL+"/motorcycle", "application/json", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			defer resp.Body.Close()

			msg, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Read body error: %v", err)
				return
			}

			// Catches duplicate motorcycles
			if resp.StatusCode != http.StatusOK {
				fmt.Println(string(msg))
			}
		}(i)
		time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()

	// Test adding vans
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			license := fmt.Sprintf("van%d", i)
			body := map[string]string{"licensePlate": license}
			jsonBody, _ := json.Marshal(body)

			resp, err := http.Post(ts.URL+"/van", "application/json", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			defer resp.Body.Close()

			msg, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Read body error: %v", err)
				return
			}

			// Catches duplicate vans
			if resp.StatusCode != http.StatusOK {
				fmt.Println(string(msg))
			}
		}(i)
		time.Sleep(1 * time.Millisecond)
	}
	wg.Wait()

	// Test removing cars
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			j := i % 9
			license := fmt.Sprintf("car%d", j)
			url := fmt.Sprintf("%s/car?id=%s", ts.URL, license)

			// Delete the car
			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				log.Fatal(err)
			}

			// Send the request
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			// Handle empty body
			if resp.StatusCode != http.StatusOK {
				// Only read the body if it's not empty or needed for error handling
				msg, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Errorf("Read body error: %v", err)
					return
				}
				fmt.Println("Response:", string(msg))
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("Testing /cap")
	resp, err := http.Get(ts.URL + "/cap")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read body error: %v", err)
		return
	}
	fmt.Println(string(msg))

	fmt.Println("Testing /available")
	resp, err = http.Get(ts.URL + "/available")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	msg, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read body error: %v", err)
		return
	}
	fmt.Println(string(msg))
}
