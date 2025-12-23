````````````````````````# 🚆 Train Seat Booking System,

## 📌 Problem Statement

Design and implement a train seat booking system that allocates seats to passengers based on availability and seat preference.

---

## 🧩 Design a Train Seat Booking System

### Problem Summary

- A train has multiple **bogies**
- Each bogey has a **fixed seat layout**
- Passengers provide **seat preferences**
- Seats can be booked **only once**
- The system:
  - Allocates preferred seats if available
  - Falls back to any available seat if not
  - Displays remaining seat availability

---

## 🏗️ Detailed Problem Statement

### Train Layout Rules

- The train consists of multiple **bogies (coaches)**
- Each bogey has **8 seats** in a fixed repeating pattern:

| Position | Seat Type |
|--------|----------|
| 1 | Lower Berth (LB) |
| 2 | Middle Berth (MB) |
| 3 | Upper Berth (UB) |
| 4 | Lower Berth (LB) |
| 5 | Middle Berth (MB) |
| 6 | Upper Berth (UB) |
| 7 | Side Lower (SL) |
| 8 | Side Upper (SU) |

- Total number of seats in the train: **72**

---

## 👤 Booking Requirements

Each passenger provides:
- **Name**
- **Preferred seat type** (`LB`, `MB`, `UB`, `SL`, `SU`)


### Seat Allocation Rules

1. Allocate the **preferred seat type** if available
2. If no preferred seat is available, allocate **any available seat**
3. If no seats are available, the booking **fails**
4. A seat can be booked **only once**

---

## ⚙️ Functional Requirements

The system must:

### 1️⃣ Initialize Seats
Each seat contains:
- Seat number
- Bogey number
- Seat type
- Booking status

### 2️⃣ Allocate Seats
- Process passengers **in order**
- Allocate seats based on preference and availability

### 3️⃣ Display Booking Details
For each passenger:
- Passenger name
- Preferred seat type
- Allocated seat type
- Seat number
- Bogey number

### 4️⃣ Display Remaining Seat Availability
- `1` → Available
- `0` → Booked
- Displayed **bogey-wise**

### 5️⃣ Booking Summary
Return:
- Booking ID
- List of allocated seats

---

## 📐 Constraints

- Total seats: **72**
- Seats per bogey: **8**
- Seat allocation follows **first-available strategy**

---

## Expected Result

The program should simulate a complete seat booking flow, correctly assigning seats based on preferences and availability, while maintaining accurate seat status throughout the process.


## 🚀 Possible Enhancements

- Add seat cancellation
- Introduce RAC / Waiting List
- Support multiple coaches dynamically
- Add REST API or UI layer
- Persist booking data

---
````````````````````````